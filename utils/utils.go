package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func validateFilepath(path string, expectedExtension string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", path)
	}

	if filepath.Ext(path) != expectedExtension {
		return false, fmt.Errorf("file has incorrect extension, expected %s", expectedExtension)
	}

	return true, nil
}

func ValidateYglFilePath(path string) (bool, error) {
	return validateFilepath(path, ".ygl")
}
