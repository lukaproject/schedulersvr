package apitest

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	apitest_common "github.com/lukaproject/schedulersvr/apitest/common"
	"github.com/lukaproject/schedulersvr/apitest/swagger"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type TaskTypeTestSuite struct {
	suite.Suite
	cfg        apitest_common.Config
	swaggerCfg *swagger.Configuration

	conn sqlx.SqlConn
}

func (ttts *TaskTypeTestSuite) SetupSuite() {
	conf.MustLoad(*ConfigFile, &ttts.cfg)
	ttts.T().Log(ttts.cfg.Addr)
	ttts.T().Log(ttts.cfg.MysqlDataSource)
	ttts.conn = sqlx.NewMysql(ttts.cfg.MysqlDataSource)
	ttts.swaggerCfg = swagger.NewConfiguration()
	ttts.swaggerCfg.BasePath = ttts.cfg.Addr
}

func (ttts *TaskTypeTestSuite) SetupTest() {
	// clean the before each testcase
	_, err := ttts.conn.Exec("DELETE FROM `task_type`")
	if err != nil {
		panic(err)
	}
}

func (ttts *TaskTypeTestSuite) Test_TaskTypeNotFound() {
	c := swagger.NewAPIClient(ttts.swaggerCfg)
	_, httpResp, _ := c.SchedulerSvrApi.GetTaskType(context.Background(), swagger.GetTaskTypeReq{
		SessionId: uuid.NewString(),
		Name:      "2333",
	})
	require.Equal(ttts.T(), httpResp.StatusCode, http.StatusNotFound)
}

func (ttts *TaskTypeTestSuite) Test_TaskTypeAddAndGetOk() {
	c := swagger.NewAPIClient(ttts.swaggerCfg)

	ttc := swagger.TaskTypeContent{
		Name:              "test_task_type",
		MaxTaskInQueLimit: -1,
		ExtraInfo:         "extra info",
	}

	_, _, err := c.SchedulerSvrApi.AddTaskType(context.Background(), swagger.AddTaskTypeReq{
		SessionId:         uuid.NewString(),
		Name:              ttc.Name,
		MaxTaskInQueLimit: ttc.MaxTaskInQueLimit,
		ExtraInfo:         ttc.ExtraInfo,
	})
	require.Nil(ttts.T(), err)

	resp, _, err := c.SchedulerSvrApi.GetTaskType(context.Background(), swagger.GetTaskTypeReq{
		SessionId: uuid.NewString(),
		Name:      ttc.Name,
	})
	require.Nil(ttts.T(), err)
	require.Equal(ttts.T(), resp.TaskType.ExtraInfo, ttc.ExtraInfo)
	require.Equal(ttts.T(), resp.TaskType.MaxTaskInQueLimit, ttc.MaxTaskInQueLimit)
	require.Equal(ttts.T(), resp.TaskType.Name, ttc.Name)
}

func Test_TaskTypeTestSuite(t *testing.T) {
	suite.Run(t, new(TaskTypeTestSuite))
}
