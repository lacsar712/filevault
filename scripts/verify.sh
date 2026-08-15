#!/usr/bin/env bash
set -euo pipefail
export GOTOOLCHAIN=local
cd "$(dirname "$0")/.."
echo "==> go test ./..."
go test ./...
