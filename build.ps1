# GoatCounter Build Script (Windows PowerShell)

$ErrorActionPreference = "Stop"

# Ensure bin directory exists
if (-not (Test-Path "bin")) {
    New-Item -ItemType Directory -Path "bin" | Out-Null
}

# Get version tag
$tag = git describe --tags --always 2>$null
if (-not $tag) { $tag = "dev" }

$build_flags = @(
    "-trimpath",
    "-mod=vendor",
    "-ldflags=-extldflags=-static -w -s -X zgo.at/goatcounter/v2.Version=$tag",
    "-tags=osusergo,netgo,sqlite_omit_load_extension"
)

Write-Host "Version: $tag" -ForegroundColor Cyan

# 1. Build for Windows (amd64)
Write-Host "Building for Windows (amd64)..."
$env:CGO_ENABLED = "0"
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o bin/goatcounter-windows-amd64.exe @build_flags ./cmd/goatcounter
if ($?) { Write-Host "Successfully built bin/goatcounter-windows-amd64.exe" -ForegroundColor Green }

# 2. Build for Linux (amd64)
Write-Host "`nBuilding for Linux (amd64)..."
$env:CGO_ENABLED = "0"
$env:GOOS = "linux"
$env:GOARCH = "amd64"

try {
    go build -o bin/goatcounter-linux-amd64 @build_flags ./cmd/goatcounter
    if ($?) { Write-Host "Successfully built bin/goatcounter-linux-amd64" -ForegroundColor Green }
} catch {
    Write-Host "Failed to build for Linux. This is likely due to missing cross-compilation toolchain for CGO." -ForegroundColor Red
}
