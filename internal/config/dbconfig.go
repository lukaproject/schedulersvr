package config

import "github.com/zeromicro/go-zero/core/logx"

type DbConfig struct {
	AturKv AturKvConfig
}

type AturKvConfig struct {
	DirPath string
	Shards  int
}

func (dbc DbConfig) Tell() {
	logx.Infof("atur.DirPath=%s", dbc.AturKv.DirPath)
	logx.Infof("atur.Shards=%d", dbc.AturKv.Shards)
}
