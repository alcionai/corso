# syntax=docker/dockerfile:1

# This dockerfile is configured to be run by the /corso/build/build-container.sh
# script.  Using docker to build this file directly will fail.

FROM gcr.io/distroless/base-debian10
# FROM gcr.io/distroless/base:debug

WORKDIR /

ARG TARGETOS
ARG TARGETARCH
COPY ./bin/${TARGETOS}-${TARGETARCH}/corso ./

USER nonroot:nonroot

ENTRYPOINT ["/corso"]
