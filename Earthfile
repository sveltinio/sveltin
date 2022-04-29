# Earthfile
FROM golang:1.17.2-alpine3.14

LABEL maintainer="Sveltin contributors <github@sveltin.io>"

WORKDIR /sveltin

RUN apk add --no-cache git

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    ENV GOARCH=amd64
    FROM +deps
    COPY --dir cmd common config helpers resources sveltinlib utils ./
    COPY main.go ./
    RUN go build -o build/sveltin main.go
    SAVE ARTIFACT build/sveltin /sveltin AS LOCAL build/sveltin

#docker:
#    ARG TAG
#    COPY +build/sveltin .
#    ENTRYPOINT ["/sveltin/sveltin"]
#    SAVE IMAGE sveltinio/sveltin:latest
#    SAVE IMAGE sveltinio/sveltin:$TAG
