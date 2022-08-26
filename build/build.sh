#!/bin/sh

set -e

SCRIPT_ROOT=$(dirname $(readlink -f $0))
PROJECT_ROOT=$(dirname ${SCRIPT_ROOT})
SRC_DIR=${PROJECT_ROOT}
CORSO_BUILD_ARGS=''

CORSO_BUILD_CONTAINER_DIR=/go/src/github.com/alcionai/corso
CORSO_BUILD_CONTAINER_SRC_DIR=${CORSO_BUILD_CONTAINER_DIR}/src

GOOS=linux
GOARCH=amd64

while [ "$#" -gt 0 ]
do
  case "$1" in
  --arch)
    GOARCH=$2
    shift
    ;;
  esac
  shift
done

# temporary directory for caching go build
mkdir -p /tmp/.corsobuild/cache
# temporary directory for caching go modules (needed for fast cross-platform build)
mkdir -p /tmp/.corsobuild/mod

echo "building corso"
set -x
docker run --rm --mount type=bind,src=${SRC_DIR},dst=${CORSO_BUILD_CONTAINER_DIR}    \
       --mount type=bind,src=/tmp/.corsobuild/cache,dst=/tmp/.corsobuild/cache       \
       --mount type=bind,src=/tmp/.corsobuild/mod,dst=/go/pkg/mod                    \
       --workdir ${CORSO_BUILD_CONTAINER_SRC_DIR}                                    \
       --env GOCACHE=/tmp/.corsobuild/cache                                          \
       --env GOOS=${GOOS}                                                            \
       --env GOARCH=${GOARCH}                                                        \
       --env GOCACHE=/tmp/.corsobuild/cache                                          \
       --entrypoint /usr/local/go/bin/go                                             \
       golang:1.18                                                                   \
       build ${CORSO_BUILD_ARGS}

mkdir -p ${PROJECT_ROOT}/bin
set +x

echo "creating binary image in bin/corso"
mv ${PROJECT_ROOT}/src/corso ${PROJECT_ROOT}/bin/corso
