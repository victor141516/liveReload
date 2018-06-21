#!/bin/sh

GOOS=windows go build -o build/liveReload.exe liveReload.go ; GOOS=linux go build -o build/liveReload.linux liveReload.go ; GOOS=darwin go build -o build/liveReload.macos liveReload.go
