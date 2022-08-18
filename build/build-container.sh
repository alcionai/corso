#!/bin/sh

set -ex

SCRIPT_ROOT=$(dirname $(readlink -f $0))
PROJECT_ROOT=$(dirname ${SCRIPT_ROOT})

OS=linux
ARCH=amd64

IMAGE_TAG=${OS}-${ARCH}-$(git describe --tags --always --dirty)
IMAGE_NAME=alcionai/corso:${IMAGE_TAG}

${SCRIPT_ROOT}/build.sh

echo "building container"
docker buildx build --tag ${IMAGE_NAME}          \
       --platform ${OS}/${ARCH}                  \
       --file ${PROJECT_ROOT}/docker/Dockerfile  \
       ${PROJECT_ROOT}

echo "container built successfully ${IMAGE_NAME}"
