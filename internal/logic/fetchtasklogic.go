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
	taskContent := types.TaskContent{}
	err = l.svcCtx.TaskStore.GetCtx(l.ctx, []byte(task.GetId()), &taskContent)
	gerrx.Must(err)
	taskContent.BeginTime = uint64(time.Now().UnixMilli())
	taskContent.WorkerId = req.WorkerId
	err = l.svcCtx.TaskStore.SetCtx(l.ctx, []byte(taskContent.Id), &taskContent)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
		return
	}
	resp.Task = taskContent
	return
}
