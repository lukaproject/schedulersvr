package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveWorkerByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveWorkerByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveWorkerByNameLogic {
	return &RemoveWorkerByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveWorkerByNameLogic) RemoveWorkerByName(req *types.RemoveWorkerByNameReq) (resp *types.GeneralResp, err error) {
	// todo: add your logic here and delete this line

	return
}
