package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/internal/logic/validate"
	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTaskTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskTypeLogic {
	return &AddTaskTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTaskTypeLogic) AddTaskType(req *types.AddTaskTypeReq) (resp *types.GeneralResp, err error) {
	l.Logger.Infof("addTaskTypeReq=%v", req)
	err = validate.LegalAddTaskTypeReq(req)
	if err != nil {
		return
	}
	err = l.svcCtx.Dbc.AddTaskType(l.ctx, types.TaskTypeContent{
		Id:                types.GenTaskTypeId(),
		Name:              req.Name,
		MaxTaskInQueLimit: req.MaxTaskInQueLimit,
		ExtraInfo:         req.ExtraInfo,
	})
	if err != nil {
		l.Logger.Error(err)
		return
	}
	return
}
