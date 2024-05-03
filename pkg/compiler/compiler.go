package compiler

import (
	"log/slog"
	"os"

	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/synthesis"
)

// Compile runs the lexer, parser and code generator
func (compiler *Compiler) Compile() {
	var syntaxTree ast.Attribute
	var err error
	tokens := compiler.lexer.FindTokens()
	scanner := NewScanner(tokens)

	if syntaxTree, err = compiler.parser.Parse(scanner); err != nil {
		slog.Error("failed to parse: %s", err.Error())
		return
	}

	if programNode, ok := syntaxTree.(ast.Program); ok {
		compiler.outBuffer = synthesis.GenWrapper(&programNode)
	}
}

// WriteOutputToFile writes the stored webassembly to a new file, run this after Compile
func (compiler *Compiler) WriteOutputToFile(outputFilePath string) {
	if compiler.outBuffer == nil {
		slog.Error("failed to write to file: buffer is empty")
	}

	err := os.WriteFile(outputFilePath, compiler.outBuffer.Bytes(), 0644)
	if err != nil {
		slog.Error("failed to write webassembly to file: %s", err.Error())
	}
}
