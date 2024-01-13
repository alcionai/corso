$tenantId = $ENV:AZURE_TENANT_ID
$clientId = $ENV:AZURE_CLIENT_ID
$clientSecret = $ENV:AZURE_CLIENT_SECRET
$useBeta = ($ENV:MSGRAPH_USE_BETA -eq 1) -or ($ENV:MSGRAPH_USE_BETA -eq "1") -or ($ENV:MSGRAPH_USE_BETA -eq "true")

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

if ($useBeta) {
    Write-Host "Switching to Beta Graph API..."
    Select-MgProfile -Name "beta"
}




