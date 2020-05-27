#!/bin/sh

if [ $# != 1 ]; then
	echo "Usage: $0 [binary name]"
	exit 0
fi

GOOS=linux GOARCH=amd64 go build -o ./bin/$1_linux64
GOOS=linux GOARCH=386 go build -o ./bin/$1_linux386

GOOS=windows GOARCH=386 go build -o ./bin/$1_windows386.exe
GOOS=windows GOARCH=amd64 go build -o ./bin/$1_windows64.exe

GOOS=darwin GOARCH=386 go build -o ./bin/$1_darwin386
GOOS=darwin GOARCH=amd64 go build -o ./bin/$1_darwin64
