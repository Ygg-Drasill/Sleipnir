package compiler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/bytecodealliance/wasmtime-go"
	"log"
	"os"

	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator"
)

// Compile runs the lexer, parser and code generator
func (compiler *Compiler) Compile() error {
	tokens := compiler.lexer.FindTokens()
	scanner := NewScanner(tokens)

	syntaxTree, err := compiler.getSyntaxTree(scanner)
	if err != nil {
		return err
	}

	programNode, ok := syntaxTree.(ast.Program)
	if !ok {
		return errors.New("root is not a program")
	}

	ctx := compiler.parser.Context.(ast.ParseContext)
	gen := generator.New(&programNode, &ctx)
	compiler.syntaxTree = &syntaxTree
	compiler.outBuffer = gen.GenWrapper()
	return nil
}

func (compiler *Compiler) getSyntaxTree(scanner *Scanner) (ast.Attribute, error) {
	var syntaxTree ast.Attribute
	var err error
	if syntaxTree, err = compiler.parser.Parse(scanner); err != nil {
		return nil, err
	}
	return syntaxTree, nil
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

func (compiler *Compiler) WriteJsonFile(outputFilePath string) {
	syntaxTree := compiler.syntaxTree
	jsonBytes, err := json.MarshalIndent(syntaxTree, "", "\t")
	if err != nil {
		log.Fatalf("failed to return converted json with indents: %s", err.Error())
	}
	file, err := os.OpenFile(outputFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("failed to create file: %s", err.Error())
	}
	_, err = file.Write(jsonBytes)
	if err != nil {
		log.Fatalf("failed to write json: %s", err.Error())
	}
	err = file.Close()
	if err != nil {
		log.Fatalf("failed to close file: %s", err.Error())
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
func (compiler *Compiler) WriteOutputToBuffer(outputBuffer *bytes.Buffer) (int, error) {
	if compiler.outBuffer == nil {
		log.Fatalf("failed to write to buffer: buffer is empty")
	}

	return outputBuffer.Write(compiler.outBuffer.Bytes())
}
