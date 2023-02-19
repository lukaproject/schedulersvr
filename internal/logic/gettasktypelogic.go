package logic

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lukaproject/schedulersvr/gerrx"
	"github.com/lukaproject/schedulersvr/internal/logic/validate"
	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskTypeLogic {
	return &GetTaskTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskTypeLogic) GetTaskType(req *types.GetTaskTypeReq) (resp *types.GetTaskTypeResp, err error) {
	err = validate.LegalGetTaskTypeReq(req)
	if err != nil {
		return
	}
	resp = &types.GetTaskTypeResp{
		SessionId: req.SessionId,
	}
	l.Logger.Infof("GetTaskTypeReq=%v", req)
	ttc, err := l.svcCtx.Dbc.GetTaskTypeByName(l.ctx, req.Name)
	if err != nil {
		l.Logger.Error(err)
		err = gerrx.NewDefaultHttpError(http.StatusNotFound, fmt.Sprintf("task type name = %s not found", req.Name))
		return
	}
	resp.TaskType = ttc
	return
}
