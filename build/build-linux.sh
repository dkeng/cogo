#!/bin/bash
cp ../cogo.toml ../bin/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../bin/cogo ../