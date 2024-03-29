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

function Get-TimestampFromName {
    param (
        [Parameter(Mandatory = $True, HelpMessage = "name")]
        [String]$name,

        [Parameter(Mandatory = $True, HelpMessage = "Default timestamp if not found in name")]
        [datetime]$defaultTimestamp
    )

    #fallback on folder create time 
    [datetime]$timestamp = $defaultTimestamp

    try {
        # Assumes that the timestamp is at the end and starts with yyyy-mm-ddT and is ISO8601
        if ($name -imatch "(\d{4}-\d{2}-\d{2}T[\S]*)") {
            $timestamp = [System.Convert]::ToDatetime($Matches.0)
        }

        # Assumes that the timestamp is at the end and starts with dd-MMM-yyyy_HH-MM-SS
        if ($name -imatch "(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}-\d{2}-\d{2})") {
            $timestamp = [datetime]::ParseExact($Matches.0, "dd-MMM-yyyy_HH-mm-ss", [CultureInfo]::InvariantCulture, "AssumeUniversal")
        }
    }
    catch {}

    Write-Verbose "Folder: $name, create timestamp: $timestamp"

    return $timestamp
}

function IsPrefixAndAgeMatch {
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Folder name to evaluate for match against a list of targets")]
        [string]$FolderName,

        [Parameter(Mandatory = $True, HelpMessage = "Folder creation times")]
        [datetime]$FolderCreateTime,

        [Parameter(Mandatory = $True, HelpMessage = "Folder name prefixes to evaluate for match")]
        [string[]]$FolderPrefixPurgeList,

        [Parameter(Mandatory = $TRUE, HelpMessage = "Purge folders before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp
    )
    
    $folderTimestamp = Get-TimestampFromName -name $FolderName -defaultTimestamp $FolderCreateTime
    
    if ($PurgeBeforeTimestamp -gt $folderTimestamp ) {
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
        [string[]]$FolderPrefixPurgeList = @(),

        [Parameter(Mandatory = $False, HelpMessage = "Perform shallow traversal only")]
        [bool]$PurgeTraversalShallow = $false
    )

    Write-Host "`nLooking for folders under well-known folder: $WellKnownRoot matching folders: $FolderNamePurgeList or prefixes: $FolderPrefixPurgeList for user: $User"
    
    $foldersToDelete = @()
    $traversal = "Deep"
    if ($PurgeTraversalShallow) {
        $traversal = "Shallow"
    }

    $offset = 0
    $moreToList = $true

    # get all folder pages
    while ($moreToList) {
        # SOAP message for getting the folders
        $body = @"
<FindFolder Traversal="$traversal" xmlns="http://schemas.microsoft.com/exchange/services/2006/messages">
    <FolderShape>
        <t:BaseShape>Default</t:BaseShape>
        <t:AdditionalProperties>
            <t:ExtendedFieldURI PropertyTag="0x3007" PropertyType="SystemTime"/>
        </t:AdditionalProperties>
    </FolderShape>
    <m:IndexedPageFolderView MaxEntriesReturned="1000" Offset="$offset" BasePoint="Beginning" />
    <ParentFolderIds>
        <t:DistinguishedFolderId Id="$WellKnownRoot"/>
    </ParentFolderIds>
</FindFolder>
"@

        try {
            Write-Host "`nRetrieving folders starting from offset: $offset"

            $getFolderIdMsg = Initialize-SOAPMessage -User $User -Body $body
            $response = Invoke-SOAPRequest -Token $Token -Message $getFolderIdMsg

            # Are there more folders to list
            $rootFolder = $response | Select-Xml -XPath "//m:RootFolder" -Namespace @{m = "http://schemas.microsoft.com/exchange/services/2006/messages" } | 
            Select-Object -ExpandProperty Node
            $moreToList = ![System.Convert]::ToBoolean($rootFolder.IncludesLastItemInRange)
        }
        catch {
            Write-Host "Error retrieving folders"

            Write-Host $response.OuterXml
            Exit
        }
    
        # Get the folders from the response
        $folders = $response | Select-Xml -XPath "//t:Folders/*" -Namespace @{t = "http://schemas.microsoft.com/exchange/services/2006/types" } | 
        Select-Object -ExpandProperty Node

        # Loop through folders
        foreach ($folder in $folders) {
            $folderName = $folder.DisplayName
            $folderCreateTime = $folder.ExtendedProperty
            | Where-Object { $_.ExtendedFieldURI.PropertyTag -eq "0x3007" }
            | Select-Object -ExpandProperty Value
            | Get-Date

            if ($FolderNamePurgeList.count -gt 0) {
                $IsNameMatchParams = @{
                    'FolderName'          = $folderName;
                    'FolderNamePurgeList' = $FolderNamePurgeList
                } 

                if ((IsNameMatch @IsNameMatchParams)) {
                    Write-Host "• Found name match: $folderName ($folderCreateTime)"
                    $foldersToDelete += $folder
                    continue
                }
            }

            if ($FolderPrefixPurgeList.count -gt 0) {
                $IsPrefixAndAgeMatchParams = @{
                    'FolderName'            = $folderName;
                    'FolderCreateTime'      = $folderCreateTime;
                    'FolderPrefixPurgeList' = $FolderPrefixPurgeList;
                    'PurgeBeforeTimestamp'  = $PurgeBeforeTimestamp;
                }

                if ((IsPrefixAndAgeMatch @IsPrefixAndAgeMatchParams)) {
                    Write-Host "• Found prefix match: $folderName ($folderCreateTime)" 
                    $foldersToDelete += $folder
                }
            }
        }

        if (!$moreToList -or $null -eq $folders) {
            Write-Host "Retrieved all folders."
        }
        else {
            $offset += $folders.count
        }
    }

    # powershel does not do well when returning empty arrays
    return , $foldersToDelete
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
        Write-Host "`nEmptying $foldersToEmptyCount folders..."
        foreach ($folder in $FolderNameList) {
            Write-Host "• $folder"
        }
        foreach ($folder in $WellKnownRootList) {
            Write-Host "• $folder"
        }

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
        foreach ($folder in $FolderNameList) {
            Write-Host "• $folder"
        }

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
        [datetime]$PurgeBeforeTimestamp,

        [Parameter(Mandatory = $False, HelpMessage = "Perform shallow traversal only")]
        [bool]$PurgeTraversalShallow = $false
    )  

    if (($FolderNamePurgeList.count -eq 0) -and 
        ($FolderPrefixPurgeList.count -eq 0 -or $PurgeBeforeTimestamp -eq $null )) {
        Write-Host "Either a list of specific folders or a list of prefixes and purge timestamp is required"
        Exit
    }

    if ($FolderNamePurgeList.count -gt 0) {
        Write-Host "Folders with names: $FolderNamePurgeList"
    }

    if ($FolderPrefixPurgeList.count -gt 0 -and $PurgeBeforeTimestamp -ne $null) {
        Write-Host "Folders older than $PurgeBeforeTimestamp with prefix:"
        foreach ($folder in $FolderPrefixPurgeList) {
            Write-Host "• $folder"
        }
    }

    $foldersToDeleteParams = @{
        'WellKnownRoot'         = $WellKnownRoot;
        'FolderNamePurgeList'   = $FolderNamePurgeList;
        'FolderPrefixPurgeList' = $FolderPrefixPurgeList;
        'PurgeBeforeTimestamp'  = $PurgeBeforeTimestamp;
        'PurgeTraversalShallow' = $PurgeTraversalShallow
    }

    $foldersToDelete = Get-FoldersToPurge @foldersToDeleteParams
    $foldersToDeleteCount = $foldersToDelete.count
    $foldersToDeleteIds = @()
    $folderNames = @()

    if ($foldersToDeleteCount -eq 0) {
        Write-Host "`nNo folders to purge matching the criteria"
        return
    }

    foreach ($folder in $foldersToDelete) {
        $foldersToDeleteIds += $folder.FolderId.Id
        $folderNames += $folder.DisplayName
    }
     
    Empty-Folder -FolderIdList $foldersToDeleteIds -FolderNameList $folderNames
    Delete-Folder -FolderIdList $foldersToDeleteIds -FolderNameList $folderNames
}

function Create-Contact {
    [CmdletBinding(SupportsShouldProcess)]

    $now = (Get-Date (Get-Date).ToUniversalTime() -Format "o")
    #used to create a recent seed contact that will be shielded from cleanup. CI tests rely on this    
    $body = @"
<CreateItem xmlns="http://schemas.microsoft.com/exchange/services/2006/messages" >
    <SavedItemFolderId>
        <t:DistinguishedFolderId Id="contacts"/>
    </SavedItemFolderId>
    <Items>
        <t:Contact>
            <t:GivenName>Sanitago</t:GivenName>
            <t:Surname>TestContact - $now</t:Surname>
            <t:CompanyName>Corso test enterprises</t:CompanyName>
            <t:EmailAddresses>
                <t:Entry Key="EmailAddress1">sanitago@example.com</t:Entry>
            </t:EmailAddresses>
            <t:PhoneNumbers>
                <t:Entry Key="BusinessPhone">4255550199</t:Entry>
            </t:PhoneNumbers>
            <t:Birthday>2000-01-01T11:59:00Z</t:Birthday>
            <t:JobTitle>Tester</t:JobTitle>
        </t:Contact>
    </Items>
</CreateItem>
"@

    if ($PSCmdlet.ShouldProcess("Creating seed contact...", "", "Create contact")) {
        Write-Host "`nCreating seed contact..."
        $createContactMsg = Initialize-SOAPMessage -User $User -Body $body
        $response = Invoke-SOAPRequest -Token $Token -Message $createContactMsg
    }
}

function Get-ItemsToPurge {
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Folder under which to look for items matching removal criteria")]
        [String]$WellKnownRoot,

        [Parameter(Mandatory = $False, HelpMessage = "Immediate subfolder within well known folder")]
        [String]$SubFolderName = $null,

        [Parameter(Mandatory = $True, HelpMessage = "Purge items before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp
    )

    $itemsToDelete = @()
    $foldersToSearchBody = "<t:DistinguishedFolderId Id='$WellKnownRoot'/>"

    if (![String]::IsNullOrEmpty($SubFolderName)) {
        $subFolders = Get-FoldersToPurge -WellKnownRoot $WellKnownRoot -FolderNamePurgeList $SubFolderName -PurgeBeforeTimestamp $PurgeBeforeTimestamp

        if ($subFolders.count -gt 0 ) {
            $foldersToSearchBody = ""
            foreach ($sub in $subFolders) {
                $subName = $sub.DisplayName
                $subId = $sub.FolderId.Id
                Write-Host "Found subfolder from which to purge items: $subName"
                $foldersToSearchBody = "<t:FolderId Id='$subId'/>`n"
            }
        }
        else {
            Write-Host "Requested subfolder $SubFolderName in folder $WellKnownRoot was not found"
            return
        }
    }

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
        $FoldersToSearchBody
    </ParentFolderIds>
</FindItem>
"@

    Write-Host "`nLooking for items under well-known folder: $WellKnownRoot($SubFolderName) older than $PurgeBeforeTimestamp for user: $User"
    $getItemsMsg = Initialize-SOAPMessage -User $User -Body $body
    $response = Invoke-SOAPRequest -Token $Token -Message $getItemsMsg

    # Get the contacts from the response
    $items = $response | Select-Xml -XPath "//t:Items/*" -Namespace @{t = "http://schemas.microsoft.com/exchange/services/2006/types" } | 
    Select-Object -ExpandProperty Node

    # Are there more folders to list
    $rootFolder = $response | Select-Xml -XPath "//m:RootFolder" -Namespace @{m = "http://schemas.microsoft.com/exchange/services/2006/messages" } | 
    Select-Object -ExpandProperty Node
    $moreToList = ![System.Convert]::ToBoolean($rootFolder.IncludesLastItemInRange)

    Write-Host "Total items under $WellKnownRoot/$SubFolderName"$rootFolder.TotalItemsInView

    foreach ($item in $items) {
        $itemId = $item.ItemId.Id
        $changeKey = $item.ItemId.Changekey
        $itemName = ""
        $itemCreateTime = $item.ExtendedProperty
        | Where-Object { $_.ExtendedFieldURI.PropertyTag -eq "0x3007" }
        | Select-Object -ExpandProperty Value
        | Get-Date

        # can be improved to pass the field to use as a name as a parameter but this is good for now
        switch -casesensitive ($WellKnownRoot) {
            "calendar" { $itemName = $item.Subject }
            "contacts" { $itemName = $item.DisplayName }
            Default { $itemName = $item.DisplayName }
        }

        if ([String]::IsNullOrEmpty($itemId) -or [String]::IsNullOrEmpty($changeKey)) {
            continue
        }

        $itemTimestamp = Get-TimestampFromName -name $itemName -defaultTimestamp $itemCreateTime

        if (![String]::IsNullOrEmpty($PurgeBeforeTimestamp) -and $itemTimestamp -gt $PurgeBeforeTimestamp) {
            continue
        }

        Write-Verbose "Item Id and ChangeKey for ""$itemName"": $itemId, $changeKey"
        $itemsToDelete += $item
    }

    if ($WhatIfPreference) {
        # not actually deleting items so only do a single iteration
        $moreToList = $false
    }

    return $itemsToDelete, $moreToList
}

function Purge-Items {
    [CmdletBinding(SupportsShouldProcess)]
    Param(
        [Parameter(Mandatory = $True, HelpMessage = "Purge items before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp, 

        [Parameter(Mandatory = $True, HelpMessage = "Items folder")]
        [string]$ItemsFolder,

        [Parameter(Mandatory = $False, HelpMessage = "Items sub-folder")]
        [string]$ItemsSubFolder = $null

    )

    $additionalAttributes = "SendMeetingCancellations='SendToNone'"

    Write-Host "`nCleaning up items from folder $ItemsFolder($ItemsSubFolder) older than $PurgeBeforeTimestamp" 
    Write-Host "-----------------------------------------------------------------------------" 

    if ($ItemsFolder -eq "contacts") {
        $ItemsSubFolder = $null
        $additionalAttributes = ""

        # Create one seed contact which will have recent create date and will not be sweapt
        # This is needed since tests rely on some contact data being present
        Create-Contact
    }

    $moreToList = $True
    # only get max of 1000 results so we may need to iterate over eligible contacts  
    while ($moreToList) {
        $itemsToDelete, $moreToList = Get-ItemsToPurge -WellKnownRoot $ItemsFolder -SubFolderName $ItemsSubFolder -PurgeBeforeTimestamp $PurgeBeforeTimestamp
        $itemsToDeleteCount = $itemsToDelete.count
        $itemsToDeleteBody = ""

        if ($itemsToDeleteCount -eq 0) {
            Write-Host "`nNo more items to delete matching criteria"
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
<m:DeleteItem DeleteType="HardDelete" $additionalAttributes>
    <m:ItemIds>
        $itemsToDeleteBody
    </m:ItemIds>
</m:DeleteItem>
"@

        if ($PSCmdlet.ShouldProcess("Deleting $itemsToDeleteCount items...", "$itemsToDeleteCount items", "Delete items")) {
            Write-Host "`nDeleting $itemsToDeleteCount items..."
        
            $emptyFolderMsg = Initialize-SOAPMessage -User $User -Body $body
            $response = Invoke-SOAPRequest -Token $Token -Message $emptyFolderMsg

            Write-Verbose "Delete response:`n"
            Write-Verbose $response.OuterXml
        
            Write-Host "`nDeleted $itemsToDeleteCount items..."
        }
    }
}

### MAIN ####

Write-Host 'Authenticating with Exchange Web Services ...'
$global:Token = Get-AccessToken | ConvertTo-SecureString -AsPlainText -Force 

# ensure that there are no unexpanded entries in the list of parameters
$FolderNamePurgeList = $FolderNamePurgeList | ForEach-Object { @($_.Split(',').Trim()) }
$FolderPrefixPurgeList = $FolderPrefixPurgeList | ForEach-Object { @($_.Split(',').Trim()) }

Write-Host "`nPurging CI-produced folders under 'msgfolderroot' ..."
Write-Host "--------------------------------------------------------"

$purgeFolderParams = @{
    'WellKnownRoot'         = "msgfolderroot";
    'FolderNamePurgeList'   = $FolderNamePurgeList;
    'FolderPrefixPurgeList' = $FolderPrefixPurgeList;
    'PurgeBeforeTimestamp'  = $PurgeBeforeTimestamp
}

#purge older prefix folders from msgfolderroot
Purge-Folders @purgeFolderParams

#purge older contacts 
Purge-Items -ItemsFolder "contacts" -PurgeBeforeTimestamp $PurgeBeforeTimestamp

#purge older contact birthday events
Purge-Items -ItemsFolder "calendar" -ItemsSubFolder "Birthdays" -PurgeBeforeTimestamp $PurgeBeforeTimestamp

# Empty Deleted Items and then purge all recoverable items. Deletes the following
# -/Recoverable Items/Audits
# -/Recoverable Items/Deletion
# -/Recoverable Items/Purges
# -/Recoverable Items/Versions
# -/Recoverable Items/Calendar Logging
# -/Recoverable Items/SubstrateHolds
Write-Host "`nProcess well-known folders that are always purged" 
Write-Host "---------------------------------------------------" 

# We explicitly also clean direct folders under Deleted Items since there is some evidence
# that suggests that emptying alone may not be reliable
Write-Host "`nExplicit delete of all folders under 'DeletedItems' ..."
Write-Host "----------------------------------------------------------"

$purgeFolderParams = @{
    'WellKnownRoot'         = "deleteditems";
    'FolderNamePurgeList'   = $FolderNamePurgeList;
    'FolderPrefixPurgeList' = @('*');
    'PurgeBeforeTimestamp'  = (Get-Date);
    'PurgeTraversalShallow' = $true
}

Purge-Folders @purgeFolderParams

Empty-Folder -WellKnownRootList "deleteditems", "recoverableitemsroot"
