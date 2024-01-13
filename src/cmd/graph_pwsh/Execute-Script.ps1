[CmdletBinding()]

Param (
    [Parameter(Mandatory = $True, HelpMessage = "Powershell script URL")]
    [string]$ScriptURL
)

Invoke-WebRequest -Uri $ScriptURL -OutFile ./script.ps1

Write-Host "Executing the following script"
Write-Host "=============================="
cat ./script.ps1
Write-Host "=============================="

./script.ps1