package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Scheduler SchedulerConfig
	DB        DbConfig
}

func (c Config) Tell() {
	c.Scheduler.Tell()
	c.DB.Tell()
}
