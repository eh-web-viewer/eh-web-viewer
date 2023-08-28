#!/bin/bash

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
# export CC=aarch64-linux-gnu-gcc

go build 
rm back.bin
go build -o back.bin
