#!/bin/bash

go run incs.go --service :1111 --data-path /tmp/1111 &
go run incs.go --service :2222 --data-path /tmp/2222 &
go run incs.go --service :3333 --data-path /tmp/3333 &
go run incs.go --service :4444 --data-path /tmp/4444 &

wait
