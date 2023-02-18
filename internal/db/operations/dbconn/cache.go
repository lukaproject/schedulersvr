package dbconn

import "github.com/lukaproject/schedulersvr/internal/types"

type Cache interface {
	AddTaskType(taskType *types.TaskTypeContent)
	GetTaskType(taskTypeName string)
	DeleteTaskType(taskTypeName string)
}
