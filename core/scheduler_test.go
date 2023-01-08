package core_test

import (
	"sync"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	"github.com/google/uuid"
	"github.com/lukaproject/schedulersvr/core"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func withScheduler(t *testing.T, c *core.Config, testFn func(sch core.Scheduler)) {
	s, err := miniredis.Run()
	require.Nil(t, err)
	c.RedisConfig = redis.RedisConf{
		Host: s.Addr(),
		Type: redis.NodeType,
	}
	sch, err := core.NewScheduler(c)
	require.Nil(t, err)
	testFn(sch)
}

func getTaskLoop(t *testing.T, wg *sync.WaitGroup, taskType string, expectTask []core.Task, sch core.Scheduler) {
	defer wg.Done()
	tick := time.NewTicker(time.Millisecond * 50)
	receiveTask := make([]core.Task, 0)
	for range tick.C {
		task, err := sch.FetchTask(taskType)
		if err != nil {
			t.Log(err)
			continue
		}
		receiveTask = append(receiveTask, task)
		if len(receiveTask) >= len(expectTask) {
			break
		}
	}
	t.Logf("taskType=%s, len expect=%d", taskType, len(expectTask))
	require.Equal(t, len(expectTask), len(receiveTask))
}
func Test_RedisByCommitTimeScheduler(t *testing.T) {
	c := core.DefaultConfig()
	c.MQType = core.MQType_Redis
	c.SchedulerMode = core.SchedulerMode_ByCommitTime

	testTask1_task := []core.Task{
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask1",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask1",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask1",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask1",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask1",
		},
	}

	testTask2_task := []core.Task{
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask2",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask2",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask2",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask2",
		},
		&core.UsualTask{
			Id:   uuid.NewString(),
			Type: "testTask2",
		},
	}

	withScheduler(t, c, func(sch core.Scheduler) {
		wg := &sync.WaitGroup{}
		wg.Add(3)
		go func() {
			defer wg.Done()
			for _, v := range testTask1_task {
				sch.AddTask(v)
			}
			for _, v := range testTask2_task {
				sch.AddTask(v)
			}
		}()
		go getTaskLoop(t, wg, "testTask1", testTask1_task, sch)
		go getTaskLoop(t, wg, "testTask2", testTask2_task, sch)
		wg.Wait()
	})
}
