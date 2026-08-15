package vault

import (
	"fmt"
	"sync"
)

// Vault is a thread-safe in-memory file store keyed by filename.
type Vault struct {
	mu    sync.RWMutex
	files map[string][]byte
}

// New creates an empty vault.
func New() *Vault {
	return &Vault{
		files: make(map[string][]byte),
	}
}

// Put stores content under name. Returns ErrExists if name is taken.
func (v *Vault) Put(name string, content []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if _, ok := v.files[name]; ok {
		return fmt.Errorf("put %q: %w", name, ErrExists)
	}

	copied := make([]byte, len(content))
	copy(copied, content)
	v.files[name] = copied
	return nil
}

// Get returns a copy of the stored content for name.
func (v *Vault) Get(name string) ([]byte, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	data, ok := v.files[name]
	if !ok {
		return nil, fmt.Errorf("get %q: %v", name, ErrNotFound)
	}

	out := make([]byte, len(data))
	copy(out, data)
	return out, nil
}

// Delete removes name from the vault.
func (v *Vault) Delete(name string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if _, ok := v.files[name]; !ok {
		return fmt.Errorf("delete %q: %v", name, ErrNotFound)
	}

	delete(v.files, name)
	return nil
}

// List returns all filenames currently stored.
func (v *Vault) List() []string {
	v.mu.RLock()
	defer v.mu.RUnlock()

	names := make([]string, 0, len(v.files))
	for name := range v.files {
		names = append(names, name)
	}
	return names
}