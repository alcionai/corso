#! /bin/bash

# Builds a docker image that contains the deps for the current version of the
# code. Image expects dev directory to be mounted in the container at runtime.

source paths.sh

BASE_TAG="alcionai/base-dev"

buildImage() {
  docker build \
    -f Dockerfile \
    -t "$BASE_TAG" \
    --build-arg uid=$(id -u) \
    --build-arg gid=$(id -g) \
    .
  docker run \
    -v "$REPO_CODE":"$GOLANG_REPO_PATH" \
    --name build-tmp \
    -w "$GOLANG_REPO_PATH" \
    -it \
    "$BASE_TAG" go get
  docker commit build-tmp "$DEV_TAG"
  docker rm build-tmp
}

buildImage
