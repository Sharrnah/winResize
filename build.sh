#!/bin/bash

env GOARCH=amd64 GOOS=windows CGO_ENABLED=1 go build -tags osusergo -v -ldflags "-s -w" -o bin/winResize.exe
