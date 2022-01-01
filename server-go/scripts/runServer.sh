#!/bin/bash
# run server

go run ./cmd/server/server.go -conf="./test/conf.ucapiv1-dev.json" -debuglog="./debug.log"