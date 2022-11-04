[CmdletBinding(SupportsShouldProcess)]

Param (
    [Parameter(Mandatory = $True, HelpMessage = "User for which to delete folders")]
    [String]$User,

    [Parameter(Mandatory = $False, HelpMessage = "Well-known name of folder under which to clean")]
    [String]$WellKnownRoot = "deleteditems",

    [Parameter(Mandatory = $False, HelpMessage = "Purge folders before this date time (UTC)")]
    [datetime]$FolderBeforePurge,

    [Parameter(Mandatory = $False, HelpMessage = "Name of specific folder to purge under root")]
    [String]$FolderNamePurge,

    [Parameter(Mandatory = $False, HelpMessage = "Purge folders with this prefix")]
    [String]$FolderPrefixPurge,

    [Parameter(Mandatory = $False, HelpMessage = "Azure TenantId")]
    [String]$TenantId = $ENV:AZURE_TENANT_ID,

    [Parameter(Mandatory = $False, HelpMessage = "Azure ClientId")]
    [String]$ClientId = $ENV:AZURE_CLIENT_ID,

    [Parameter(Mandatory = $False, HelpMessage = "Azure ClientSecret")]
    [String]$ClientSecret = $ENV:AZURE_CLIENT_SECRET
)

function Get-AccessToken {
    [CmdletBinding()]
    Param()

    if ([String]::IsNullOrEmpty($TenantId) -or [String]::IsNullOrEmpty($ClientId) -or [String]::IsNullOrEmpty($ClientSecret)) {
        Write-Host "Need to specify TenantId, ClientId, and ClientSecret as parameters or ENVs"
    }

    $body=@{
        client_id=$ClientId
        client_secret=$ClientSecret
        scope="https://outlook.office365.com/.default"
        grant_type="client_credentials"
    }
    
    $res = Invoke-WebRequest -Uri "https://login.microsoftonline.com/$TenantId/oauth2/v2.0/token" -ContentType "application/x-www-form-urlencoded" -Body $body -Method Post

    return $res.content | ConvertFrom-Json | Select-Object -ExpandProperty access_token
}

function Initialize-SOAPMessage {
    [CmdletBinding()]
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "User for which to delete folders")]
        [String]$User,

        [Parameter(Mandatory = $True, HelpMessage = "The message body")]
        [String]$Body
    )

    $Message = @"
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
 xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" 
 xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types"
 xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">

    <soap:Header>
        <t:ExchangeImpersonation>
            <t:ConnectingSID>
                <t:PrimarySmtpAddress>$User</t:PrimarySmtpAddress>
            </t:ConnectingSID>
        </t:ExchangeImpersonation>
    </soap:Header>
    
    <soap:Body>
        $Body
    </soap:Body>
</soap:Envelope>
"@

    return $Message
}

function Invoke-SOAPRequest {
    [CmdletBinding()]
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "OAuth token to connect to Exchange Online.")]
        [securestring]$Token,

        [Parameter(Mandatory = $True, HelpMessage = "The message ")]
        [String]$Message,

        [Parameter(Mandatory = $False, HelpMessage = "The http method")]
        [String]$Method = "Post"
    )
    
    # EWS service url for Exchange Online
    $webServiceUrl = "https://outlook.office365.com/EWS/Exchange.asmx"
    
    $Response = Invoke-WebRequest -Uri $webServiceUrl -Authentication Bearer -Token $Token -Body $Message -Method $Method 
    [xml]$xmlResponse = $Response.Content

    return $xmlResponse
}

function Remove-Folder {
    [CmdletBinding(SupportsShouldProcess)]
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "OAuth token to connect to Exchange Online.")]
        [Securestring]$Token,

        [Parameter(Mandatory = $True, HelpMessage = "User for which to delete folders")]
        [String]$User,

        [Parameter(Mandatory = $False, HelpMessage = "Well-known name of folder under which to clean")]
        [String]$WellKnownRoot = "deleteditems"
    )    

    # SOAP message for getting the folder id
    $body = @"
<FindFolder Traversal="Deep" xmlns="http://schemas.microsoft.com/exchange/services/2006/messages">
    <FolderShape>
        <t:BaseShape>Default</t:BaseShape>
        <t:AdditionalProperties>
            <t:ExtendedFieldURI PropertyTag="0x3007" PropertyType="SystemTime"/>
        </t:AdditionalProperties>
    </FolderShape>
    <ParentFolderIds>
        <t:DistinguishedFolderId Id="$WellKnownRoot"/>
    </ParentFolderIds>
</FindFolder>
"@

    Write-Host "Looking for folders under well-known folder: $WellKnownRoot & matching folder: $FolderNamePurge$FolderNamePrefixPurge & for user: $User"
    $getFolderIdMsg = Initialize-SOAPMessage -User $User -Body $body
    $response = Invoke-SOAPRequest -Token $Token -Message $getFolderIdMsg

    # Get the folders from the response
    $folders = $response | Select-Xml -XPath "//t:Folders/*" -Namespace @{t = "http://schemas.microsoft.com/exchange/services/2006/types"} | 
                    Select-Object -ExpandProperty Node

    $folderId = $null
    $changeKey = $null
    $totalCount = $null

    # Loop through folders
    foreach ($folder in $folders) {
        $folderId = $folder.FolderId.Id
        $changeKey = $folder.FolderId.Changekey
        $totalCount = $folder.TotalCount
        $folderName = $folder.DisplayName
        $folderCreateTime = $folder.ExtendedProperty
            | Where-Object { $_.ExtendedFieldURI.PropertyTag -eq "0x3007" }
            | Select-Object -ExpandProperty Value
            | Get-Date

        if ((![String]::IsNullOrEmpty($FolderNamePurge) -and $folderName -ne $FolderNamePurge) -or
            (![String]::IsNullOrEmpty($FolderPrefixPurge) -and $folderName -notlike "$FolderPrefixPurge*") -or
            (![String]::IsNullOrEmpty($FolderBeforePurge) -and $folderCreateTime -gt $FolderBeforePurge)) {
            continue
        }

        if (![String]::IsNullOrEmpty($FolderNamePurge)) {
            Write-Host "Found desired folder to purge: $FolderNamePurge"
        }

        Write-Verbose "Folder Id and ChangeKey for ""$folderName"": $folderId, $changeKey"

        # Empty and delete the folder if found
        if (![String]::IsNullOrEmpty($folderId) -and ![String]::IsNullOrEmpty($changeKey)) {
            if ($PSCmdlet.ShouldProcess("$folderName ($totalCount items) created $folderCreateTime", "Emptying folder")) {
                Write-Host "Emptying folder $folderName ($totalCount items)..."

                # DeleteType = HardDelete, MoveToDeletedItems, or SoftDelete
                $body = @"
        <m:EmptyFolder DeleteType="HardDelete" DeleteSubFolders="true">
            <m:FolderIds>
                <t:FolderId Id="$folderId" ChangeKey="$changeKey" />
            </m:FolderIds>
        </m:EmptyFolder>
"@
                $emptyFolderMsg = Initialize-SOAPMessage -User $User -Body $body
                $response = Invoke-SOAPRequest -Token $Token -Message $emptyFolderMsg
            }

            if ($PSCmdlet.ShouldProcess($folderName, "Deleting folder")) {
                Write-Host "Deleting folder $folderName..."

                # DeleteType = HardDelete, MoveToDeletedItems, or SoftDelete
                $body = @"
            <m:DeleteFolder  DeleteType="HardDelete" DeleteSubFolders="true">
                <m:FolderIds>
                    <t:FolderId Id="$folderId" ChangeKey="$changeKey" />
                </m:FolderIds>
            </m:DeleteFolder>
"@
                $deleteFolderMsg = Initialize-SOAPMessage -User $User -Body $body
                $response = Invoke-SOAPRequest -Token $Token -Message $deleteFolderMsg
            }

            Write-Host "Deleted folder $folderName ($totalCount items)"
        }
    }
}

$token = Get-AccessToken | ConvertTo-SecureString -AsPlainText -Force 

Remove-Folder -Token $token  -User $User -WellKnownRoot $WellKnownRoot