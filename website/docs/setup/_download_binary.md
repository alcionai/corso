<!-- markdownlint-disable MD034 MD041 -->
<!-- vale Vale.Spelling = NO -->

import CodeBlock from '@theme/CodeBlock';
import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import {Version} from '@site/src/corsoEnv';

<Tabs groupId="download">
<TabItem value="win" label="Windows (Powershell)">

<CodeBlock language="powershell">{
`$ProgressPreference = 'SilentlyContinue'
Invoke-WebRequest \`
  -Uri https://github.com/alcionai/corso/releases/download/${Version()}/corso_${Version()}_Windows_x86_64.zip \`
  -UseBasicParsing -Outfile corso_${Version()}_Windows_x86_64.zip
Expand-Archive .\\corso_${Version()}_Windows_x86_64.zip`
}</CodeBlock>

</TabItem>
<TabItem value="linux-arm" label="Linux - arm64">

<CodeBlock language="bash">{
`curl -L -O https://github.com/alcionai/corso/releases/download/${Version()}/corso_${Version()}_Linux_arm64.tar.gz && \\
  tar zxvf corso_${Version()}_Linux_arm64.tar.gz`
}</CodeBlock>

</TabItem>
<TabItem value="linux-x86-64" label="Linux - x86_64">

<CodeBlock language="bash">{
`curl -L -O https://github.com/alcionai/corso/releases/download/${Version()}/corso_${Version()}_Linux_x86_64.tar.gz && \\
  tar zxvf corso_${Version()}_Linux_x86_64.tar.gz`
}</CodeBlock>

</TabItem>
<TabItem value="macos-arm" label="macOS - arm64">

<CodeBlock language="bash">{
`curl -L -O https://github.com/alcionai/corso/releases/download/${Version()}/corso_${Version()}_Darwin_arm64.tar.gz && \\
  tar zxvf corso_${Version()}_Darwin_arm64.tar.gz`
}</CodeBlock>

</TabItem>
<TabItem value="macos-x86-64" label="macOS - x86_64">

<CodeBlock language="bash">{
`curl -L -O https://github.com/alcionai/corso/releases/download/${Version()}/corso_${Version()}_Darwin_x86_64.tar.gz && \\
  tar zxvf corso_${Version()}_Darwin_x86_64.tar.gz`
}</CodeBlock>

</TabItem>
</Tabs>

<!-- vale Vale.Spelling = YES -->
<!-- markdownlint-enable MD034 MD041 -->
