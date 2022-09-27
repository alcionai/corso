# syntax=docker/dockerfile:1

# This dockerfile is configured to be run by the /corso/build/build-container.sh
# script.  Using docker to build this file directly will fail.

FROM ubuntu:latest

WORKDIR /

ARG TARGETOS
ARG TARGETARCH
COPY ./bin/${TARGETOS}-${TARGETARCH}/corso ./

USER nobody

ENTRYPOINT ["/corso"]
