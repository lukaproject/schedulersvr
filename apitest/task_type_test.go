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
	require.Equal(ttts.T(), ttc.ExtraInfo, resp.TaskType.ExtraInfo)
	require.Equal(ttts.T(), ttc.MaxTaskInQueLimit, resp.TaskType.MaxTaskInQueLimit)
	require.Equal(ttts.T(), ttc.Name, resp.TaskType.Name)
}

func (ttts *TaskTypeTestSuite) Test_TaskTypeAddFailed_BadRequest() {
	c := swagger.NewAPIClient(ttts.swaggerCfg)
	ttc := swagger.TaskTypeContent{
		Name:              "",
		MaxTaskInQueLimit: -1,
		ExtraInfo:         "extra info",
	}
	_, httpResp, err := c.SchedulerSvrApi.AddTaskType(context.Background(), swagger.AddTaskTypeReq{
		SessionId:         uuid.NewString(),
		Name:              ttc.Name,
		MaxTaskInQueLimit: ttc.MaxTaskInQueLimit,
		ExtraInfo:         ttc.ExtraInfo,
	})
	require.NotNil(ttts.T(), err)
	require.Equal(ttts.T(), http.StatusBadRequest, httpResp.StatusCode)
}

func (ttts *TaskTypeTestSuite) Test_TaskTypeAddFailed_Exist() {
	c := swagger.NewAPIClient(ttts.swaggerCfg)
	ttc := swagger.TaskTypeContent{
		Name:              "123",
		MaxTaskInQueLimit: -1,
		ExtraInfo:         "extra info",
	}
	_, httpResp, err := c.SchedulerSvrApi.AddTaskType(context.Background(), swagger.AddTaskTypeReq{
		SessionId:         uuid.NewString(),
		Name:              ttc.Name,
		MaxTaskInQueLimit: ttc.MaxTaskInQueLimit,
		ExtraInfo:         ttc.ExtraInfo,
	})
	require.Nil(ttts.T(), err)
	require.Equal(ttts.T(), http.StatusOK, httpResp.StatusCode)
	_, httpResp, err = c.SchedulerSvrApi.AddTaskType(context.Background(), swagger.AddTaskTypeReq{
		SessionId:         uuid.NewString(),
		Name:              ttc.Name,
		MaxTaskInQueLimit: ttc.MaxTaskInQueLimit,
		ExtraInfo:         ttc.ExtraInfo,
	})
	require.NotNil(ttts.T(), err)
	require.Equal(ttts.T(), http.StatusInternalServerError, httpResp.StatusCode)
}

func (ttts *TaskTypeTestSuite) Test_TaskTypeDeleteSuccess() {
	c := swagger.NewAPIClient(ttts.swaggerCfg)
	ttc := swagger.TaskTypeContent{
		Name:              "123",
		MaxTaskInQueLimit: -1,
		ExtraInfo:         "extra info",
	}
	_, httpResp, err := c.SchedulerSvrApi.AddTaskType(context.Background(), swagger.AddTaskTypeReq{
		SessionId:         uuid.NewString(),
		Name:              ttc.Name,
		MaxTaskInQueLimit: ttc.MaxTaskInQueLimit,
		ExtraInfo:         ttc.ExtraInfo,
	})
	require.Nil(ttts.T(), err)
	require.Equal(ttts.T(), http.StatusOK, httpResp.StatusCode)
	_, httpResp, err = c.SchedulerSvrApi.DeleteTaskType(context.Background(), swagger.DeleteTaskTypeReq{
		SessionId: uuid.NewString(),
		Name:      ttc.Name,
	})
	require.Nil(ttts.T(), err)
	require.Equal(ttts.T(), http.StatusOK, httpResp.StatusCode)

	_, httpResp, err = c.SchedulerSvrApi.GetTaskType(context.Background(), swagger.GetTaskTypeReq{
		SessionId: uuid.NewString(),
		Name:      ttc.Name,
	})
	require.NotNil(ttts.T(), err)
	require.Equal(ttts.T(), http.StatusNotFound, httpResp.StatusCode)
}

func Test_TaskTypeTestSuite(t *testing.T) {
	suite.Run(t, new(TaskTypeTestSuite))
}
