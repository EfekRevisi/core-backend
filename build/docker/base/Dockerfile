################################################################################
# Generate Docker Base Image
# develop
#
# VERSION:
# 0.0.1 - init
################################################################################
FROM golang:1.15

# Enable Go Modules.
ENV GO111MODULE=on

RUN go get  github.com/canthefason/go-watcher/cmd/watcher && \
    mkdir -p /go/src/core-services