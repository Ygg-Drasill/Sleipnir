package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath := "test.bay" // Change this to the path of the file you want to validate(eg. "test.ygl / bay.test")
	valid, err := validateFilepath(filepath)
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

// First we check if the path does exist, if not we return an error.
func validateFilepath(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, fmt.Errorf("file does not exist: %s", path)
	}

	//Then we check if the file is a ygl file, if not we return an error.
	if filepath.Ext(path) != ".ygl" {
		return false, fmt.Errorf("file is not a go file")
	}
	//If the file exists and is a ygl file, we return true and nil.
	return true, nil
}
