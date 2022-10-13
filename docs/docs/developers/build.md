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
If you dong have Go available, you can find installation instructions [here](https://go.dev/doc/install)
:::

This will generate a binary named `corso` in the directory where you run the build.

### Building via Docker

For convenience, the Corso build tooling is containerized. To take advantage, you need
[Docker](https://www.docker.com/) installed on your machine.

To build Corso via docker, use the following command from the root of your repo:

```bash
./build/build.sh
```

You can pass in all the architectures/platforms you would like to
build it for using the `--platforms` flag as a comma separated
list. For example, if you would like to build `amd64` and `arm64`
versions for Linux, you can run the following command:

```bash
./build/build.sh --platforms linux/amd64,linux/arm64
```

Once built, the resulting binaries will be available in `<repo root>/bin` for all the different platforms you specified.

## Container Image

If you prefer to build Corso as a container image, use the following command instead:

```bash
# Use --help to see all available options
./build/build-container.sh
```

Below are the main customization flags that you can set when building a container image:

- `-a|--arch`: Set the architecture to the specified value. (default: amd64)
- `-l|--local`: Build the corso binary on your local system, rather than a go image.
- `-p|--prefix`: Prefix for the image name.
- `-s|--suffix`: Suffix for the version.

For example, you can use the following command to create a `arm64` image with prefix of `ghcr.io` and the tag as `nightly`.

```bash
./build/build-container.sh --arch arm64 --prefix ghcr.io --suffix nightly
```
