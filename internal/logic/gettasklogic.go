package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/gerrx"
	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskLogic {
	return &GetTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskLogic) GetTask(req *types.GetTaskReq) (resp *types.GetTaskResp, err error) {
	resp = &types.GetTaskResp{
		SessionId: req.SessionId,
	}
	taskContent := types.TaskContent{}
	err = l.svcCtx.TaskStore.GetCtx(l.ctx, []byte(req.TaskId), &taskContent)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
		return
	}
	resp.Task = taskContent
	return
}
