package config

import "github.com/zeromicro/go-zero/core/logx"

type DbConfig struct {
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Datasource string
}

func (dbc DbConfig) Tell() {
	logx.Infof("mysql.datasource=%s", dbc.Mysql.Datasource)
}
