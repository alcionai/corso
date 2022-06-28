#! /bin/bash

# Builds a docker image that wraps the corso binary

IMAGE_NAME="alcionai/corso"
VERSION=$(git describe --tags --always --dirty)
CORSO_BINARY="./bin/corso"

if [ ! -f "$CORSO_BINARY" ]; then
    echo "$CORSO_BINARY does not exist. Build corso and ensure the binary is available at $CORSO_BINARY"
    exit 1
fi

buildImage() {
  docker build . \
    -t "$IMAGE_NAME:$VERSION"
}

buildImage