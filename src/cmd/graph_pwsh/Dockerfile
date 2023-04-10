from m365pnp/powershell:2.1.1-alpine-3.14

RUN Install-Module PowerShellGet -Force
RUN Install-Module Microsoft.Graph -Force -RequiredVersion 1.25.0 -Scope AllUsers

COPY ./Execute-Script.ps1 /tpm/Execute-Script.ps1
RUN Move-Item -Path /tpm/Execute-Script.ps1 -Destination $ENV:PATH.Split(":")[0]

# If you add other powershell commands this needs to be last since it sets 
# the profile to run auth for the user which will then try to run for subsequent connads
COPY ./Auth-Graph.ps1 /tmp/Auth-Graph.ps1
RUN Move-Item -Path /tmp/Auth-Graph.ps1 -Destination $PROFILE.AllUsersAllHosts

WORKDIR /usr/pwsh