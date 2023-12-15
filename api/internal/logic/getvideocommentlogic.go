package logic

import (
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
	for _, videocomment := range res.VideoComment {
		childlist := make([]*types.VideoComment, 0)
		for _, child := range videocomment.ChildList {
			childlist = append(childlist, &types.VideoComment{
				VideoId:     child.VideoId,
				UserId:      child.UserId,
				Comment:     child.Comment,
				ReplyUserId: child.ReplyUserId,
				RootId:      child.RootId,
				UserInfo: &types.UserInfo{
					Id:     child.UserInfo.Id,
					Nick:   child.UserInfo.Nick,
					Avatar: child.UserInfo.Avatar,
					Sign:   child.UserInfo.Sign,
					Gender: child.UserInfo.Gender,
					Birth:  child.UserInfo.Birth,
				},
				ReplyUserInfo: &types.UserInfo{
					Id:     child.ReplyUserInfo.Id,
					Nick:   child.ReplyUserInfo.Nick,
					Avatar: child.ReplyUserInfo.Avatar,
					Sign:   child.ReplyUserInfo.Sign,
					Gender: child.ReplyUserInfo.Gender,
					Birth:  child.ReplyUserInfo.Birth,
				},
				ChildList: nil,
			})
		}
		videoComments = append(videoComments, &types.VideoComment{
			VideoId:     videocomment.VideoId,
			UserId:      videocomment.UserId,
			Comment:     videocomment.Comment,
			ReplyUserId: videocomment.ReplyUserId,
			RootId:      videocomment.RootId,
			UserInfo: &types.UserInfo{
				Id:     videocomment.UserInfo.Id,
				Nick:   videocomment.UserInfo.Nick,
				Avatar: videocomment.UserInfo.Avatar,
				Sign:   videocomment.UserInfo.Sign,
				Gender: videocomment.UserInfo.Gender,
				Birth:  videocomment.UserInfo.Birth,
			},
			ReplyUserInfo: &types.UserInfo{
				Id:     videocomment.ReplyUserInfo.Id,
				Nick:   videocomment.ReplyUserInfo.Nick,
				Avatar: videocomment.ReplyUserInfo.Avatar,
				Sign:   videocomment.ReplyUserInfo.Sign,
				Gender: videocomment.ReplyUserInfo.Gender,
				Birth:  videocomment.ReplyUserInfo.Birth,
			},
			ChildList: childlist,
		})
	}

	return
}
