# Corso Documentation

Corso documentation is built using [Docusaurus 2](https://docusaurus.io/), a modern static website generator.

### Installation

```
$ make buildimage
```

### Live Docs

```
$ make dev
```

This command starts a local development server within the Docker container and will expose docs at http://localhost:3000.

### Build Docs

```
$ make build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Check Docs

```
$ make check
```

This command will lint all Markdown files and check them for style issues.

### Documentation Platform Development


```
$ make shell
```

This command is when you want to interactively (and temporarily!) change the contents or configuration of the live documentation container image.
