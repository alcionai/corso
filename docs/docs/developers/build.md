# Building Corso


For convenience, the Corso build tooling is containerized. To take advantage, you need the following on your machine:

- Docker


To build Corso locally, use the following command from the root of your repo:

```bash
./build/build.sh 

```

The resulting binary will be under `<repo root>/bin`

If you prefer to build Corso as a container, use the following command:

```bash
# Use --help to see all available options
./build/build-container.sh 
```
