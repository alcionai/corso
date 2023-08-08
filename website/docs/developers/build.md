# Building Corso

## Binary

### Building locally

Corso is a Go project, and you can build it with `go build` from `<repo root>/src`
if you have Go installed on your system.

```bash
# run from within `./src`
go build -o corso 
```

:::info
If you don't have Go available, you can find installation instructions [here](https://go.dev/doc/install)
:::

This will generate a binary named `corso` in the directory where you run the build.

:::note
Prebuilt binary artifacts of the latest commit are available on GitHub.
You can access them by navigating to the "Summary" page of
the [`Build/Release Corso` CI job](https://github.com/alcionai/corso/actions/workflows/ci.yml?query=branch%3Amain)
that was run for the latest commit on the `main` branch.
The downloads will be available in the "Artifacts" section towards the
bottom of the page.
:::

### Building via Docker

For convenience, the Corso build tooling is containerized. To take advantage, you need
[Docker](https://www.docker.com/) installed on your machine.

To build Corso via docker, use the following command from the root of your repo:

```bash
./build/build.sh binary
```

By default, we will build for your current platform. You can pass in
all the architectures/platforms you would like to build it for using
the `--platforms` flag as a comma separated list. For example, if you
would like to build `amd64` and `arm64` versions for Linux, you can
run the following command:

```bash
./build/build.sh binary --platforms linux/amd64,linux/arm64
```

Once built, the resulting binaries will be available in `<repo root>/bin` for all the different platforms you specified.

## Container Image

If you prefer to build Corso as a container image, use the following command instead:

```bash
# Use --help to see all available options
./build/build.sh image
```

:::note
`Dockerfile` used to build the image is available at [`build/Dockerfile`](https://github.com/alcionai/corso/blob/main/build/Dockerfile)
:::

Similar to binaries, we build your a container image for your current
platform by default, but you can change it by explicitly passing in
the platforms that you would like to build for.
In addition, you can optionally pass the tag that you would like to
apply for the image using `--tag` option.

For example, you can use the following command to create a `arm64`
image with the tag `ghcr.io/alcionai/corso:latest`, you can run:

```bash
./build/build.sh image --platforms linux/arm64 --tag ghcr.io/alcionai/corso:latest
```

:::info
If you run into any issues with building cross platform images, make
sure to follow the instructions on [Docker
docs](https://docs.docker.com/build/building/multi-platform/) to setup
the build environment for Multi-platform images.
:::
