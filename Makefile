.PHONY: build test verify verify-broken

GOTOOLCHAIN ?= local
export GOTOOLCHAIN

build:
	go build ./...

test:
	go test ./...

verify: test

verify-broken:
	go test ./internal/vault/ -run 'TestGetMissingReturnsErrNotFound|TestDeleteMissingReturnsErrNotFound|TestDeleteRemovesFile' -count=1 -v
