package validate

import (
	"fmt"

	"github.com/lukaproject/schedulersvr/internal/types"
)

const (
	TaskType_LengthLimit             = 64
	TaskType_MaxTaskInQueLimit_Limit = 100000
)

const (
	TaskType_NameLegnthError = "length is illegal (except length 1 ~ 64)"
)

var (
	TaskType_MaxTaskInQueLimitError = fmt.Sprintf("is illegal (except -1 ~ %d)", TaskType_MaxTaskInQueLimit_Limit)
)

func LegalAddTaskTypeReq(req *types.AddTaskTypeReq) error {
	err := NewValidateError(req.SessionId)
	if req.Name == "" || len(req.Name) > TaskType_LengthLimit {
		err.Param2Error["req.Name"] = TaskType_NameLegnthError
	}
	if req.MaxTaskInQueLimit < -1 || req.MaxTaskInQueLimit > TaskType_MaxTaskInQueLimit_Limit {
		err.Param2Error["req.MaxTaskInQueLimit"] = TaskType_MaxTaskInQueLimitError
	}
	return ReturnValidateError(err)
}

func LegalGetTaskTypeReq(req *types.GetTaskTypeReq) error {
	err := NewValidateError(req.SessionId)
	if req.Name == "" || len(req.Name) > TaskType_LengthLimit {
		err.Param2Error["req.Name"] = TaskType_NameLegnthError
	}
	return ReturnValidateError(err)
}

func LegalDeleteTaskTypeReq(req *types.DeleteTaskTypeReq) error {
	err := NewValidateError(req.SessionId)
	if req.Name == "" || len(req.Name) > TaskType_LengthLimit {
		err.Param2Error["req.Name"] = TaskType_NameLegnthError
	}
	if len(err.Param2Error) == 0 {
		return nil
	}
	return ReturnValidateError(err)
}
