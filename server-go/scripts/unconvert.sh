#!/bin/bash
# The unconvert program analyzes Go packages to identify unnecessary type conversions.

go get github.com/mdempsky/unconvert
go mod vendor
unconvert -v ./...