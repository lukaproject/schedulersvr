package core

import (
	"errors"
)

var (
	ErrNoSuitableScheduler = errors.New("no suitable scheduler")
)

func NewScheduler(c *Config) (ret Scheduler, err error) {
	if c.MQType == MQType_Redis {
		if c.SchedulerMode == SchedulerMode_ByCommitTime {
			return newRedisByCommitTime(c)
		}
	}
	return nil, ErrNoSuitableScheduler
}
