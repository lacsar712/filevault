# filevault

In-memory file vault library (Go) with a small CLI.

## Requirements

- Go 1.21+ (container image uses golang:1.22)
- `GOTOOLCHAIN=local` recommended on host when using a pinned toolchain

## Build

```bash
go build ./...
```

## Run

```bash
go run ./cmd/filevault get missing.txt
go run ./cmd/filevault delete ghost.txt
```

## Test

```bash
go test ./...
```

## Docker (benzhi)

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh filevault-001 linux/amd64
./build_benzhi_docker.sh filevault-001 linux/arm64
docker run -it filevault-001:latest
# inside container:
go build ./...
go test ./...
```
