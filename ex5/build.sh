#!/bin/bash

if [[ $1 = "lint" ]]; then
	golint -set_exit_status *.go

	if [[ $? -ne 0 ]]; then
		exit 1
	fi
fi

go build -o gameOfLife *.go
