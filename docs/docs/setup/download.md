# Downloads

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

Corso is available as a Docker container image or as native binaries.

## Docker container images

The Corso Docker container image is available for Linux (`x86_64` and `arm64`) and this can be used on Linux, with
Docker Desktop on macOS, and on Windows in
[Linux Mode](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick-start/quick-start-windows-10-linux).
These Docker containers can be pulled from [GitHub's container registry](https://github.com/alcionai/corso/pkgs/container/corso).

We strongly recommend using a container image with the release version tag (for example,
`ghcr.io/alcionai/corso:v0.1.0`) but container images with the `latest` tag are also available. Unreleased builds
with the `nightly` tag are also provided for testing but these are likely to be unstable.

<Tabs groupId="docker">
<TabItem value="release" label="Official Release">

  ```bash
  docker pull ghcr.io/alcionai/corso:<version here>
  ```

</TabItem>
<TabItem value="latest" label="Latest">

   ```bash
   docker pull ghcr.io/alcionai/corso:latest
   ```

</TabItem>
<TabItem value="nightly" label="Nightly (Unstable)">

   ```bash
   docker pull ghcr.io/alcionai/corso:nightly
   ```

</TabItem>
</Tabs>

## Native binaries

Corso is also available as an `x86_64` and `arm64` executable for Windows, Linux and macOS. These can be downloaded from
the [GitHub releases page](https://github.com/alcionai/corso/releases).

<Tabs groupId="download">
<TabItem value="win" label="Windows (Powershell)">

  ```powershell
   Invoke-WebRequest `
     -Uri https://github.com/alcionai/corso/releases/download/<VERSION>/corso_<VERSION>_Windows_x86_64.tar.gz `
     -UseBasicParsing -Outfile corso_<VERSION>_Windows_x86_64.tar.gz
   tar zxvf .\corso_<VERSION>_Windows_x86_64.tar.gz
  ```

</TabItem>
<TabItem value="linux-arm" label="Linux - arm64">

   ```bash
   curl -L -O https://github.com/alcionai/corso/releases/download/<VERSION>/corso_<VERSION>_Linux_arm64.tar.gz && \
     tar zxvf corso_<VERSION>_Linux_arm.tar.gz
   ```

</TabItem>
<TabItem value="linux-x86-64" label="Linux - x86_64">

   ```bash
   curl -L -O https://github.com/alcionai/corso/releases/download/<VERSION>/corso_<VERSION>_Linux_x86_64.tar.gz && \
     tar zxvf corso_<VERSION>_Linux_x86_64.tar.gz
   ```

</TabItem>
<TabItem value="macos-arm" label="macOS - arm64">

   ```bash
   curl -L -O https://github.com/alcionai/corso/releases/download/<VERSION>/corso_<VERSION>_Darwin_arm64.tar.gz && \
     tar zxvf corso_<VERSION>_Darwin_arm.tar.gz
   ```

</TabItem>
<TabItem value="macos-x86-64" label="macOS - x86_64">

   ```bash
   curl -L -O https://github.com/alcionai/corso/releases/download/<VERSION>/corso_<VERSION>_Darwin_x86_64.tar.gz && \
     tar zxvf corso_<VERSION>_Darwin_x86_64.tar.gz
   ```

</TabItem>
</Tabs>
