package core

import "errors"

const (
	SchedulerMode_ByCommitTime = "BY_COMMIT_TIME"
	SchedulerMode_ByPriority   = "BY_PRIORITY"
	SchedulerMode_ByRandom     = "BY_RANDOM"

	MQType_Redis = "REDIS"
)

var (
	Err_TaskQueueEmpty = errors.New("scheduler: task queue empty")
)
