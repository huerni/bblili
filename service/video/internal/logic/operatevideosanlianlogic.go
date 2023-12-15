package logic

import (
	"bblili/service/video/internal/db"
	"context"
	"fmt"
	"gorm.io/gorm"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateVideoSanLianLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperateVideoSanLianLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateVideoSanLianLogic {
	return &OperateVideoSanLianLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperateVideoSanLianLogic) OperateVideoSanLian(in *video.OperateVideoSanLianRequest) (*video.OperateVideoSanLianResponse, error) {
	// 0: 点赞 1: 投币 2: 收藏
	switch in.Operator {
	case 0:
		videoLikes, err := db.QueryVideoLikeByUserIdAndVideoId(l.ctx, in.UserId, in.VideoId)
		if err != nil {
			return nil, err
		}
		if videoLikes == nil || len(videoLikes) == 0 {
			if err = db.InsertVideoLike(l.ctx, &db.VideoLike{
				Model:   gorm.Model{},
				UserId:  in.UserId,
				VideoId: in.VideoId,
			}); err != nil {
				return nil, err
			}
		} else {
			if err = db.DeleteVideoLike(l.ctx, uint64(videoLikes[0].ID)); err != nil {
				return nil, err
			}
		}
	case 1:
		videoCoins, err := db.QueryVideoCoinByVideoIdAndUserId(l.ctx, in.VideoId, in.UserId)
		if err != nil {
			return nil, err
		}
		if videoCoins == nil || videoCoins.ID == 0 {
			if err := db.InsertVideoCoin(l.ctx, &db.VideoCoin{
				Model:   gorm.Model{},
				UserId:  in.UserId,
				VideoId: in.VideoId,
				Amount:  1,
			}); err != nil {
				return nil, err
			}
		} else if videoCoins.Amount == 1 {
			if err := db.UpdateVideoCoin(l.ctx, &db.VideoCoin{
				Model:   gorm.Model{ID: videoCoins.ID},
				UserId:  in.UserId,
				VideoId: in.VideoId,
				Amount:  2,
			}); err != nil {
				return nil, err
			}
		} else if videoCoins.Amount == 2 {
			return nil, fmt.Errorf("只能投两币！")
		}
	case 2:
		videoCollections, err := db.QueryVideoCollectionByVideoIdAndUserIdAndGroupId(l.ctx, in.VideoId, in.UserId, in.GroupId)
		if err != nil {
			return nil, err
		}
		if videoCollections == nil || len(videoCollections) == 0 {
			if err = db.InsertVideoCollection(l.ctx, &db.VideoCollection{
				Model:   gorm.Model{},
				VideoId: in.VideoId,
				UserId:  in.UserId,
				GroupId: in.GroupId,
			}); err != nil {
				return nil, err
			}
		} else {
			if err = db.DeleteVideoCollection(l.ctx, uint64(videoCollections[0].ID)); err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("请确认参数正确")
	}
	return &video.OperateVideoSanLianResponse{}, nil
}
