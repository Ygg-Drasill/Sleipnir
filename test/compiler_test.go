package test

import (
	"bytes"
	"github.com/Ygg-Drasill/Sleipnir/pkg/compiler"
	"github.com/bytecodealliance/wasmtime-go"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
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

const resultSeparator = "expect:"

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
			sampleData := strings.Split(string(file), resultSeparator)
			c := compiler.NewFromString(sampleData[0])
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

			//validate result wasm
			wasm := make([]byte, 0)
			module := new(wasmtime.Module)
			engine := wasmtime.NewEngine()
			instance := new(wasmtime.Instance)
			store := wasmtime.NewStore(engine)
			var result interface{}
			wasm, err = wasmtime.Wat2Wasm(output.String())
			if err != nil {
				t.Fatalf("failed to parse resulting webassembly: %s", err.Error())
			}
			module, err = wasmtime.NewModule(engine, wasm)
			if err != nil {
				t.Fatalf("failed to convert wasm to module: %s", err.Error())
			}

			var sampleDataValue int64
			var want int32
			if len(sampleData) > 1 {
				sampleDataValue, err = strconv.ParseInt(sampleData[1], 10, 32)
				if err != nil {
					t.Fatalf("failed to parse expected int: %s", err.Error())
				}
				want = int32(sampleDataValue)
			}

			logImport := wasmtime.WrapFunc(store, func(got int32) {
				if got != want {
					t.Fatalf("resulting wasm was incorrect, got: %d but want: %d", got, want)
				}
			})

			dummyImport := wasmtime.WrapFunc(store, func(got int32) { return })

			instance, err = wasmtime.NewInstance(store, module, []wasmtime.AsExtern{
				logImport,   // console.log
				dummyImport, // screeps.move
			})
			if err != nil {
				t.Fatalf("failed to create wasmtime instance: %s", err.Error())
			}
			rootNode := instance.GetExport(store, "root")
			result, err = rootNode.Func().Call(store)
			if err != nil {
				t.Fatalf("execution of resulting wasm failed: %s, %v", err.Error(), result)
			}
		})
	}
}

func TestCompileInvalidSourceCode(t *testing.T) {
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
