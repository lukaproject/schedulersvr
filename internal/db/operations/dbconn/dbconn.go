package dbconn

import (
	"context"
	"time"

	"github.com/lukaproject/schedulersvr/internal/db/operations/model"
	"github.com/lukaproject/schedulersvr/internal/types"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DBConn interface {
	AddTaskType(context.Context, types.TaskTypeContent) error
	GetTaskTypeByName(ctx context.Context, taskTypeName string) (ttc types.TaskTypeContent, err error)
	DeleteTaskTypeByName(ctx context.Context, taskTypeName string) (err error)
}

type impl struct {
	mTaskType model.TaskTypeModel
}

func (idbc *impl) AddTaskType(ctx context.Context, ttc types.TaskTypeContent) (err error) {
	_, err = idbc.mTaskType.AddTaskTypeWithCreateTime(ctx, &model.TaskType{
		TaskTypeId:        ttc.Id,
		TaskTypeName:      ttc.Name,
		MaxTaskInQueLimit: ttc.MaxTaskInQueLimit,
		CreateTime:        time.Now().UnixMilli(),
		ExtraInfo:         ttc.ExtraInfo,
	})
	return
}

func (idbc *impl) GetTaskTypeByName(ctx context.Context, taskTypeName string) (ttc types.TaskTypeContent, err error) {
	result, err := idbc.mTaskType.FindOneByTaskTypeName(ctx, taskTypeName)
	if err != nil {
		return
	}
	ttc = types.TaskTypeContent{
		Id:                result.TaskTypeId,
		Name:              result.TaskTypeName,
		MaxTaskInQueLimit: result.MaxTaskInQueLimit,
		CreateTime:        result.CreateTime,
		ExtraInfo:         result.ExtraInfo,
	}
	return
}

func (idbc *impl) DeleteTaskTypeByName(ctx context.Context, taskTypeName string) (err error) {
	err = idbc.mTaskType.DeleteTaskTypeByName(ctx, taskTypeName)
	return
}

func NewDBConn(conn sqlx.SqlConn) DBConn {
	return &impl{
		mTaskType: model.NewTaskTypeModel(conn),
	}
}
