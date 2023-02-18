#!/bin/bash

# $1 is language
# $2 is output path [$src_root/out/$2]
# ./script/gen-client.sh go go

# -l "$language"
# -o "output path"
docker run --rm -v \
  "$(pwd):/go-work" swaggerapi/swagger-codegen-cli:2.4.29 generate \
  -i "/go-work/schedulersvr.json" -l $1 -o "/go-work/out/$2"

rm -rf ./apitest/swagger/*
cp -r ./out/$2/*.go ./apitest/swagger/ # only copy the go source file