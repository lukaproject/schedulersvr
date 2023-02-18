package apitest

import (
	"context"
	"testing"

	"github.com/google/uuid"
	apitest_common "github.com/lukaproject/schedulersvr/apitest/common"
	"github.com/lukaproject/schedulersvr/apitest/swagger"
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
	conf.MustLoad(*configFile, &ttts.cfg)
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
	resp, httpResp, err := c.SchedulerSvrApi.GetTaskType(context.Background(), swagger.GetTaskTypeReq{
		SessionId: uuid.NewString(),
		Name:      "2333",
	})
	ttts.T().Log(err)
	ttts.T().Log(httpResp)
	ttts.T().Log(resp)
}

func Test_TaskTypeTestSuite(t *testing.T) {
	suite.Run(t, new(TaskTypeTestSuite))
}
