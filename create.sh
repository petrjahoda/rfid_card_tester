#!/usr/bin/env bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o run_me_linux
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o run_me_windows.exe
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o run_me_mac
