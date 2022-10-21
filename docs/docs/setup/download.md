# Downloads

Corso is available as a Docker container image or as native binaries.

## Docker container images

The Corso Docker container image is available for Linux (`x86_64` and `arm64`) and this can be used on Linux, with
Docker Desktop on macOS, and on Windows in
[Linux Mode](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick-start/quick-start-windows-10-linux).
These Docker containers can be pulled from [GitHub's container registry](https://github.com/alcionai/corso/pkgs/container/corso).

We strongly recommend using a container image with the release version tag (for example,
`ghcr.io/alcionai/corso:v0.1.0`) but container images with the `latest` tag are also available. Unreleased builds
with the `nightly` tag are also provided for testing but these are likely to be unstable.

## Native binaries

Corso is also available as an `x86_64` and `arm64` executable for Windows, Linux and macOS. These can be downloaded from
the [GitHub releases page](https://github.com/alcionai/corso/releases).
