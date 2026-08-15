# filevault

In-memory file vault library with a small CLI.

## Requirements

- Go 1.21+
- `GOTOOLCHAIN=local` recommended when using a pinned toolchain

## Setup

```bash
cd filevault
go mod tidy
go build ./...
```

## Verify

```bash
go test ./...
```

Or use the Makefile / scripts:

```bash
make test
# Windows PowerShell
./scripts/verify.ps1
```

## CLI

```bash
go run ./cmd/filevault get missing.txt   # exits 1 with "not found"
go run ./cmd/filevault delete ghost.txt  # exits 1 with "not found"
```

The CLI relies on `errors.Is` to detect `vault.ErrNotFound`.
