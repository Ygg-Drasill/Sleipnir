package cmd

import (
	"fmt"
	"github.com/bytecodealliance/wasmtime-go"
	"log"
	"os"
)

func runBinary(binary []byte) (any, error) {
	module := new(wasmtime.Module)
	engine := wasmtime.NewEngine()
	instance := new(wasmtime.Instance)
	store := wasmtime.NewStore(engine)

	var result interface{}
	var err error

	module, err = wasmtime.NewModule(engine, binary)
	if err != nil {
		log.Fatalf("failed to convert wasm to module: %s", err.Error())
	}

	logImport := wasmtime.WrapFunc(store, func(value int32) {
		fmt.Println("print: ", value)
	})

	var memory int32 = 0

	setImport := wasmtime.WrapFunc(store, func(value int32) {
		memory = value
	})

	getImport := wasmtime.WrapFunc(store, func() int32 {
		return memory
	})

	dummyImport := wasmtime.WrapFunc(store, func(value int32) { return })

	instance, err = wasmtime.NewInstance(store, module, []wasmtime.AsExtern{
		logImport,   // console.log
		dummyImport, // screeps.move
		setImport,   // screeps.set
		getImport,   // screeps.get
	})

	if err != nil {
		log.Fatalf("failed to create wasmtime instance: %s", err.Error())
	}
	rootNode := instance.GetExport(store, "root")
	result, err = rootNode.Func().Call(store)
	return result, err
}

func readBinaryFromFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
