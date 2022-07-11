#!/bin/sh

set -ex

SCRIPT_ROOT=$(dirname $(readlink -f $0))
PROJECT_ROOT=$(dirname ${SCRIPT_ROOT})

IMAGE_TAG=$(git describe --tags --always --dirty)
IMAGE_NAME=alcionai/corso:${IMAGE_TAG}

${SCRIPT_ROOT}/build.sh

echo "building container"
docker build -t ${IMAGE_NAME}            \
       -f ${PROJECT_ROOT}/Dockerfile     \
       ${PROJECT_ROOT}

echo "container built successfully ${IMAGE_NAME}"
