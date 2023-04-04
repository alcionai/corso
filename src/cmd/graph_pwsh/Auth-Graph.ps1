[CmdletBinding()]

Param (
    [Parameter(Mandatory = $False, HelpMessage = "Use beta Graph API")]
    [bool]$UseBeta

)

$tenantId = $ENV:AZURE_TENANT_ID
$clientId = $ENV:AZURE_CLIENT_ID
$clientSecret = $ENV:AZURE_CLIENT_SECRET

# This version of Graph Powershell does not support app secret auth yet so roll our own
$body = @{
    Grant_Type    = "client_credentials"
    Scope         = "https://graph.microsoft.com/.default"
    Client_Id     = $clientId
    Client_Secret = $clientSecret
}

$ConectionRequest = @{
    Uri    = "https://login.microsoftonline.com/$tenantId/oauth2/v2.0/token"
    Method = "POST"
    Body   = $body
}

$connection = Invoke-RestMethod @ConectionRequest

Write-Host "Authenticating with tenantId: $tenantId ..."
try {
    Connect-MgGraph -AccessToken $connection.access_token 
    Write-Host "Successfully authenticated with tenantId: $tenantId ..."
}
catch {
    Write-Host "Authentication failed..."
    Write-Output $_
}

if ($UseBeta) {
    Write-Host "Switching to Beta Graph API..."
    Select-MgProfile -Name "beta"
}




