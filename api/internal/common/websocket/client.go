package mywebsoc

import (
	"bblili/api/internal/svc"
	"bblili/service/video/videoclient"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var videoRooms = make(map[uint64]*VideoRoom)

type Client struct {
	videoRoom *VideoRoom
	conn      *websocket.Conn
	send      chan []byte
}

func (c *Client) readPump(ctx context.Context, svcCtx *svc.ServiceContext) {
	// socket连接断开时执行defer，在此处对用户离开直播间行为进行处理
	defer func() {
		// 离开视频 观看数量少一
		if c.videoRoom != nil {
			delete(c.videoRoom.clients, c) // 移除直播间的当前连接
			close(c.send)
			msg := make(map[string]interface{})
			msg["msgtype"] = 1
			msg["viewNum"] = len(c.videoRoom.clients)
			msgbytes, _ := json.Marshal(msg)
			c.videoRoom.broadcast <- msgbytes
		}
		c.conn.Close()
	}()
	//c.conn.SetReadLimit(maxMessageSize)
	//c.conn.SetReadDeadline(time.Now().Add(pongWait))
	//c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		log.Println(string(message))
		// 解析json文本
		var r map[string]interface{}
		if err := json.Unmarshal(message, &r); err != nil {
			log.Println(err)
			continue
		}
		videoId := uint64(r["videoId"].(float64))

		msgtype := int(r["msgtype"].(float64))
		msg := make(map[string]interface{})
		msg["msgtype"] = msgtype
		var msgbytes []byte
		// 根据不同的消息类型分别处理
		switch msgtype {
		case 1: // 进入视频 观看人数变化时自动更新

			if _, ok := videoRooms[videoId]; !ok {
				log.Println(videoId, "房间发言频道不存在 -> 创建")
				// 模拟根据房间号找主播昵称
				videoRooms[videoId] = newVideoRoom(videoId)
				go videoRooms[videoId].run()
			}
			c.videoRoom = videoRooms[videoId]
			c.videoRoom.clients[c] = true
			msg["viewNum"] = len(c.videoRoom.clients)
			msgbytes, _ = json.Marshal(msg)
			c.videoRoom.broadcast <- msgbytes
		case 2: // 更新弹幕
			currentTime := int32(r["currentTime"].(float64))
			resp, err := svcCtx.VideoClient.GetBarrages(ctx, &videoclient.GetBarragesRequest{VideoId: videoId, CurrentTime: currentTime})
			if err != nil {
				log.Println(err)
				break
			}
			msg["barrageInfos"] = resp.BarrageInfos
			msgbytes, _ = json.Marshal(msg)
			c.send <- msgbytes
		}
	}
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// 定时发送
func (c *Client) writePump() {
	ticker := time.NewTicker(60 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServerBarrageWs(w http.ResponseWriter, r *http.Request, svcCtx *svc.ServiceContext) {
	// 允许跨域
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// 将请求升级为websocket连接，代表一个观众进来
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 观看人数变化时自动更新； 每隔60s定时获取后60s弹幕；
	client := &Client{conn: conn, send: make(chan []byte, 256)}
	go client.writePump()
	go client.readPump(context.Background(), svcCtx)
}
