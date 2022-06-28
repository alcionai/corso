#! /bin/bash

# TODO(ashmrtnz): Needs adjusting based on where source code lands in the repo
# and where the scripts land in the repo.
REPO_ROOT=$(dirname $(pwd))
REPO_CODE="${REPO_ROOT}/src"
GOLANG_BASE_REPO_PATH="/go/src/github.com/alcionai"
GOLANG_REPO_PATH="${GOLANG_BASE_REPO_PATH}/$(basename $REPO_ROOT)"
DEV_TAG="alcionai/dev"
