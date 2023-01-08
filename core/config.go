package core

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

var (
	ErrInvalidMQTpye = errors.New("invalid MQType")
)

type Config struct {
	SchedulerMode string
	MQType        string
	RedisConfig   redis.RedisConf
}

func (c *Config) Validate() error {
	if c.MQType == MQType_Redis {
		err := c.RedisConfig.Validate()
		if err != nil {
			return err
		}
	} else {
		return ErrInvalidMQTpye
	}
	return nil
}

func DefaultConfig() *Config {
	return &Config{
		SchedulerMode: SchedulerMode_ByCommitTime,
		MQType:        MQType_Redis,
	}
}
