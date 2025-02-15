#!/bin/bash

rm -rf bin
mkdir -vp bin
GOOS=darwin GOARCH=amd64 go build -o bin/go-ec2-instannce-connect-amd64-darwin cmd/go-ec2-instance-connect-ssh/main.go
GOOS=darwin GOARCH=arm64 go build -o bin/go-ec2-instannce-connect-arm64-darwin cmd/go-ec2-instance-connect-ssh/main.go
GOOS=linux GOARCH=amd64 go build -o bin/go-ec2-instannce-connect-amd64-linux cmd/go-ec2-instance-connect-ssh/main.go
GOOS=windows GOARCH=amd64 go build -o bin/go-ec2-instannce-connect-amd64.exe cmd/go-ec2-instance-connect-ssh/main.go

cd bin
tar czvf go-ec2-instance-connect-amd64-darwin.tgz go-ec2-instannce-connect-amd64-darwin
tar czvf go-ec2-instance-connect-arm64-darwin.tgz go-ec2-instannce-connect-arm64-darwin
tar czvf go-ec2-instance-connect-amd64-linux.tgz go-ec2-instannce-connect-amd64-linux
tar czvf go-ec2-instance-connect-amd64-windows.tgz go-ec2-instannce-connect-amd64.exe

