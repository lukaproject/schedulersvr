package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HeartBeatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHeartBeatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HeartBeatLogic {
	return &HeartBeatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HeartBeatLogic) HeartBeat(req *types.HeartBeatReq) (resp *types.HeartBeatResp, err error) {
	// todo: add your logic here and delete this line

	return
}
