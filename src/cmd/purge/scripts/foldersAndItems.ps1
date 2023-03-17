[CmdletBinding(SupportsShouldProcess)]

Param (
    [Parameter(Mandatory = $True, HelpMessage = "User for which to delete folders")]
    [String]$User,

    [Parameter(Mandatory = $True, HelpMessage = "Purge folders or contacts before this date time (UTC)")]
    [datetime]$PurgeBeforeTimestamp,

    [Parameter(Mandatory = $True, HelpMessage = "Name of specific folder to purge under root")]
    [String[]]$FolderNamePurgeList,

    [Parameter(Mandatory = $True, HelpMessage = "Purge folders with this prefix")]
    [String[]]$FolderPrefixPurgeList,

    [Parameter(Mandatory = $False, HelpMessage = "Azure TenantId")]
    [String]$TenantId = $ENV:AZURE_TENANT_ID,

    [Parameter(Mandatory = $False, HelpMessage = "Azure ClientId")]
    [String]$ClientId = $ENV:AZURE_CLIENT_ID,

    [Parameter(Mandatory = $False, HelpMessage = "Azure ClientSecret")]
    [String]$ClientSecret = $ENV:AZURE_CLIENT_SECRET
)

Set-StrictMode -Version 2.0
# Attempt to set network timeout to 10min
[System.Net.ServicePointManager]::MaxServicePointIdleTime = 600000

function Get-AccessToken {
    [CmdletBinding()]
    Param()

    if ([String]::IsNullOrEmpty($TenantId) -or [String]::IsNullOrEmpty($ClientId) -or [String]::IsNullOrEmpty($ClientSecret)) {
        Write-Host "`nNeed to specify TenantId, ClientId, and ClientSecret as parameters or ENVs"
    }

    $body = @{
        client_id     = $ClientId
        client_secret = $ClientSecret
        scope         = "https://outlook.office365.com/.default"
        grant_type    = "client_credentials"
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

function IsNameMatch {
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Folder name to evaluate for match against a list of targets")]
        [string]$FolderName,

        [Parameter(Mandatory = $True, HelpMessage = "Folder names to evaluate for match")]
        [string[]]$FolderNamePurgeList = @()
    )

    return ($FolderName -in $FolderNamePurgeList)
}

function IsPrefixAndAgeMatch {
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Folder name to evaluate for match against a list of targets")]
        [string]$FolderName,

        [Parameter(Mandatory = $True, HelpMessage = "Folder creation times")]
        [string]$FolderCreateTime,

        [Parameter(Mandatory = $True, HelpMessage = "Folder name prefixes to evaluate for match")]
        [string[]]$FolderPrefixPurgeList,

        [Parameter(Mandatory = $TRUE, HelpMessage = "Purge folders before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp
    )

    if ($PurgeBeforeTimestamp -gt $folderCreateTime ) {
        foreach ($prefix in $FolderPrefixPurgeList) {
            if ($FolderName -like "$prefix*") {
                return $true
            }
        }
    }

    return $false
}

function Get-FoldersToPurge {
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Folder under which to look for items matching removal criteria")]
        [String]$WellKnownRoot,

        [Parameter(Mandatory = $False, HelpMessage = "Purge folders before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp,

        [Parameter(Mandatory = $False, HelpMessage = "Purge folders with these names")]
        [string[]]$FolderNamePurgeList = @(),

        [Parameter(Mandatory = $False, HelpMessage = "Purge folders with these prefixes")]
        [string[]]$FolderPrefixPurgeList = @()
    )

    $foldersToDelete = @()
    
    # SOAP message for getting the folders
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

    Write-Host "`nLooking for folders under well-known folder: $WellKnownRoot matching folders: $FolderNamePurgeList or prefixes: $FolderPrefixPurgeList for user: $User"
    $getFolderIdMsg = Initialize-SOAPMessage -User $User -Body $body
    $response = Invoke-SOAPRequest -Token $Token -Message $getFolderIdMsg

    # Get the folders from the response
    $folders = $response | Select-Xml -XPath "//t:Folders/*" -Namespace @{t = "http://schemas.microsoft.com/exchange/services/2006/types" } | 
    Select-Object -ExpandProperty Node

    # Are there more folders to list
    $rootFolder = $response | Select-Xml -XPath "//m:RootFolder" -Namespace @{m = "http://schemas.microsoft.com/exchange/services/2006/messages" } | 
    Select-Object -ExpandProperty Node
    $moreToList = ![System.Convert]::ToBoolean($rootFolder.IncludesLastItemInRange)

    # Loop through folders
    foreach ($folder in $folders) {
        $folderName = $folder.DisplayName
        $folderCreateTime = $folder.ExtendedProperty
        | Where-Object { $_.ExtendedFieldURI.PropertyTag -eq "0x3007" }
        | Select-Object -ExpandProperty Value
        | Get-Date

        $IsNameMatchParams = @{
            'FolderName'          = $folderName;
            'FolderNamePurgeList' = $FolderNamePurgeList
        } 

        $IsPrefixAndAgeMatchParams = @{
            'FolderName'            = $folderName;
            'FolderCreateTime'      = $folderCreateTime;
            'FolderPrefixPurgeList' = $FolderPrefixPurgeList;
            'PurgeBeforeTimestamp'  = $PurgeBeforeTimestamp;
        }

        if ((IsNameMatch @IsNameMatchParams) -or (IsPrefixAndAgeMatch @IsPrefixAndAgeMatchParams)) {
            Write-Host "`nFound desired folder to purge: $folderName ($folderCreateTime)"
            $foldersToDelete += $folder
        }
    }

    # powershel does not do well when returning empty arrays
    return $foldersToDelete, $moreToList
}

function Empty-Folder {
    [CmdletBinding(SupportsShouldProcess)]
    Param(
        [Parameter(Mandatory = $False, HelpMessage = "List of well-known folders to empty ")]
        [String[]]$WellKnownRootList = @(),

        [Parameter(Mandatory = $False, HelpMessage = "List of folderIds to empty ")]
        [string[]]$FolderIdList = @(),

        [Parameter(Mandatory = $False, HelpMessage = "List of folder names to empty ")]
        [string[]]$FolderNameList = @()
    )

    $folderIdsBody = ""
    $foldersToEmptyCount = $FolderIdList.count + $WellKnownRootList.count

    foreach ($wnr in $WellKnownRootList) {
        $folderIdsBody += "<t:DistinguishedFolderId Id='$wnr'/>"
    }

    foreach ($fid in $FolderIdList) {
        $folderIdsBody += "<t:FolderId Id='$fid'/>"
    }

    if ($PSCmdlet.ShouldProcess("Emptying $foldersToEmptyCount folders ($WellKnownRootList $FolderNameList)", "$foldersToEmptyCount folders ($WellKnownRootList $FolderNameList)", "Empty folders")) {
        Write-Host "`nEmptying $foldersToEmptyCount folders ($WellKnownRootList $FolderNameList)"

        # DeleteType = HardDelete, MoveToDeletedItems, or SoftDelete
        $body = @"
<m:EmptyFolder DeleteType="HardDelete" DeleteSubFolders="true">
    <m:FolderIds>
    $folderIdsBody
    </m:FolderIds>
</m:EmptyFolder>
"@

        $emptyFolderMsg = Initialize-SOAPMessage -User $User -Body $body
        $response = Invoke-SOAPRequest -Token $Token -Message $emptyFolderMsg
    }
}

function Delete-Folder {
    [CmdletBinding(SupportsShouldProcess)]
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "List of folderIds to remove ")]
        [String[]]$FolderIdList,

        [Parameter(Mandatory = $False, HelpMessage = "List of folder names to remove ")]
        [String[]]$FolderNameList = @()
    )

    $folderIdsBody = ""
    $foldersToRemoveCount = $FolderIdList.count

    foreach ($fid in $FolderIdList) {
        $folderIdsBody += "<t:FolderId Id='$fid'/>"
    }

    if ($PSCmdlet.ShouldProcess("Removing $foldersToRemoveCount folders ($FolderNameList)", "$foldersToRemoveCount folders ($FolderNameList)", "Delete folders")) {
        Write-Host "`nRemoving $foldersToRemoveCount folders ($FolderNameList)"

        # DeleteType = HardDelete, MoveToDeletedItems, or SoftDelete
        $body = @"
<m:DeleteFolder DeleteType="HardDelete" DeleteSubFolders="true">
    <m:FolderIds>
    $folderIdsBody
    </m:FolderIds>
</m:DeleteFolder>
"@

        $emptyFolderMsg = Initialize-SOAPMessage -User $User -Body $body
        $response = Invoke-SOAPRequest -Token $Token -Message $emptyFolderMsg
    }
}

function Purge-Folders {
    [CmdletBinding(SupportsShouldProcess)]
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Folder under which to look for items matching removal criteria")]
        [String]$WellKnownRoot,

        [Parameter(Mandatory = $False, HelpMessage = "Purge folders with these names")]
        [string[]]$FolderNamePurgeList = @(),

        [Parameter(Mandatory = $False, HelpMessage = "Purge folders with these prefixes")]
        [string[]]$FolderPrefixPurgeList = @(),

        [Parameter(Mandatory = $False, HelpMessage = "Purge folders before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp
    )  

    if (($FolderNamePurgeList.count -eq 0) -and 
        ($FolderPrefixPurgeList.count -eq 0 -or $PurgeBeforeTimestamp -eq $null )) {
        Write-Host "Either a list of specific folders or a list of prefixes and purge timestamp is required"
        Exit
    }

    Write-Host "`nPurging CI-produced folders..."
    Write-Host "--------------------------------"

    if ($FolderNamePurgeList.count -gt 0) {
        Write-Host "Folders with names: $FolderNamePurgeList"
    }

    if ($FolderPrefixPurgeList.count -gt 0 -and $PurgeBeforeTimestamp -ne $null) {
        Write-Host "Folders older than $PurgeBeforeTimestamp with prefix: $FolderPrefixPurgeList"
    }

    $foldersToDeleteParams = @{
        'WellKnownRoot'         = $WellKnownRoot;
        'FolderNamePurgeList'   = $FolderNamePurgeList;
        'FolderPrefixPurgeList' = $FolderPrefixPurgeList;
        'PurgeBeforeTimestamp'  = $PurgeBeforeTimestamp
    }

    $moreToList = $True
    # only get max of 1000 results so we may need to iterate over eligible folders  
    while ($moreToList) {
        $foldersToDelete, $moreToList = Get-FoldersToPurge @foldersToDeleteParams
        $foldersToDeleteCount = $foldersToDelete.count
        $foldersToDeleteIds = @()
        $folderNames = @()

        if ($foldersToDeleteCount -eq 0) {
            Write-Host "`nNo folders to purge matching the criteria"
            break
        }

        foreach ($folder in $foldersToDelete) {
            $foldersToDeleteIds += $folder.FolderId.Id
            $folderNames += $folder.DisplayName
        }
        
        Empty-Folder -FolderIdList $foldersToDeleteIds -FolderNameList $folderNames
        Delete-Folder -FolderIdList $foldersToDeleteIds -FolderNameList $folderNames
    }
}

function Get-ItemsToPurge {
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Folder under which to look for items matching removal criteria")]
        [String]$WellKnownRoot,

        [Parameter(Mandatory = $True, HelpMessage = "Purge items before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp
    )

    $itemsToDelete = @()

    # SOAP message for getting the folder id
    $body = @"
<FindItem Traversal="Shallow" xmlns="http://schemas.microsoft.com/exchange/services/2006/messages">
    <ItemShape>
        <t:BaseShape>Default</t:BaseShape>
        <t:AdditionalProperties>
            <t:ExtendedFieldURI PropertyTag="0x3007" PropertyType="SystemTime"/>
        </t:AdditionalProperties>
    </ItemShape>
    <ParentFolderIds>
        <t:DistinguishedFolderId Id="$WellKnownRoot"/>
    </ParentFolderIds>
</FindItem>
"@

    Write-Host "`nLooking for items under well-known folder: $WellKnownRoot older than $PurgeBeforeTimestamp for user: $User"
    $getItemsMsg = Initialize-SOAPMessage -User $User -Body $body
    $response = Invoke-SOAPRequest -Token $Token -Message $getItemsMsg

    # Get the contacts from the response
    $items = $response | Select-Xml -XPath "//t:Items/*" -Namespace @{t = "http://schemas.microsoft.com/exchange/services/2006/types" } | 
    Select-Object -ExpandProperty Node

    # Are there more folders to list
    $rootFolder = $response | Select-Xml -XPath "//m:RootFolder" -Namespace @{m = "http://schemas.microsoft.com/exchange/services/2006/messages" } | 
    Select-Object -ExpandProperty Node
    $moreToList = ![System.Convert]::ToBoolean($rootFolder.IncludesLastItemInRange)

    foreach ($item in $items) {
        $itemId = $item.ItemId.Id
        $changeKey = $item.ItemId.Changekey
        $itemName = $item.DisplayName
        $itemCreateTime = $item.ExtendedProperty
        | Where-Object { $_.ExtendedFieldURI.PropertyTag -eq "0x3007" }
        | Select-Object -ExpandProperty Value
        | Get-Date

        if ([String]::IsNullOrEmpty($itemId) -or [String]::IsNullOrEmpty($changeKey)) {
            continue
        }

        if (![String]::IsNullOrEmpty($PurgeBeforeTimestamp) -and $itemCreateTime -gt $PurgeBeforeTimestamp) {
            continue
        }

        Write-Verbose "Item Id and ChangeKey for ""$itemName"": $itemId, $changeKey"
        $itemsToDelete += $item
    }

    return $itemsToDelete, $moreToList
}

function Purge-Contacts {
    [CmdletBinding(SupportsShouldProcess)]
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Purge items before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp
    )

    Write-Host "`nCleaning up contacts older than $PurgeBeforeTimestamp" 
    Write-Host "-------------------------------------------------------" 

    $moreToList = $True
    # only get max of 1000 results so we may need to iterate over eligible contacts  
    while ($moreToList) {
        $itemsToDelete, $moreToList = Get-ItemsToPurge -WellKnownRoot "contacts" -PurgeBeforeTimestamp $PurgeBeforeTimestamp
        $itemsToDeleteCount = $itemsToDelete.count
        $itemsToDeleteBody = ""

        if ($itemsToDeleteCount -eq 0) {
            Write-Host "`nNo more contacts to delete matching criteria"
            break
        }

        Write-Host "`nQueueing $itemsToDeleteCount items to delete"
        foreach ($item in $itemsToDelete) {
            $itemId = $item.ItemId.Id
            $changeKey = $item.ItemId.Changekey
            $itemsToDeleteBody += "<t:ItemId Id='$itemId' ChangeKey='$changeKey' />`n"
        }

        # Do the actual deletion in a batch request 
        # DeleteType = HardDelete, MoveToDeletedItems, or SoftDelete
        $body = @"
<m:DeleteItem DeleteType="HardDelete">
    <m:ItemIds>
        $itemsToDeleteBody
    </m:ItemIds>
</m:DeleteItem>
"@
        
        if ($PSCmdlet.ShouldProcess("Deleting $itemsToDeleteCount items...", "$itemsToDeleteCount items", "Delete items")) {
            Write-Host "`nDeleting $itemsToDeleteCount items..."
        
            $emptyFolderMsg = Initialize-SOAPMessage -User $User -Body $body
            $response = Invoke-SOAPRequest -Token $Token -Message $emptyFolderMsg
        
            Write-Host "`nDeleted $itemsToDeleteCount items..."
        
        }
    }
}

Write-Host 'Authenticating with Exchange Web Services ...'
$global:Token = Get-AccessToken | ConvertTo-SecureString -AsPlainText -Force 

$purgeFolderParams = @{
    'WellKnownRoot'         = "root";
    'FolderNamePurgeList'   = $FolderNamePurgeList;
    'FolderPrefixPurgeList' = $FolderPrefixPurgeList;
    'PurgeBeforeTimestamp'  = $PurgeBeforeTimestamp
}

#purge older prefix folders
Purge-Folders @purgeFolderParams

#purge older contacts 
Purge-Contacts -PurgeBeforeTimestamp $PurgeBeforeTimestamp

# Empty Deleted Items and then purge all recoverable items. Deletes the following
# -/Recoverable Items/Audits
# -/Recoverable Items/Deletion
# -/Recoverable Items/Purges
# -/Recoverable Items/Versions
# -/Recoverable Items/Calendar Logging
# -/Recoverable Items/SubstrateHolds
Write-Host "`nProcess well-known folders that are always purged" 
Write-Host "---------------------------------------------------" 
Empty-Folder -WellKnownRoot "deleteditems", "recoverableitemsroot"