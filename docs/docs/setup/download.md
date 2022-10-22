# Downloads

import CodeBlock from '@theme/CodeBlock';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import {Version} from '@site/src/corsoEnv';
import DownloadBinaries from './_download_binary.md';

Corso is available as a Docker container image or as native binaries.

## Native binaries

Corso is available as an `x86_64` and `arm64` executable for Windows, Linux and macOS. These can be downloaded from
the [GitHub releases page](https://github.com/alcionai/corso/releases).

<DownloadBinaries />

## Docker container images

Corso is also available as a Docker container image for Linux (`x86_64` and `arm64`). The image can also be used on
Linux, with Docker Desktop on macOS, and on Windows in
[Linux Mode](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick-start/quick-start-windows-10-linux).
These Docker containers can be pulled from [GitHub's container registry](https://github.com/alcionai/corso/pkgs/container/corso).

We strongly recommend using a container image with the release version tag (for example,
`ghcr.io/alcionai/corso:v0.1.0`) but container images with the `latest` tag are also available. Unreleased builds
with the `nightly` tag are also provided for testing but these are likely to be unstable.

<Tabs groupId="docker">
<TabItem value="release" label="Official Release">

<CodeBlock language="bash">{
`docker pull ghcr.io/alcionai/corso:${Version()}`
}</CodeBlock>

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
