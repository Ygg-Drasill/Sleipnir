package test

import (
	"bytes"
	"github.com/Ygg-Drasill/Sleipnir/pkg/compiler"
	"log"
	"os"
	"testing"
	"fmt"
)

func samplesInDir(dirPath string) []os.FileInfo {
	sampleEntries := make([]os.FileInfo, 0)
	var entries []os.DirEntry
	var err error
	if entries, err = os.ReadDir(dirPath); err != nil {
		log.Fatalln(err.Error())
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			log.Fatalln(err.Error())
		}
		sampleEntries = append(sampleEntries, info)
	}

	return sampleEntries
}

const validPath = "./samples/valid/"
const invalidPath = "./samples/invalid/"

func TestCompileValidSourceCode(t *testing.T) {
	sampleFiles := samplesInDir(validPath)
	for _, fileInfo := range sampleFiles {
		t.Run(fileInfo.Name(), func(t *testing.T) {
			t.Parallel()
			var err error
			var file []byte
			file, err = os.ReadFile(validPath + fileInfo.Name())
			if err != nil {
				t.Fatalf("failed to read file: %s\n%s", fileInfo.Name(), err.Error())
			}
			sample := string(file)
			c := compiler.NewFromString(sample)
			err = c.Compile()
			if err != nil {
				t.Fatalf("failed to compile valid source code sample: %s\n%s", fileInfo.Name(), err.Error())
			}
			output := new(bytes.Buffer)
			var n int
			n, err = c.WriteOutputToBuffer(output)
			if err != nil {
				t.Fatalf("failed to write sample output: %s\n%s", fileInfo.Name(), err.Error())
			}
			if n == 0 {
				t.Fatalf("no bytes was produced after compilation of valid sample: %s", fileInfo.Name())
			}
		})
	}
}

func TestCompileInValidSourceCode(t *testing.T) {
	sampleDirPath := invalidPath
	sampleFiles := samplesInDir(sampleDirPath)
	for _, fileInfo := range sampleFiles {
		t.Run(fileInfo.Name(), func(t *testing.T) {
			t.Parallel()
			var err error
			var file []byte
			file, err = os.ReadFile(sampleDirPath + fileInfo.Name())
			if err != nil {
				t.Fatalf("failed to read file: %s\n%s", fileInfo.Name(), err.Error())
			}
			sample := string(file)
			c := compiler.NewFromString(sample)
			err = c.Compile()
			if err == nil {
				t.Fatalf("invalid source code was accepted by compiler: %s", fileInfo.Name())
			}
		})
	}
}
