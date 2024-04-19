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
	if res, e := p.Parse(scanner); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(res)
		bytes, _ := json.MarshalIndent(res, "", "\t")
		file, _ := os.OpenFile("AST_out.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		file.Write(bytes)
		file.Close()
	}
}
