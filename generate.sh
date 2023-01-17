#!/bin/bash
goctl api go -api schedulersvr.api -dir .
goctl model mysql ddl -src="./_sql/*.sql" -dir="./db/model"
goctl api plugin -plugin goctl-swagger="swagger -filename schedulersvr.json" -api schedulersvr.api -dir .