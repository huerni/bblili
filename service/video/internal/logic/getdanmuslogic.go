package logic

import (
	"bblili/service/video/internal/db"
	"context"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDanmusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDanmusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDanmusLogic {
	return &GetDanmusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDanmusLogic) GetDanmus(in *video.GetDanmusRequest) (*video.GetDanmusResponse, error) {

	danmus, err := db.QueryDanMuByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, err
	}
	danmuInfos := make([]*video.DanMuInfo, 0)
	for _, danmu := range danmus {
		danmuInfos = append(danmuInfos, &video.DanMuInfo{
			UserId:    danmu.UserId,
			VideoId:   danmu.VideoId,
			Content:   danmu.Content,
			DanmuTime: danmu.DanmuTime,
		})
	}
	return &video.GetDanmusResponse{DanmuInfo: danmuInfos}, nil
}
