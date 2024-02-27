package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath := "test.bay" // Change to the path of the file you want to validate

	valid, err := validateFilepath(filepath, ".ygl") // Specify the expected extension
	if err != nil {
		fmt.Println(err)
		return
	}
	if valid {
		fmt.Println("File is valid")
	} else {
		fmt.Println("File is not valid")
	}
}

func validateFilepath(path string, expectedExtension string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, fmt.Errorf("file %s does not exist", path)
	}

	if filepath.Ext(path) != expectedExtension {
		return false, fmt.Errorf("file has incorrect extension, expected %s", expectedExtension)
	}

	return true, nil
}
}
