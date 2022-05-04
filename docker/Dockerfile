FROM golang:1.18

ARG uid=1000
ARG gid=1000

# Dockerfile with static deps that won't be changed. Right now there's no deps
# outside of what the go.mod file has to manage. Makes a user in the container
# with the same UID and GID as the current user so that the mounted directory
# doesn't end up with files with strange permissions.

RUN groupadd -g "${gid}" build && \
    useradd -m -u "${uid}" -g "${gid}" build

USER build:build
