[CmdletBinding()]

Param (
    [Parameter(Mandatory = $True, HelpMessage = "Powershell script URL")]
    [string]$ScriptURL
)

Invoke-WebRequest -Uri $ScriptURL -OutFile ./script.ps1

./script.ps1