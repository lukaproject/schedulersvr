# API TEST
The end-to-end test for schedulersvr.

# Usage
you should deploy the redis and mysql first.
```
go test -v [test file] flag.go -args -f apitest.yml
```