# Corso build image

The Corso build image is a docker image that
can be used to dev/test/build Corso

## Creating the build image

Run `build.sh` to create `docker.io/alcionai/base-dev`
```
$ ./build.sh
```

## Launching the build image

Run `run_dev.sh` to launch into the build image. It will mount the corso src directory at `/`
```
$ ./run_dev.sh
```