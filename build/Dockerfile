# syntax=docker/dockerfile:1

# This dockerfile is configured to be run by the /corso/build/build-container.sh
# script.  Using docker to build this file directly will fail.

FROM alpine:3.16

# Add a well-know corso user to use by default
RUN addgroup -g 1001 corso && \
    adduser --shell /sbin/nologin --disabled-password \
    --home /app/corso --uid 1001 --ingroup corso corso

# Update to latest cert bundle in case there are changes
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /

ARG TARGETOS
ARG TARGETARCH
COPY ./bin/${TARGETOS}-${TARGETARCH}/corso ./

USER corso

ENTRYPOINT ["/corso"]
