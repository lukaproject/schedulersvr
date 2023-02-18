#!/bin/bash
# generate the swagger json
goctl api plugin -plugin goctl-swagger="swagger -filename schedulersvr.json" -api schedulersvr.api -dir .