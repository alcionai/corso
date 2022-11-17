# Corso website documentation

## Requirements

Building the Corso website requires the following tools on your machine:

- `make`
- Docker

## Installation

```bash
make buildimage
```

## Live documentation development

```bash
make dev
```

This command starts a local development server within the Docker container and will expose a live website preview at [http://localhost:5050](http://localhost:5050).

## Building a static website

```bash
make build
```

This command generates static content into the `dist` directory for integration with any static contents hosting service. If you are using AWS S3 + CloudFront, you can run `make publish` to upload to the configured S3 bucket.

## Website platform development

```bash
make shell
```

Use this command to interactively (and temporarily!) change the contents or
configuration of the live website container image (for example, when
experimenting with new packages).
