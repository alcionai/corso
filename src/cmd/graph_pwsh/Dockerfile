from m365pnp/powershell:2.1.1-alpine-3.14

RUN Install-Module PowerShellGet -Force
RUN Install-Module Microsoft.Graph -Force -RequiredVersion 1.25.0 -Scope AllUsers

COPY ./Auth-Graph.ps1 /tmp/Auth-Graph.ps1
RUN Move-Item -Path /tmp/Auth-Graph.ps1 -Destination $PROFILE.AllUsersAllHosts

WORKDIR /usr/pwsh