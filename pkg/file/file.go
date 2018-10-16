package file

import (
	"fmt"
	"os"
)

func EnsureFileMode(path string) error {
	if err := os.Chmod(path, os.FileMode(0600)); err != nil {
		return fmt.Errorf("failed to set 0600 permissions on file '%s'", path)
	}

	return nil
}
