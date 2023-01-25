#!/bin/bash
goctl api go -api schedulersvr.api -dir .
goctl api plugin -plugin goctl-swagger="swagger -filename schedulersvr.json" -api schedulersvr.api -dir .