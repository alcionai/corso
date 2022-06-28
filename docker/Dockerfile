# syntax=docker/dockerfile:1

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY ./bin/corso ./
    
USER nonroot:nonroot

ENTRYPOINT ["/corso"] 