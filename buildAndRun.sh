#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
rm -f vlibrary-linux.tar.gz
rm -rf vlibrary-linux
mkdir vlibrary-linux
cp main vlibrary-linux/cintool
cp -R views vlibrary-linux/
cp -R static vlibrary-linux/
cp -R conf vlibrary-linux/
tar -zcf vlibrary-linux.tar.gz vlibrary-linux
rm -rf vlibrary-linux