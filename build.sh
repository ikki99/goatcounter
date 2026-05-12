#!/usr/bin/env bash
set -e

# Ensure bin directory exists
mkdir -p bin

# Get version tag
tag=$(git describe --tags --always 2>/dev/null || echo "dev")

build_flags="-trimpath -mod=vendor -ldflags='-extldflags=-static -w -s -X zgo.at/goatcounter/v2.Version=$tag' -tags=osusergo,netgo,sqlite_omit_load_extension"

echo "Version: $tag"

# 1. Build for Linux (amd64)
echo "Building for Linux (amd64)..."
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o bin/goatcounter-linux-amd64 $build_flags ./cmd/goatcounter
echo "Successfully built bin/goatcounter-linux-amd64"

# 2. Build for Windows (amd64)
echo -e "\nBuilding for Windows (amd64)..."
export CGO_ENABLED=0
export GOOS=windows
export GOARCH=amd64
go build -o bin/goatcounter-windows-amd64.exe $build_flags ./cmd/goatcounter
echo "Successfully built bin/goatcounter-windows-amd64.exe"
