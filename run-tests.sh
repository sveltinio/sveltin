#!/bin/bash

go test ./... -race -covermode=atomic -coverprofile=coverage.out
