package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type extraTaskTypeModel interface {
	DeleteTaskTypeByName(ctx context.Context, taskTypeName string) (err error)
	AddTaskTypeWithCreateTime(ctx context.Context, taskType *TaskType) (ret sql.Result, err error)
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
	query := fmt.Sprintf("delete from %s where `task_type_name` = ?", m.table)
	_, err = m.conn.ExecCtx(ctx, query, taskTypeName)
	return err
}

func (m *iExtraTaskTypeModel) AddTaskTypeWithCreateTime(ctx context.Context, taskType *TaskType) (ret sql.Result, err error) {
	query := fmt.Sprintf("insert into %s values (?, ?, ?, ?, ?)", m.table)
	ret, err = m.conn.ExecCtx(ctx, query,
		taskType.TaskTypeId,
		taskType.TaskTypeName,
		taskType.MaxTaskInQueLimit,
		taskType.CreateTime,
		taskType.ExtraInfo)
	return ret, err
}
