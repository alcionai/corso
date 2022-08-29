#!/bin/bash

set -e

SCRIPT_ROOT=$(dirname $(readlink -f $0))
PROJECT_ROOT=$(dirname ${SCRIPT_ROOT})

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

CORSO_BUILD_ARGS="$@"

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

  OS_ARCH_DIR=${PROJECT_ROOT}/bin/${GOOS}-${GOARCH}

  set -x

  mkdir -p ${OS_ARCH_DIR}

  cd ${PROJECT_ROOT}/src; \
  GOOS=${GOOS}            \
  GOARCH=${GOARCH}        \
  go build -o ${OS_ARCH_DIR} "$CORSO_BUILD_ARGS"

  set +x

  echo "-----"
  echo "created binary ${PROJECT_ROOT}/bin/${GOOS}-${GOARCH}/corso"
  echo "-----"
done