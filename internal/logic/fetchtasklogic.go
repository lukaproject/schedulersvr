package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchTaskLogic {
	return &FetchTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchTaskLogic) FetchTask(req *types.FetchTaskReq) (resp *types.FetchTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
