package types

import (
	"time"

	"github.com/lukaproject/schedulersvr/core"
)

func (atr *AddTaskReq) GenCoreTask() core.Task {
	utask := &core.UsualTask{
		Id:    GenTaskId(),
		Type:  atr.TaskType,
		Input: atr.Input,
	}
	return utask
}

func GenTaskContent(atr *AddTaskReq, t core.Task) *TaskContent {
	return &TaskContent{
		Id:         t.GetId(),
		Name:       atr.Name,
		Input:      t.GetInput(),
		TaskType:   t.GetType(),
		CommitTime: uint64(time.Now().UnixMilli()),
	}
}
