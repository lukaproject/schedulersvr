package svc

import (
	"github.com/lukaproject/atur"
	"github.com/lukaproject/schedulersvr/core"
	"github.com/lukaproject/schedulersvr/gerrx"
	"github.com/lukaproject/schedulersvr/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	Scheduler core.Scheduler
	TaskStore atur.KvStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	sch, err := core.NewScheduler(c.Scheduler.MustToCoreConfig())
	if err != nil {
		panic(err)
	}
	svc := &ServiceContext{
		Config:    c,
		Scheduler: sch,
	}
	kvConf := atur.NewKvConfig()
	svc.TaskStore, err = atur.NewKvStore(kvConf)
	gerrx.Must(err)
	return svc
}
