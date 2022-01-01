#!/bin/bash
# Detect ineffectual assignments in Go code.

go get github.com/gordonklaus/ineffassign
go mod vendor
ineffassign ./...
