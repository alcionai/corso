FROM golang:1.19-alpine as builder

WORKDIR /go/src/app
COPY src .

ARG CORSO_BUILD_LDFLAGS=""
RUN go build -o corso -ldflags "$CORSO_BUILD_LDFLAGS"

FROM alpine:3.17

LABEL org.opencontainers.image.title="Corso"
LABEL org.opencontainers.image.description="Free, Secure, and Open-Source Backup for Microsoft 365"
LABEL org.opencontainers.image.url="https://github.com/alcionai/corso"
LABEL org.opencontainers.image.source="https://github.com/alcionai/corso"
LABEL org.opencontainers.image.vendor="Alcion, Inc."

COPY --from=builder /go/src/app/corso /corso

#  Pull tls certs directly from latest upstream image
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV CORSO_HOME=/app/corso
ENV CORSO_CONFIG_DIR=$CORSO_HOME \
	KOPIA_CONFIG_PATH=$CORSO_HOME/kopia/config/repository.config \
	KOPIA_LOG_DIR=$CORSO_HOME/kopia/logs \
	KOPIA_CACHE_DIRECTORY=$CORSO_HOME/kopia/cache \
	RCLONE_CONFIG=$CORSO_HOME/kopia/rclone/rclone.conf \
	KOPIA_PERSIST_CREDENTIALS_ON_CONNECT=false \
	KOPIA_CHECK_FOR_UPDATES=false

ENTRYPOINT ["/corso"]