package main

import (
	"encoding/json"
	"fmt"
	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"github.com/Ygg-Drasill/Sleipnir/compiler/ast"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"os"
)

func main() {
	var filePath string = os.Args[1]
	lexer := NewLexerFromString(filePath)
	tokens := lexer.FindTokens()
	scanner := NewScanner(tokens)
	p := parser.NewParser()
	p.Context = ast.NewParseContext()
	var ast interface{}
	var err error
	if ast, err = p.Parse(scanner); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ast)
	bytes, _ := json.MarshalIndent(ast, "", "\t")
	file, _ := os.OpenFile("AST_out.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	file.Write(bytes)
	file.Close()
}
