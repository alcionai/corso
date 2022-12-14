#!/bin/bash

set -e

usage() {
        echo "Usage: $(basename $0) binary | image [--platforms ...] [--tag ...]"
        echo ""
        echo "OPTIONS"
        echo " -p | --platforms  Platforms to build for (default: $PLATFORMS)"
        echo "                 Specify multiple platforms using ',' (eg: linux/amd64,darwin/arm)"
        echo " -t | --tag        Tag for container image (default: $TAG)"
}

ROOT=$(dirname $(dirname $(readlink -f $0)))
GOVER=1.19                           # go version
CORSO_BUILD_CACHE="/tmp/.corsobuild" # shared persistent cache

# Figure out os and architecture
case "$(uname -m)" in
x86_64) GOARCH="amd64" ;;
aarch64) GOARCH="arm64" ;;
arm64) GOARCH="arm64" ;;
i386) GOARCH="386" ;;
*) echo "Unknown architecture" && exit 0 ;;
esac
case "$(uname)" in
Linux) GOOS="linux" ;;
Darwin) GOOS="darwin" ;;
*) echo "Unknown OS" && exit 0 ;;
esac

PLATFORMS="$GOOS/$GOARCH" # default platform
TAG="alcionai/corso"      # default image tag

MODE="binary"
case "$1" in
binary) MODE="binary" && shift ;;
image)
	MODE="image"
	shift
	GOOS="linux" # darwin container images are not a thing
	;;
-h | --help) usage && exit 0 ;;
*) usage && exit 1 ;;
esac

while [ "$#" -gt 0 ]; do
	case "$1" in
	-p | --platforms) PLATFORMS="$2" && shift ;;
	-t | --tag) TAG="$2" && shift ;;
	*) echo "Invalid argument $1" && usage && exit 1 ;;
	esac
	shift
done

if [ "$MODE" == "binary" ]; then
	mkdir -p ${CORSO_BUILD_CACHE} # prep env
	for platform in ${PLATFORMS/,/ }; do
		IFS='/' read -r -a platform_split <<<"$platform"
		GOOS=${platform_split[0]}
		GOARCH=${platform_split[1]}

		printf "Building for %s...\r" "$platform"
		docker run --rm \
			--mount type=bind,src="${ROOT}",dst="/app" \
			--mount type=bind,src="${CORSO_BUILD_CACHE}",dst="${CORSO_BUILD_CACHE}" \
			--env GOMODCACHE="${CORSO_BUILD_CACHE}/mod" --env GOCACHE="${CORSO_BUILD_CACHE}/cache" \
			--env GOOS=${GOOS} --env GOARCH=${GOARCH} \
			--workdir "/app/src" \
			golang:${GOVER} \
			go build -o corso -ldflags "${CORSO_BUILD_LDFLAGS}"

		OUTFILE="corso"
		[ "$GOOS" == "windows" ] && OUTFILE="corso.exe"

		mkdir -p "${ROOT}/bin/${GOOS}-${GOARCH}"
		mv "${ROOT}/src/corso" "${ROOT}/bin/${GOOS}-${GOARCH}/${OUTFILE}"
		echo Corso $platform binary available in "${ROOT}/bin/${GOOS}-${GOARCH}/${OUTFILE}"
	done
else
	for platform in ${PLATFORMS/,/ }; do
		echo "$platform" | grep -Eq "^darwin" &&
			echo Cannot create darwin images "($platform)" && exit 1
	done
	echo Building "$TAG" image for "$PLATFORMS"
	docker buildx build --tag ${TAG} \
		--platform ${PLATFORMS} \
		--file "${ROOT}/build/Dockerfile" \
		--build-arg CORSO_BUILD_LDFLAGS="$CORSO_BUILD_LDFLAGS" \
		--load "${ROOT}"
	echo Built container image "$TAG"
fi
