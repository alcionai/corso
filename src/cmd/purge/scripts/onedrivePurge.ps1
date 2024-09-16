[CmdletBinding(SupportsShouldProcess)]
Param (
    [Parameter(Mandatory = $False, HelpMessage = "User for which to delete folders in OneDrive")]
    [String]$User,

    [Parameter(Mandatory = $False, HelpMessage = "Site for which to delete folders in SharePoint")]
    [String]$Site,

    [Parameter(Mandatory = $False, HelpMessage = "Document library root. Can add multiple comma-separated values")]
    [String[]]$LibraryNameList = @(),

    [Parameter(Mandatory = $True, HelpMessage = "Purge folders before this date time (UTC)")]
    [datetime]$PurgeBeforeTimestamp,

    [Parameter(Mandatory = $True, HelpMessage = "Purge folders with this prefix")]
    [String[]]$FolderPrefixPurgeList,

    [Parameter(Mandatory = $False, HelpMessage = "Delete document libraries with this prefix")]
    [String[]]$LibraryPrefixDeleteList = @(),

    [Parameter(Mandatory = $False, HelpMessage = "Tenant domain")]
    [String]$TenantDomain = $ENV:TENANT_DOMAIN,

    [Parameter(Mandatory = $False, HelpMessage = "Azure ClientId")]
    [String]$ClientId = $ENV:AZURE_CLIENT_ID,

    [Parameter(Mandatory = $False, HelpMessage = "Azure AppCert")]
    [String]$AppCert = $ENV:AZURE_APP_CERT
)

Set-StrictMode -Version 2.0
# Attempt to set network timeout to 10min
[System.Net.ServicePointManager]::MaxServicePointIdleTime = 600000

function Get-TimestampFromFolderName {
    param (
        [Parameter(Mandatory = $True, HelpMessage = "Folder ")]
        [Microsoft.SharePoint.Client.Folder]$folder
    )

    $name = $folder.Name

    #fallback on folder create time
    [datetime]$timestamp = $folder.TimeCreated

    try {
        # Assumes that the timestamp is at the end and starts with yyyy-mm-ddT and is ISO8601
        if ($name -imatch "(\d{4}}-\d{2}-\d{2}T.*)") {
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

function Get-TimestampFromListName {
    param (
        [Parameter(Mandatory = $True, HelpMessage = "List ")]
        [Microsoft.SharePoint.Client.List]$list
    )

    $name = $list.Title

    #fallback on list create time
    [datetime]$timestamp = $list.LastItemUserModifiedDate

    try {
        # Assumes that the timestamp is at the end and starts with yyyy-mm-ddT and is ISO8601
        if ($name -imatch "(\d{4}}-\d{2}-\d{2}T.*)") {
            $timestamp = [System.Convert]::ToDatetime($Matches.0)
        }

        # Assumes that the timestamp is at the end and starts with dd-MMM-yyyy_HH-MM-SS
        if ($name -imatch "(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}-\d{2}-\d{2})") {
            $timestamp = [datetime]::ParseExact($Matches.0, "dd-MMM-yyyy_HH-mm-ss", [CultureInfo]::InvariantCulture, "AssumeUniversal")
        }
    }
    catch {}

    Write-Verbose "List: $name, create timestamp: $timestamp"

    return $timestamp
}

function Purge-Library {
    [CmdletBinding(SupportsShouldProcess)]
    Param (
        [Parameter(Mandatory = $True, HelpMessage = "Document library root")]
        [String]$LibraryName,

        [Parameter(Mandatory = $True, HelpMessage = "Purge folders before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp,

        [Parameter(Mandatory = $True, HelpMessage = "Purge folders with this prefix")]
        [String[]]$FolderPrefixPurgeList,

        [Parameter(Mandatory = $True, HelpMessage = "Site suffix")]
        [String[]]$SiteSuffix
    )

    Write-Host "`nPurging library: $LibraryName"

    $foldersToPurge = @()
    $folders = Get-PnPFolderItem -FolderSiteRelativeUrl $LibraryName -ItemType Folder

    Write-Host "`nFolders: $folders"
    foreach ($f in $folders) {
        $folderName = $f.Name
        $createTime = Get-TimestampFromFolderName -Folder $f

        if ($PurgeBeforeTimestamp -gt $createTime) {
            foreach ($p in $FolderPrefixPurgeList) {
                if ($folderName -like "$p*") {
                    $foldersToPurge += $f
                }
            }
        }
    }

    Write-Host "Found"$foldersToPurge.count"folders to purge"

    foreach ($f in $foldersToPurge) {
        $folderName = $f.Name
        $siteRelativeParentPath = ""

        if ($f.ServerRelativeUrl -imatch "$SiteSuffix/{0,1}(.+?)/{0,1}$folderName$") {
            $siteRelativeParentPath = $Matches.1
        }

        if ($PSCmdlet.ShouldProcess("Name: " + $f.Name + " Parent: " + $siteRelativeParentPath, "Remove folder")) {
            Write-Host "Deleting folder: "$f.Name" with parent: $siteRelativeParentPath"
            try {
                Remove-PnPFolder -Name $f.Name -Folder $siteRelativeParentPath -Force
            }
            catch [ System.Management.Automation.ItemNotFoundException ] {
                Write-Host "Folder: "$f.Name" with parent: $siteRelativeParentPath is already deleted. Skipping..."
            }
        }
    }
}

function Delete-LibraryByPrefix {
    [CmdletBinding(SupportsShouldProcess)]
    Param (
        [Parameter(Mandatory = $True, HelpMessage = "Document library root")]
        [String]$LibraryNamePrefix,

        [Parameter(Mandatory = $True, HelpMessage = "Purge folders before this date time (UTC)")]
        [datetime]$PurgeBeforeTimestamp,

        [Parameter(Mandatory = $True, HelpMessage = "Site suffix")]
        [String[]]$SiteSuffix
    )

    Write-Host "`nDeleting library: $LibraryNamePrefix"

    $listsToDelete = @()
    $lists = Get-PnPList

    foreach ($l in $lists) {
        $listName = $l.Title
        $createTime = Get-TimestampFromListName -List $l

        if ($PurgeBeforeTimestamp -gt $createTime) {
            foreach ($p in $FolderPrefixPurgeList) {
                if ($listName -like "$p*") {
                    $listsToDelete += $l
                }
            }
        }
    }

    Write-Host "Found"$listsToDelete.count"lists to delete"

    foreach ($l in $listsToDelete) {
        $listName = $l.Title

        if ($PSCmdlet.ShouldProcess("Name: " + $l.Title + "Remove folder")) {
            Write-Host "Deleting list: "$l.Title
            try {
                $listInfo = Get-PnPList -Identity $l.Id | Select-Object -Property Hidden

                # Check if the 'hidden' property is true
                if ($listInfo.Hidden) {
                    Write-Host "List: $($l.Title) is hidden. Skipping..."
                    continue
                }

                Remove-PnPList -Identity $l.Id  -Force
            }
            catch [ System.Management.Automation.ItemNotFoundException ] {
                Write-Host "List: "$f.Name" is already deleted. Skipping..."
            }
        }
    }
}

######## MAIN #########

# Setup SharePointPnP
if (-not (Get-Module -ListAvailable -Name PnP.PowerShell)) {
    $ProgressPreference = 'SilentlyContinue'
    Install-Module -Name PnP.PowerShell -Force
    $ProgressPreference = 'Continue'
}


if ([string]::IsNullOrEmpty($ClientId) -or [string]::IsNullOrEmpty($AppCert)) {
    Write-Host "ClientId and AppCert required as arguments or environment variables."
    Exit
}

# Connet to OneDrive or Sharepoint
$siteUrl = $null
if (![string]::IsNullOrEmpty($User)) {
    # Works for dev domains where format is <user name>@<domain>.onmicrosoft.com
    $domain = $User.Split('@')[1].Split('.')[0]
    $userNameEscaped = $User.Replace('.', '_').Replace('@', '_')

    $siteUrl = "https://$domain-my.sharepoint.com/personal/$userNameEscaped/"

    if ($LibraryNameList.count -eq 0) {
        $LibraryNameList = @("Documents")
        Write-Host "`nUsing default OneDrive library: $LibraryNameList"
    }
}
elseif (![string]::IsNullOrEmpty($Site)) {
    $siteUrl = $Site

    if ($LibraryNameList.count -eq 0) {
        $LibraryNameList = @("Shared Documents")
        Write-Host "`nUsing default SharePoint library: $LibraryNameList"
    }
}
else {
    Write-Host "User (for OneDrive) or Site (for Sharepoint) is required"
    Exit
}

#extract the suffix after the domain
$siteSuffix = ""
if ($siteUrl -imatch "^.*?(?<=sharepoint.com)(.*?$)") {
    $siteSuffix = $Matches.1
}
else {
    Write-Host "Site url appears to be malformed"
    Exit
}

Write-Host "`nAuthenticating and connecting to $SiteUrl"
Connect-PnPOnline -Url $siteUrl -ClientId $ClientId -CertificateBase64Encoded $AppCert -Tenant $TenantDomain
Write-Host "Connected to $siteUrl`n"

# ensure that there are no unexpanded entries in the list of parameters
$LibraryNameList = $LibraryNameList | ForEach-Object { @($_.Split(',').Trim()) }
$FolderPrefixPurgeList = $FolderPrefixPurgeList | ForEach-Object { @($_.Split(',').Trim()) }

foreach ($library in $LibraryNameList) {
    Purge-Library -LibraryName $library -PurgeBeforeTimestamp $PurgeBeforeTimestamp -FolderPrefixPurgeList $FolderPrefixPurgeList -SiteSuffix $siteSuffix
}

foreach ($libraryPfx in $LibraryPrefixDeleteList) {
    Delete-LibraryByPrefix -LibraryNamePrefix $libraryPfx -PurgeBeforeTimestamp $PurgeBeforeTimestamp -SiteSuffix $siteSuffix
}
