#! /bin/bash

# Runs the dev docker image with the current user and mounts the source code
# directory in the container at the proper go path.
# NOTE: The container is ephemeral and destroyed after it is exited (but changes
# in the repo's code directory will be available to the host).

source paths.sh

docker run \
  --rm \
  -it \
  -v "$REPO_CODE":"$GOLANG_REPO_PATH" \
  -w "$GOLANG_REPO_PATH" \
  "$DEV_TAG" \
  /bin/bash
