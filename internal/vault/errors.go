package vault

import "errors"

var (
	// ErrNotFound indicates the requested file does not exist in the vault.
	ErrNotFound = errors.New("file not found")
	// ErrExists indicates a file with the same name already exists.
	ErrExists = errors.New("file already exists")
)
