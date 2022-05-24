# Corso documentation

Corso documentation uses [Docusaurus 2](https://docusaurus.io/), a modern static website generator.
[Mermaid](https://mermaid-js.github.io/mermaid/) provides support for native diagrams in Markdown.

## Requirements

Developing documentation for Corso requires the following tools on your machine:

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

This command starts a local development server within the Docker container and will expose docs at [http://localhost:3000](http://localhost:3000).

## Building static documentation

```bash
make build
```

This command generates static content into the `build` directory for integration with any static contents hosting service.

## Style and linting

```bash
make check
```

This command will lint all Markdown files and check them for style issues.

## Documentation platform development

```bash
make shell
```

Use this command to interactively (and temporarily!) change the contents or
configuration of the live documentation container image (for example, when
experimenting with new plugins).
