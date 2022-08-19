#!/bin/sh

set -e

usage() {
  echo " "
  echo "-----"
  echo "Builds a Corso docker container image."
  echo " "
  echo "-----"
  echo "Flags"
  echo "  -h|--help         Help"
  echo "    |--arm          Set the architecture to arm64 (default: amd64)"
  echo " "
  echo "-----"
  echo "Example Usage:"
  echo "  ./build/build-container.sh"
  echo "  ./build/build-container.sh --arm" 
  echo " "
  exit 0
}

SCRIPT_ROOT=$(dirname $(readlink -f $0))
PROJECT_ROOT=$(dirname ${SCRIPT_ROOT})

OS=linux
ARCH=amd64

while [ "$#" -gt 0 ]
do
  case "$1" in
  -h|--help)
    usage
    exit 0
    ;;
  --arm)
    ARCH=arm64
    ;;
  -*)
    echo "Invalid option '$1'. Use -h|--help to see the valid options" >&2
    return 1
    ;;
  *)
    echo "Invalid option '$1'. Use -h|--help to see the valid options" >&2
    return 1
  ;;
  esac
  shift
done

IMAGE_TAG=${OS}-${ARCH}-$(git describe --tags --always --dirty)
IMAGE_NAME=alcionai/corso:${IMAGE_TAG}

${SCRIPT_ROOT}/build.sh

echo "building container"
set -x
docker buildx build --tag ${IMAGE_NAME}     \
  --platform ${OS}/${ARCH}                  \
  --file ${PROJECT_ROOT}/docker/Dockerfile  \
  ${PROJECT_ROOT}
unset -x
echo "container built successfully ${IMAGE_NAME}"
