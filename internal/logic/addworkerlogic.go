package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddWorkerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddWorkerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddWorkerLogic {
	return &AddWorkerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddWorkerLogic) AddWorker(req *types.AddWorkerReq) (resp *types.GeneralResp, err error) {
	// todo: add your logic here and delete this line

	return
}
