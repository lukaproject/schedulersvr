package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type extraTaskTypeModel interface {
	DeleteTaskTypeByName(ctx context.Context, taskTypeName string) (err error)
}

type iExtraTaskTypeModel struct {
	conn  sqlx.SqlConn
	table string
}

func newExtraTaskTypeModel(conn sqlx.SqlConn) *iExtraTaskTypeModel {
	return &iExtraTaskTypeModel{
		conn:  conn,
		table: "`task_type`",
	}
}

func (m *iExtraTaskTypeModel) DeleteTaskTypeByName(ctx context.Context, taskTypeName string) (err error) {
	return
}
