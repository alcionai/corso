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
  echo "  -a|--arch         Set the architecture to the specified value (default: amd64)"
  echo "  -p|--prefix       Prefixes the image name."
  echo "  -s|--suffix       Suffixes the version."
  echo " "
  echo "-----"
  echo "Example Usage:"
  echo "  ./build/build-container.sh"
  echo "  ./build/build-container.sh --arch arm64"
  echo "  ./build/build-container.sh --arch arm64 --prefix ghcr.io --suffix nightly"
  echo " "
  exit 0
}

SCRIPT_ROOT=$(dirname $(readlink -f $0))
PROJECT_ROOT=$(dirname ${SCRIPT_ROOT})

OS=linux
ARCH=amd64
IMAGE_NAME_PREFIX=
IMAGE_TAG_SUFFIX=

while [ "$#" -gt 0 ]
do
  case "$1" in
  -h|--help)
    usage
    exit 0
    ;;
  -a|--arch)
    ARCH=$2
    shift
    ;;
  -p|--prefix)
    IMAGE_NAME_PREFIX=$2
    shift
    ;;
  -s|--suffix)
    IMAGE_TAG_SUFFIX=$2
    shift
    ;;
  -*)
    echo "Invalid flag '$1'. Use -h|--help to see the valid options" >&2
    return 1
    ;;
  *)
    echo "Invalid arg '$1'. Use -h|--help to see the valid options" >&2
    return 1
    ;;
  esac
  shift
done

IMAGE_TAG=${OS}-${ARCH}
if [ ! -z "${IMAGE_TAG_SUFFIX}" ]; then
  IMAGE_TAG=${IMAGE_TAG}-${IMAGE_TAG_SUFFIX}
fi

IMAGE_NAME=alcionai/corso:${IMAGE_TAG}
if [ ! -z "${IMAGE_NAME_PREFIX}" ]; then
  IMAGE_NAME=${IMAGE_NAME_PREFIX}/${IMAGE_NAME}
fi

${SCRIPT_ROOT}/build.sh --arch ${ARCH}

echo "-----"
echo "building corso container ${IMAGE_NAME}"
echo "-----"

set -x
docker buildx build --tag ${IMAGE_NAME}     \
  --platform ${OS}/${ARCH}                  \
  --file ${PROJECT_ROOT}/build/Dockerfile  \
  ${PROJECT_ROOT}
set +x

echo "-----"
echo "container built successfully"
