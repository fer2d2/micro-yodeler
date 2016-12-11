#!/bin/bash

rm -f builds/*

env GOOS=linux GOARCH=amd64 go build -o builds/yodel.linux.amd64 main.go
env GOOS=linux GOARCH=arm go build -o builds/yodel.linux.arm main.go
env GOOS=darwin GOARCH=amd64 go build -o builds/yodel.osx.amd64 main.go
