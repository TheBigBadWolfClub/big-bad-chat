#!/bin/bash
# errcheck is a program for checking for unchecked errors in go programs.

go get github.com/kisielk/errcheck
go mod vendor
errcheck ./...