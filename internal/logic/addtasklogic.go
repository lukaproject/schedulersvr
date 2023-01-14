package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskLogic {
	return &AddTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTaskLogic) AddTask(req *types.AddTaskReq) (resp *types.AddTaskResp, err error) {
	task := req.GenCoreTask()
	resp = &types.AddTaskResp{
		SessionId: req.SessionId,
		TaskId:    task.GetId(),
		Name:      req.Name,
	}
	taskContent := types.GenTaskContent(req, task)
	_, err = l.svcCtx.TaskTable.Insert(l.ctx, types.ToTaskModel(&taskContent))
	if err != nil {
		return
	}
	err = l.svcCtx.Scheduler.AddTask(task)
	if err != nil {
		return
	}
	return
}
