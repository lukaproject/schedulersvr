package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/db/model"
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
	var modelTask *model.Task
	modelTask, err = l.svcCtx.TaskTable.FindOne(l.ctx, req.TaskId)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
		return
	}
	resp.Task = types.ToTaskContent(modelTask)
	return
}
