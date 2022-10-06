<img src="https://github.com/alcionai/corso/blob/main/docs/static/img/corso_logo.svg?raw=true" alt="Corso Logo" width="100" />

# Corso

[![Discord](https://img.shields.io/badge/discuss-discord-blue)](https://discord.gg/63DTTSnuhT)
[![License](https://img.shields.io/badge/License-Apache_2.0-green.svg)](https://opensource.org/licenses/Apache-2.0)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

Corso is the first open-source tool that aims to assist IT admins with the critical task of protecting their
Microsoft 365 data. It provides a reliable, secure, and efficient data protection engine. Admins decide where to store
the backup data and have the flexibility to perform backups of their desired service through an intuitive interface.
As Corso evolves, it can become a great building block for more complex data protection workflows.

**Corso is currently in ALPHA and should NOT be used in production.**

Corso supports M365 Exchange and OneDrive with SharePoint and Teams support in active development. Coverage for more
services, possibly beyond M365, will expand based on the interest and needs of the community.

## Getting Started

TODO - Link to the appropriate page in the published docs.

## Building Corso

```sh
# Build a binary. Will be placed in bin/
./build/build.sh 

# Build a container image
./build/build-container.sh 
```

# Containers

Corso images are hosted on [ghrc.io](https://github.com/alcionai/corso/pkgs/container/corso).

Rolling release
```sh
docker pull ghcr.io/alcionai/corso:{SHA} --platform linux/arm64
```

# Contribution Guidelines

TODO 

## Code of Conduct

It's important that our community is inclusive and respectful of everyone.
We ask that all Corso users and contributors take a few minutes to review our
[Code of Conduct](CODE_OF_CONDUCT.md)

## License

TODO
`
