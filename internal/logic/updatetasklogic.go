package logic

import (
	"context"
	"time"

	"github.com/lukaproject/schedulersvr/db/model"
	"github.com/lukaproject/schedulersvr/gerrx"
	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskLogic {
	return &UpdateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTaskLogic) UpdateTask(req *types.UpdateTaskReq) (resp *types.GeneralResp, err error) {
	resp = &types.GeneralResp{
		SessionId: req.SessionId,
	}
	var modelTask *model.Task
	modelTask, err = l.svcCtx.TaskTable.FindOne(l.ctx, req.Task.Id)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
		return
	}
	modelTask.EndTime.Scan(time.Now().UnixMilli())
	modelTask.Output.Scan(req.Task.Output)
	modelTask.Status.Scan(req.Task.Status)
	err = l.svcCtx.TaskTable.Update(l.ctx, modelTask)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
	}
	return
}
