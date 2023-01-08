package core

import "github.com/zeromicro/go-zero/core/stores/redis"

type Scheduler interface {
	AddTask(t Task) error
	FetchTask(taskType string) (Task, error)
}

// a demo scheduler implement by redis.
type redisByCommitTimeSchedulerImpl struct {
	c     *Config
	redis *redis.Redis
}

func (r *redisByCommitTimeSchedulerImpl) AddTask(t Task) (err error) {
	_, err = r.redis.Lpush(t.GetType(), t.MustMarshal())
	return
}

func (r *redisByCommitTimeSchedulerImpl) FetchTask(taskType string) (t Task, err error) {
	var value string
	value, err = r.redis.Lpop(taskType)
	if err != nil {
		return nil, err
	}
	t = &UsualTask{}
	t.MustUnmarshal([]byte(value))
	return
}

func newRedisByCommitTime(c *Config) (ret Scheduler, err error) {
	err = c.Validate()
	if err != nil {
		return nil, err
	}
	ret = &redisByCommitTimeSchedulerImpl{
		c:     c,
		redis: c.RedisConfig.NewRedis(),
	}
	return
}
