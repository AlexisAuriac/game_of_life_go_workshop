#!/bin/bash

golint -set_exit_status *.go

if [[ $? -ne 0 ]]; then
	exit 1
fi

go build -o gameOfLife *.go
