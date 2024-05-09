package compiler

import (
	"bytes"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/parser"
	"github.com/Ygg-Drasill/Sleipnir/pkg/lexer"
	"github.com/Ygg-Drasill/Sleipnir/utils"
	"log/slog"
	"os"
)

// Compiler is used to transpile .ygl source files to wasm
type Compiler struct {
	lexer     *lexer.Lexer
	parser    *parser.Parser
	outBuffer *bytes.Buffer
}

func NewFromFile(filePath string) *Compiler {
	var err error
	var sourceIsValid bool
	var source []byte

	sourceIsValid, err = utils.ValidateYglFilePath(filePath)
	if err != nil {
		slog.Error("failed to validate source file: %s\n", err.Error())
		return nil
	}
	if !sourceIsValid {
		slog.Error("source file is invalid: %s\n", err.Error())
		return nil
	}
	source, err = os.ReadFile(filePath)
	return NewFromString(string(source))
}

func NewFromString(source string) *Compiler {
	newLexer := lexer.NewFromString(source)
	newParser := parser.NewParser()
	newParser.Context = ast.NewParseContext()

	return &Compiler{
		lexer:     newLexer,
		parser:    newParser,
		outBuffer: nil,
	}
}
