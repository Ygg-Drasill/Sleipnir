package compiler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bytecodealliance/wasmtime-go/v20"
	"log"
	"os"

	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator"
)

// Compile runs the lexer, parser and code generator
func (compiler *Compiler) Compile() error {
	var syntaxTree ast.Attribute
	var err error
	tokens := compiler.lexer.FindTokens()
	scanner := NewScanner(tokens)

	if syntaxTree, err = compiler.parser.Parse(scanner); err != nil {
		return err
	}

	programNode, ok := syntaxTree.(ast.Program)
	if !ok {
		return errors.New("root is not a program")
	}

	ctx := compiler.parser.Context.(ast.ParseContext)
	gen := generator.New(&programNode, &ctx)
	compiler.outBuffer = gen.GenWrapper()
	bytes, _ := json.MarshalIndent(syntaxTree, "", "\t")
	file, _ := os.OpenFile("AST_out.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	file.Write(bytes)
	file.Close()
	return nil
}

// writeOutputToFile writes the stored webassembly to a new file, run this after Compile
func writeOutputToFile(outputFilePath string, data []byte) {
	if data == nil {
		log.Fatal("failed to write to file: buffer is empty")
	}
	err := os.WriteFile(outputFilePath, data, 0644)
	if err != nil {
		log.Fatalf("failed to write to file: %s", err.Error())
	}
}

// WriteWatFile writes the code generated webassembly into a .wat file at the specified path
func (compiler *Compiler) WriteWatFile(outputFilePath string) {
	writeOutputToFile(outputFilePath, compiler.outBuffer.Bytes())
}

// ConvertWat2Wasm converts a .wat file into a .wasm file and writes a .wasm file at the specified path
func (compiler *Compiler) ConvertWat2Wasm(outputFilePath string) {
	compiledYgl := compiler.outBuffer
	if compiledYgl == nil {
		log.Fatal("failed to write to file: buffer is empty")
	}
	wasm, err := wasmtime.Wat2Wasm(compiledYgl.String())
	if err != nil {
		log.Fatalf("failed to convert webassembly text into compiled webassembly: %s", err.Error())
	}
	writeOutputToFile(outputFilePath, wasm)
}

// WriteOutputToBuffer copies the output webassembly text buffer to another buffer, returns the amount of bytes written
func (compiler *Compiler) WriteOutputToBuffer(buffer *bytes.Buffer) (int, error) {
	outputBuffer := new(bytes.Buffer)
	if compiler.outBuffer == nil {
		log.Fatalf("failed to write to buffer: buffer is empty")
	}

	return outputBuffer.Write(compiler.outBuffer.Bytes())
}
