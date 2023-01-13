package config

import (
	"github.com/lukaproject/schedulersvr/core"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SchedulerConfig struct {
	SchedulerMode string
	MQType        string
	Redis         redis.RedisConf
}

func (schC SchedulerConfig) ToCoreConfig() (c *core.Config, err error) {
	c = &core.Config{}
	c.MQType = schC.MQType
	c.SchedulerMode = schC.SchedulerMode
	c.RedisConfig = schC.Redis
	err = c.Validate()
	return
}

func (schC SchedulerConfig) MustToCoreConfig() *core.Config {
	c, err := schC.ToCoreConfig()
	if err != nil {
		panic(err)
	}
	return c
}

func (schC SchedulerConfig) Tell() {
	logx.Infof("SchedulerMode=%s", schC.SchedulerMode)
	logx.Infof("MQType=%s", schC.MQType)
	logx.Infof("Redis.Addr=%s", schC.Redis.Host)
	logx.Infof("Redis.Type=%s", schC.Redis.Type)
	logx.Infof("Redis.Pass=%s", schC.Redis.Pass)
}
