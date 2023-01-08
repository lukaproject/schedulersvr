#!/bin/bash
goctl api go -api schedulersvr.api -dir .
goctl model mysql ddl -src="_sql/*.sql" -dir="./db/model"