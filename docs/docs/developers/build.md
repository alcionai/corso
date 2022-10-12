# Building Corso

## Binary

### Building locally

Corso source is contained in `<repo root>/src` directory of the repo. It is a go
project and can be built like any other go project using using `go build` if you have go available in your system:

> You can find instructions to install go [here](https://go.dev/doc/install)

```bash
go build -o corso # run from within `./src` directory
```

This will generate the built binary with the name `corso` in the
directory you built it.

While building, you can pass in a the following values using `ldflags` to configure corso:

- `github.com/alcionai/corso/src/cli.version`: Version reported by `corso` when running `--version`
- `github.com/alcionai/corso/src/internal/events.RudderStackWriteKey`: RudderStack key for metrics
- `github.com/alcionai/corso/src/internal/events.RudderStackDataPlaneURL`: RudderStack data place URL for metrics

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
versions for linux, you can run the following command:

```bash
./build/build.sh --platforms linux/amd64,linux/arm64
```

In order to set `ldflags`, you can make use of `CORSO_BUILD_LDFLAGS`
with the same flags as above. For example, in order specify a version,
you can set the following:

```bash
export CORSO_BUILD_LDFLAGS="-X 'github.com/alcionai/corso/src/cli.version=<version>'"
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
