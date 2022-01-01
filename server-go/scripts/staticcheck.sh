#!/bin/bash
# Using static analysis, it finds bugs and performance issues, offers simplifications, and enforces style rules.

go get honnef.co/go/tools/cmd/staticcheck
go mod vendor
staticcheck -checks all ./...
