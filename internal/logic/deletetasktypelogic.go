package logic

import (
	"context"

	"github.com/lukaproject/schedulersvr/internal/logic/validate"
	"github.com/lukaproject/schedulersvr/internal/svc"
	"github.com/lukaproject/schedulersvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTaskTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskTypeLogic {
	return &DeleteTaskTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTaskTypeLogic) DeleteTaskType(req *types.DeleteTaskTypeReq) (resp *types.GeneralResp, err error) {
	err = validate.LegalDeleteTaskTypeReq(req)
	if err != nil {
		return
	}
	resp = &types.GeneralResp{
		SessionId: req.SessionId,
	}
	err = l.svcCtx.Dbc.DeleteTaskTypeByName(l.ctx, req.Name)
	if err != nil {
		l.Logger.Error(err)
		return
	}
	return
}
