package test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func loadSamples(dirPath string) [][]byte {
	files := make([][]byte, 0)
	var entries []os.DirEntry
	var err error
	if entries, err = os.ReadDir(dirPath); err != nil {
		log.Fatalln(err.Error())
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		var file []byte
		samplePath := fmt.Sprintf("./%s/%s", dirPath, entry.Name())
		file, err = os.ReadFile(samplePath)
		files = append(files, file)
	}

	return files
}

func TestCompiler(t *testing.T) {
	validFiles := loadSamples("./samples/valid")
	for _, file := range validFiles {
		
	}
}
