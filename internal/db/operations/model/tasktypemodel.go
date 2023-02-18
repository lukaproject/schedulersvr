package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskTypeModel = (*customTaskTypeModel)(nil)

type (
	// TaskTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskTypeModel.
	TaskTypeModel interface {
		taskTypeModel
		extraTaskTypeModel
	}

	customTaskTypeModel struct {
		*defaultTaskTypeModel
		*iExtraTaskTypeModel
	}
)

// NewTaskTypeModel returns a model for the database table.
func NewTaskTypeModel(conn sqlx.SqlConn) TaskTypeModel {
	return &customTaskTypeModel{
		defaultTaskTypeModel: newTaskTypeModel(conn),
		iExtraTaskTypeModel:  newExtraTaskTypeModel(conn),
	}
}
