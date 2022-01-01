#!/bin/bash
# build the includes applications.

BRANCH=$(git branch --show-current)
STAMP=$(date +"%s")

echo "compile server ..."
go build -o ./bin/server-"$BRANCH"-"$STAMP" ./cmd/server/server.go
echo "done!"

echo "compile configurator ..."
go build -o ./bin/configurator-"$BRANCH"-"$STAMP" ./cmd/configurator/configurator.go
echo "done!"

echo "finished!"
