#!/bin/bash
# Reports unused function parameters and results in your code.

go get mvdan.cc/unparam
go mod vendor
unparam ./...