#!/bin/bash
# generate the service framework 
goctl api go -api schedulersvr.api -dir .

# generate the swagger json
goctl api plugin -plugin goctl-swagger="swagger -filename schedulersvr.json" -api schedulersvr.api -dir .

# generate the sdk 
# -l "$language"
# -o "output path"
docker run --rm -v \
  "$(pwd):/go-work" swaggerapi/swagger-codegen-cli:2.4.29 generate \
  -i "/go-work/schedulersvr.json" -l "python" -o "/go-work/out/py3"