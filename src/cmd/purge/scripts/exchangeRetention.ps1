# This is tested on Mac as well as Docker (with m365pnp/powershell image)
# To run in Docker with the script in the current working diredctory 
# docker run --rm -it -v "$(pwd):/usr/reset-retnention" -e M365TENANT_ADMIN_USER -e M365TENANT_ADMIN_PASSWORD \
#                     -w /usr/reset-retnention m365pnp/powershell pwsh -c "./setRetention.ps1"
Param (
    [Parameter(Mandatory = $False, HelpMessage = "Exchange Admin email")]
    [String]$AdminUser = $ENV:M365_TENANT_ADMIN_USER,

    [Parameter(Mandatory = $False, HelpMessage = "Exchange Admin password")]
    [String]$AdminPwd = $ENV:M365_TENANT_ADMIN_PASSWORD
)

# Setup ExchangeOnline
if (-not (Get-Module -ListAvailable -Name ExchangeOnlineManagement)) {
    $ProgressPreference = 'SilentlyContinue'
    Install-Module -Name ExchangeOnlineManagement -MinimumVersion 3.0.0 -Force
    $ProgressPreference = 'Continue'
}

Write-Host "`nConnecting to Exchange..."
$password = convertto-securestring -String "$AdminPwd" -AsPlainText -Force
$cred = New-Object -TypeName System.Management.Automation.PSCredential -ArgumentList $AdminUser, $password
Connect-ExchangeOnline -Credential $cred

Write-Host "`nResetting retention..."

# Set retention values for all mailboxes 
Get-Mailbox | ForEach-Object {
    Write-Host "...for" $_
    Set-Mailbox -Identity $_.Alias `
        -RetentionHoldEnabled $false `
        -LitigationHoldEnabled $false `
        -SingleItemRecoveryEnabled $false `
        -RetainDeletedItemsFor 0 `
        -AuditLogAgeLimit 0 `
        -Force
}
