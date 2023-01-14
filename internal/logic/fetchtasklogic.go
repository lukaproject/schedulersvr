package logic

import (
	"context"
	"time"

	"github.com/lukaproject/schedulersvr/gerrx"
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
	resp = &types.FetchTaskResp{
		SessionId: req.SessionId,
	}
	task, err := l.svcCtx.Scheduler.FetchTask(req.TaskType)
	gerrx.Must(err)
	modelTask, err := l.svcCtx.TaskTable.FindOne(l.ctx, task.GetId())
	gerrx.Must(err)
	modelTask.BeginTime.Scan(time.Now().UnixMilli())
	modelTask.WorkerId.Scan(req.WorkerId)
	err = l.svcCtx.TaskTable.Update(l.ctx, modelTask)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
		return
	}
	resp.Task = types.ToTaskContent(modelTask)
	return
}
