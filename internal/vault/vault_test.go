package vault_test

import (
	"errors"
	"testing"

	"github.com/example/filevault/internal/vault"
)

func TestPutGetRoundTrip(t *testing.T) {
	v := vault.New()
	if err := v.Put("readme.txt", []byte("hello")); err != nil {
		t.Fatalf("put: %v", err)
	}

	got, err := v.Get("readme.txt")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if string(got) != "hello" {
		t.Fatalf("got %q, want %q", got, "hello")
	}
}

func TestPutDuplicateReturnsErrExists(t *testing.T) {
	v := vault.New()
	if err := v.Put("dup.txt", []byte("a")); err != nil {
		t.Fatalf("first put: %v", err)
	}

	err := v.Put("dup.txt", []byte("b"))
	if !errors.Is(err, vault.ErrExists) {
		t.Fatalf("expected ErrExists, got %v", err)
	}
}

func TestGetMissingReturnsErrNotFound(t *testing.T) {
	v := vault.New()

	_, err := v.Get("missing.txt")
	if !errors.Is(err, vault.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDeleteMissingReturnsErrNotFound(t *testing.T) {
	v := vault.New()

	err := v.Delete("ghost.txt")
	if !errors.Is(err, vault.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDeleteRemovesFile(t *testing.T) {
	v := vault.New()
	if err := v.Put("temp.txt", []byte("x")); err != nil {
		t.Fatalf("put: %v", err)
	}
	if err := v.Delete("temp.txt"); err != nil {
		t.Fatalf("delete: %v", err)
	}

	_, err := v.Get("temp.txt")
	if !errors.Is(err, vault.ErrNotFound) {
		t.Fatalf("expected ErrNotFound after delete, got %v", err)
	}
}

func TestListReturnsStoredNames(t *testing.T) {
	v := vault.New()
	_ = v.Put("a.txt", []byte("1"))
	_ = v.Put("b.txt", []byte("2"))

	names := v.List()
	if len(names) != 2 {
		t.Fatalf("expected 2 names, got %d: %v", len(names), names)
	}
}
