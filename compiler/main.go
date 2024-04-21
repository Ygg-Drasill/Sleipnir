package main

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	ost "github.com/Ygg-Drasill/Sleipnir/compiler/ast"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"github.com/Ygg-Drasill/Sleipnir/compiler/synthesis"
)

func main() {
	var filePath string = os.Args[1]
	lexer := NewLexerFromString(filePath)
	tokens := lexer.FindTokens()
	scanner := NewScanner(tokens)
	p := parser.NewParser()
	p.Context = ost.NewParseContext()
	var ast interface{}
	var err error
	if ast, err = p.Parse(scanner); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ast)
	bytes, _ := json.MarshalIndent(ast, "", "\t")
	file, _ := os.OpenFile("AST_out.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if prog, ok := ast.(ost.Program); ok {
		test := synthesis.GenWrapper(&prog)
		fmt.Printf("%s", test.String())
		files, err := os.OpenFile("test.wat", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
		file.Write(bytes)
		file.Close()
		files.Write(test.Bytes())
		files.Close()
		bytes, _ := json.MarshalIndent(ast, "", "\t")
		file, _ := os.OpenFile("AST_out.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		file.Write(bytes)
		file.Close()

	}
}
