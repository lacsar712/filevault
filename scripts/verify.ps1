$ErrorActionPreference = "Stop"
$env:GOTOOLCHAIN = "local"

Push-Location $PSScriptRoot\..

Write-Host "==> go test ./..."
go test ./...

Pop-Location
