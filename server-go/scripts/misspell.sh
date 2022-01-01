#!/bin/bash
# Correct commonly misspelled English words... quickly.

go get github.com/client9/misspell/cmd/misspell
go mod vendor
misspell \
  -locale GB \
  -error \
  *.md *.go
