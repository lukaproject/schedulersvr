package svc

import (
	"github.com/lukaproject/schedulersvr/core"
	"github.com/lukaproject/schedulersvr/db/model"
	"github.com/lukaproject/schedulersvr/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	mysqlName = "mysql"
)

type ServiceContext struct {
	Config    config.Config
	Scheduler core.Scheduler

	SqlConn   sqlx.SqlConn
	TaskTable model.TaskModel
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
	svc.SqlConn = sqlx.NewSqlConn(mysqlName, c.DB.Mysql.Datasource)
	svc.TaskTable = model.NewTaskModel(svc.SqlConn)
	return svc
}
