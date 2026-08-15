package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/example/filevault/internal/vault"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s get|delete <name>\n", os.Args[0])
		os.Exit(2)
	}

	v := vault.New()
	action, name := os.Args[1], os.Args[2]

	switch action {
	case "get":
		data, err := v.Get(name)
		if errors.Is(err, vault.ErrNotFound) {
			fmt.Fprintf(os.Stderr, "not found: %s\n", name)
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "get failed: %v\n", err)
			os.Exit(1)
		}
		os.Stdout.Write(data)
	case "delete":
		err := v.Delete(name)
		if errors.Is(err, vault.ErrNotFound) {
			fmt.Fprintf(os.Stderr, "not found: %s\n", name)
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "delete failed: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown action %q\n", action)
		os.Exit(2)
	}
}
