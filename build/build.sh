#!/bin/bash

set -e

SCRIPT_ROOT=$(dirname $(readlink -f $0))
PROJECT_ROOT=$(dirname ${SCRIPT_ROOT})

CORSO_BUILD_CONTAINER=/go/src/github.com/alcionai/corso
CORSO_BUILD_CONTAINER_SRC=${CORSO_BUILD_CONTAINER}/src
CORSO_BUILD_PKG_MOD=/go/pkg/mod
CORSO_BUILD_TMP=/tmp/.corsobuild
CORSO_BUILD_TMP_CACHE=${CORSO_BUILD_TMP}/cache
CORSO_BUILD_TMP_MOD=${CORSO_BUILD_TMP}/mod
CORSO_CACHE=${CORSO_BUILD_TMP_CACHE}
CORSO_MOD_CACHE=${CORSO_BUILD_PKG_MOD}/cache

CORSO_BUILD_ARGS=''

platforms=
GOVER=1.18
GOOS=linux
GOARCH=amd64

while [ "$#" -gt 0 ]
do
  case "$1" in
  --platforms)
    platforms=$2
    shift
    ;;
  esac
  shift
done

# temporary directory for caching go build
mkdir -p ${CORSO_BUILD_TMP_CACHE}
# temporary directory for caching go modules (needed for fast cross-platform build)
mkdir -p ${CORSO_BUILD_TMP_MOD}

if [ -z "$platforms" ]; then
  platforms="${GOOS}/${GOARCH}"
fi

for platform in ${platforms/,/ }
do
  IFS='/' read -r -a platform_split <<< "${platform}"
  GOOS=${platform_split[0]}
  GOARCH=${platform_split[1]}

  echo "-----"
  echo "building corso binary for ${GOOS}/${GOARCH}"
  echo "-----"

  set -x
  docker run --rm \
    --mount type=bind,src=${PROJECT_ROOT},dst=${CORSO_BUILD_CONTAINER}          \
    --mount type=bind,src=${CORSO_BUILD_TMP_CACHE},dst=${CORSO_BUILD_TMP_CACHE} \
    --mount type=bind,src=${CORSO_BUILD_TMP_MOD},dst=${CORSO_BUILD_PKG_MOD}     \
    --workdir ${CORSO_BUILD_CONTAINER_SRC}                                      \
    --env GOMODCACHE=${CORSO_MOD_CACHE}                                         \
    --env GOCACHE=${CORSO_CACHE}                                                \
    --env GOOS=${GOOS}                                                          \
    --env GOARCH=${GOARCH}                                                      \
    --entrypoint /usr/local/go/bin/go                                           \
    golang:${GOVER}                                                             \
    build -o corso ${CORSO_BUILD_ARGS}
  set +x

  mkdir -p ${PROJECT_ROOT}/bin/${GOOS}-${GOARCH}
  mv ${PROJECT_ROOT}/src/corso ${PROJECT_ROOT}/bin/${GOOS}-${GOARCH}/corso

  echo "-----"
  echo "created binary image in ${PROJECT_ROOT}/bin/${GOOS}-${GOARCH}/corso"
  echo "-----"
done