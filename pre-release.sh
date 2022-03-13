#!/bin/bash

echo ""
echo '-------------------'
echo 'Verify dependencies'
echo '-------------------'
go mod verify

echo ""
echo '-------------------'
echo 'Run go build'
echo '-------------------'
go build -v ./...

echo ""
echo '-------------------'
echo 'Run go vet'
echo '-------------------'
go vet ./...


echo ""
echo '-------------------'
echo 'Install staticcheck'
echo '-------------------'
go install honnef.co/go/tools/cmd/staticcheck@latest


echo ""
echo '-------------------'
echo 'Run staticcheck'
echo '-------------------'
staticcheck ./...


echo ""
echo '-------------------'
echo 'Install golint'
echo '-------------------'
go install golang.org/x/lint/golint@latest


echo ""
echo '-------------------'
echo 'Run golint'
echo '-------------------'
golint ./...


echo ""
echo '-------------------'
echo 'Run tests'
echo '-------------------'
go test -race -vet=off ./...

