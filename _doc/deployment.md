# deployment

## docker

You should install docker and docker-compose first.

1. build scheduler image
```bash
docker build -t schedulersvr:latest -f _docker/Dockerfile .
```
2. deploy
```bash
docker-compose -f _docker/schedulersvrstandalone.yml up -d
```

## Linux

You should download redis first.

1. run the redis

2. edit configuration

Edit the `etc/schedulersvr.yaml`
```ini
Name: SchedulerSvr
Host: 0.0.0.0
Port: 8888
Mode: dev
Log:
  Encoding: plain

Scheduler:
  SchedulerMode: BY_COMMIT_TIME
  MQType: REDIS
  Redis:
    Host: "redis:6379" # you can edit it to the local redis host.
    Type: "node"
    Pass: ""

DB:
  AturKv:
    DirPath: atur/
    Shards: 5
```

3. build shedulersvr
```
go build
```

4. run the schedulersvr
```bash
./schedulersvr -f etc/schedulersvr.yaml
```