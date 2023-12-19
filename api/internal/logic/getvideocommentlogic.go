package logic

import (
	"bblili/service/user/userclient"
	"bblili/service/video/videoclient"
	"context"

	"bblili/api/internal/svc"
	"bblili/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentLogic {
	return &GetVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoCommentLogic) GetVideoComment(req *types.GetVideoCommentRequest) (resp *types.GetVideoCommentResponse, err error) {

	res, err := l.svcCtx.VideoClient.GetVideoComments(l.ctx, &videoclient.GetVideoCommentsRequest{
		Size:    req.Size,
		Page:    req.Page,
		VideoId: req.VideoId,
	})

	videoComments := make([]*types.VideoComment, 0)
	usermap := make(map[uint64]*types.UserInfo, 0)

	for _, videocomment := range res.VideoComment {
		childlist := make([]*types.VideoComment, 0)
		for _, child := range videocomment.ChildList {
			typechild := types.VideoComment{
				VideoId:       child.VideoId,
				UserId:        child.UserId,
				Comment:       child.Comment,
				ReplyUserId:   child.ReplyUserId,
				RootId:        child.RootId,
				UserInfo:      nil,
				ReplyUserInfo: nil,
				ChildList:     nil,
			}
			value, exists := usermap[child.UserId]
			if exists {
				typechild.UserInfo = value
			} else {
				userInfo, err := l.getUserInfo(child.UserId)
				if err != nil {
					return
				}
				usermap[child.UserId] = userInfo
				typechild.UserInfo = userInfo
			}

			value, exists = usermap[child.ReplyUserId]
			if exists {
				typechild.ReplyUserInfo = value
			} else {
				userInfo, err := l.getUserInfo(child.ReplyUserId)
				if err != nil {
					return
				}
				usermap[child.ReplyUserId] = userInfo
				typechild.ReplyUserInfo = userInfo
			}

			childlist = append(childlist, &typechild)
		}

		typecomment := types.VideoComment{
			VideoId:     videocomment.VideoId,
			UserId:      videocomment.UserId,
			Comment:     videocomment.Comment,
			ReplyUserId: videocomment.ReplyUserId,
			RootId:      videocomment.RootId,
			ChildList:   childlist,
		}

		value, exists := usermap[videocomment.UserId]
		if exists {
			typecomment.UserInfo = value
		} else {
			userInfo, err := l.getUserInfo(videocomment.UserId)
			if err != nil {
				return
			}
			usermap[videocomment.UserId] = userInfo
			typecomment.UserInfo = userInfo
		}

		value, exists = usermap[videocomment.ReplyUserId]
		if exists {
			typecomment.ReplyUserInfo = value
		} else {
			userInfo, err := l.getUserInfo(videocomment.ReplyUserId)
			if err != nil {
				return
			}
			usermap[videocomment.ReplyUserId] = userInfo
			typecomment.ReplyUserInfo = userInfo
		}

		videoComments = append(videoComments, &typecomment)
	}

	return
}

func (l *GetVideoCommentLogic) getUserInfo(userId uint64) (*types.UserInfo, error) {
	response, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &userclient.GetUserInfoRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		Id:     response.UserInfo.Id,
		Nick:   response.UserInfo.Nick,
		Avatar: response.UserInfo.Avatar,
		Sign:   response.UserInfo.Sign,
		Gender: response.UserInfo.Gender,
		Birth:  response.UserInfo.Birth,
	}, nil
}
