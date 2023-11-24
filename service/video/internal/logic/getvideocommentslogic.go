package logic

import (
	"bblili/service/user/userclient"
	"bblili/service/video/internal/db"
	"context"
	"fmt"

	"bblili/service/video/internal/svc"
	"bblili/service/video/video"

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
	queryVideo, res := db.QueryVideoByVideoId(l.ctx, in.VideoId)
	if res != nil {
		return nil, res
	}
	if queryVideo == nil {
		return nil, fmt.Errorf("该视频不存在！")
	}
	comments, err := db.PageQueryVideoCommentByVideoId(l.ctx, in.Page, in.Size, in.VideoId)
	if err != nil {
		return nil, res
	}
	videoComments := make([]*video.VideoComment, 0)
	if len(comments) > 0 {
		usermap := make(map[uint64]*video.UserInfo, 0)
		for i := 0; i < len(comments); i++ {
			comment := comments[i]
			childComments, err := db.QueryVideoChildCommentByVideoId(l.ctx, in.VideoId, uint64(comment.ID))
			if err != nil {
				return nil, err
			}
			childVideoComments := make([]*video.VideoComment, 0)
			for j := 0; j < len(childComments); j++ {
				childVideoComment := &video.VideoComment{
					VideoId:       childComments[j].VideoId,
					UserId:        childComments[j].UserId,
					Comment:       childComments[j].Comment,
					ReplyUserId:   childComments[j].ReplyUserId,
					RootId:        childComments[j].RootId,
					UserInfo:      nil,
					ReplyUserInfo: nil,
					ChildList:     nil,
				}

				value, exists := usermap[childComments[j].UserId]
				if exists {
					childVideoComment.UserInfo = value
				} else {
					userInfo, err := l.getUserInfo(childComments[j].UserId)
					if err != nil {
						return nil, err
					}
					childVideoComment.UserInfo = userInfo
					usermap[childComments[j].UserId] = userInfo

				}
				value, exists = usermap[childComments[j].ReplyUserId]
				if exists {
					childVideoComment.ReplyUserInfo = value
				} else {
					userInfo, err := l.getUserInfo(childComments[j].ReplyUserId)
					if err != nil {
						return nil, err
					}
					childVideoComment.UserInfo = userInfo
					usermap[childComments[j].ReplyUserId] = userInfo
				}

				childVideoComments = append(childVideoComments, childVideoComment)
			}

			videoComment := &video.VideoComment{
				VideoId:       comment.VideoId,
				UserId:        comment.UserId,
				Comment:       comment.Comment,
				ReplyUserId:   comment.ReplyUserId,
				RootId:        comment.RootId,
				UserInfo:      nil,
				ReplyUserInfo: nil,
				ChildList:     nil,
			}
			value, exists := usermap[comment.UserId]
			if exists {
				videoComment.UserInfo = value
			} else {
				userInfo, err := l.getUserInfo(comment.UserId)
				if err != nil {
					return nil, err
				}
				videoComment.UserInfo = userInfo
				usermap[comment.UserId] = userInfo
			}
			value, exists = usermap[comment.ReplyUserId]
			if exists {
				videoComment.ReplyUserInfo = value
			} else {
				//  查询userinfo
				userInfo, err := l.getUserInfo(comment.ReplyUserId)
				if err != nil {
					return nil, err
				}
				videoComment.UserInfo = userInfo
				usermap[comment.ReplyUserId] = userInfo
			}
			videoComment.ChildList = childVideoComments
			videoComments = append(videoComments, videoComment)
		}
	}
	return &video.GetVideoCommentsResponse{VideoComment: videoComments}, nil
}

func (l *GetVideoCommentsLogic) getUserInfo(userId uint64) (*video.UserInfo, error) {
	response, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &userclient.GetUserInfoRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &video.UserInfo{
		Id:     response.UserInfo.Id,
		Nick:   response.UserInfo.Nick,
		Avatar: response.UserInfo.Avatar,
		Sign:   response.UserInfo.Sign,
		Gender: response.UserInfo.Gender,
		Birth:  response.UserInfo.Birth,
	}, nil
}
