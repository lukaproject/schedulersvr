package logic

import (
	"context"
	"time"

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
	taskContent := types.TaskContent{}
	err = l.svcCtx.TaskStore.GetCtx(l.ctx, []byte(req.Task.Id), &taskContent)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
		return
	}
	taskContent.EndTime = uint64(time.Now().UnixMilli())
	taskContent.Output = req.Task.Output
	taskContent.Status = req.Task.Status
	err = l.svcCtx.TaskStore.SetCtx(l.ctx, []byte(taskContent.Id), &taskContent)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultError(err.Error(), req.SessionId)
	}
	return
}
