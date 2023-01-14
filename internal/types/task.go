package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/lukaproject/schedulersvr/core"
	"github.com/lukaproject/schedulersvr/db/model"
	"github.com/lukaproject/schedulersvr/gerrx"
)

func (atr *AddTaskReq) GenCoreTask() core.Task {
	utask := &core.UsualTask{
		Id:    uuid.NewString(),
		Type:  atr.TaskType,
		Input: atr.Input,
	}
	return utask
}

func GenTaskContent(atr *AddTaskReq, t core.Task) TaskContent {
	return TaskContent{
		Id:         t.GetId(),
		Name:       atr.Name,
		Input:      t.GetInput(),
		TaskType:   t.GetType(),
		CommitTime: uint64(time.Now().UnixMilli()),
	}
}

func ToTaskModel(tc *TaskContent) *model.Task {
	ret := &model.Task{}
	ret.Id = tc.Id
	ret.Name = tc.Name
	ret.TaskType = tc.TaskType
	gerrx.Must(ret.Input.Scan(tc.Input))
	gerrx.Must(ret.BeginTime.Scan(int64(tc.BeginTime)))
	gerrx.Must(ret.CommitTime.Scan(int64(tc.CommitTime)))
	gerrx.Must(ret.EndTime.Scan(int64(tc.EndTime)))
	gerrx.Must(ret.Output.Scan(tc.Output))
	gerrx.Must(ret.Status.Scan(tc.Status))
	gerrx.Must(ret.WorkerId.Scan(tc.WorkerId))
	return ret
}

func ToTaskContent(mt *model.Task) TaskContent {
	tc := TaskContent{}
	tc.Id = mt.Id
	tc.Name = mt.Name
	tc.TaskType = mt.TaskType
	tc.Input = mt.Input.String
	tc.BeginTime = uint64(mt.BeginTime.Int64)
	tc.CommitTime = uint64(mt.CommitTime.Int64)
	tc.EndTime = uint64(mt.EndTime.Int64)
	tc.Output = mt.Output.String
	tc.Status = mt.Status.String
	tc.WorkerId = mt.WorkerId.String
	return tc
}
