package logic

import (
	"bblili/service/video/internal/common/constant"
	"bblili/service/video/internal/db"
	"bblili/service/video/internal/svc"
	"bblili/service/video/video"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentsLogic {
	return &GetVideoCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoCommentsLogic) GetVideoComments(in *video.GetVideoCommentsRequest) (*video.GetVideoCommentsResponse, error) {
	videoComments := make([]*video.VideoComment, 0)

	key := fmt.Sprintf(constant.VIDEOCOMMENT_SORTEDSET, in.VideoId, 0)
	isexist, err := l.svcCtx.RedisClient.Exists(key)
	if err != nil {
		return nil, err
	}

	if isexist {
		lockKey := fmt.Sprintf(constant.LOCK_VIDEOCOMMENT_SORTEDSET, in.VideoId)
		mutex := l.svcCtx.RedisSync.NewMutex(lockKey)
		// 对key进行
		if err := mutex.Lock(); err != nil {
			panic(err)
		}
		defer mutex.Unlock()

		isexist, err := l.svcCtx.RedisClient.Exists(key)
		if err != nil {
			return nil, err
		}
		if isexist {
			redisres, err := l.svcCtx.RedisClient.Zrevrange(key, int64((in.Page-1)*in.Size), int64(in.Page*in.Size))
			if err != nil {
				return nil, err
			}
			childVideoComments := make([]*video.VideoComment, 0)
			for _, bytecomment := range redisres {
				var comment db.VideoComment
				err := json.Unmarshal([]byte(bytecomment), &comment)
				if err != nil {
					return nil, err
				}
				// 查询子评论 childVideoComments
				childKey := fmt.Sprintf(constant.VIDEOCOMMENT_SORTEDSET, in.VideoId, comment.ID)
				ischildexist, err := l.svcCtx.RedisClient.Exists(childKey)
				if err != nil {
					return nil, err
				}
				if ischildexist {
					childres, err := l.svcCtx.RedisClient.Zrevrange(childKey, 0, -1)
					if err != nil {
						return nil, err
					}
					for _, bytechild := range childres {
						var childcomment db.VideoComment
						err := json.Unmarshal([]byte(bytechild), &childcomment)
						if err != nil {
							return nil, err
						}
						childVideoComments = append(childVideoComments, &video.VideoComment{
							VideoId:     childcomment.VideoId,
							UserId:      childcomment.UserId,
							Comment:     childcomment.Comment,
							ReplyUserId: childcomment.ReplyUserId,
							RootId:      childcomment.RootId,
							ChildList:   nil,
						})
					}
				} else {
					// 数据库中查询
					childComments, err := db.QueryVideoChildCommentByVideoId(l.ctx, in.VideoId, uint64(comment.ID))
					if err != nil {
						return nil, err
					}
					childVideoComments := make([]*video.VideoComment, 0)
					for j := 0; j < len(childComments); j++ {
						childVideoComment := &video.VideoComment{
							VideoId:     childComments[j].VideoId,
							UserId:      childComments[j].UserId,
							Comment:     childComments[j].Comment,
							ReplyUserId: childComments[j].ReplyUserId,
							RootId:      childComments[j].RootId,
							ChildList:   nil,
						}

						childVideoComments = append(childVideoComments, childVideoComment)
					}
				}
				videoComments = append(videoComments, &video.VideoComment{
					VideoId:     comment.VideoId,
					UserId:      comment.UserId,
					Comment:     comment.Comment,
					ReplyUserId: comment.ReplyUserId,
					RootId:      comment.RootId,
					ChildList:   childVideoComments,
				})
			}
		}
	} else {
		comments, err := db.PageQueryVideoCommentByVideoId(l.ctx, in.Page, in.Size, in.VideoId)
		if err != nil {
			return nil, err
		}

		if len(comments) > 0 {
			for i := 0; i < len(comments); i++ {
				comment := comments[i]
				childComments, err := db.QueryVideoChildCommentByVideoId(l.ctx, in.VideoId, uint64(comment.ID))
				if err != nil {
					return nil, err
				}
				childVideoComments := make([]*video.VideoComment, 0)
				for j := 0; j < len(childComments); j++ {
					childVideoComment := &video.VideoComment{
						VideoId:     childComments[j].VideoId,
						UserId:      childComments[j].UserId,
						Comment:     childComments[j].Comment,
						ReplyUserId: childComments[j].ReplyUserId,
						RootId:      childComments[j].RootId,
						ChildList:   nil,
					}

					childVideoComments = append(childVideoComments, childVideoComment)
				}

				videoComment := &video.VideoComment{
					VideoId:     comment.VideoId,
					UserId:      comment.UserId,
					Comment:     comment.Comment,
					ReplyUserId: comment.ReplyUserId,
					RootId:      comment.RootId,
					ChildList:   childVideoComments,
				}
				videoComments = append(videoComments, videoComment)
			}
		}
	}

	return &video.GetVideoCommentsResponse{VideoComment: videoComments}, nil
}
