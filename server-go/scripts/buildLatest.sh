#!/bin/bash
# build the includes applications.

LATEST=$(git describe --tags)

echo "your current state will be stashed"
git stash

git checkout ${LATEST}
git reset --hard

echo "compile server @ ${LATEST}"
go build -o ./bin/server-"$LATEST" ./cmd/server/server.go
echo "done!"

echo "compile configurator @ ${LATEST}"
go build -o ./bin/configurator-"$LATEST" ./cmd/configurator/configurator.go
echo "done!"

echo "finished!"
