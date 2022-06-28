# Corso Dockerfile

## Instructions

The docker image build currently expects the `corso` binary to be available in the same directory.

-  Build the `corso` binary in the build image. See instructions in the `build` directory.

- Copy the `corso` binary to `$REPO/docker/bin` and run `build.sh`

```
$ ./build.sh
```