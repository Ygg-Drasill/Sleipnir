package main

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"github.com/Ygg-Drasill/Sleipnir/compiler/ast"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"github.com/Ygg-Drasill/Sleipnir/compiler/synthesis"
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
		if prog, ok := res.(ast.Program); ok {
			test := synthesis.GenWrapper(&prog)
			fmt.Printf("%s", test.String())
			files, err := os.OpenFile("test.wat", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				panic(err)
			}
			files.Write(test.Bytes())
			files.Close()
			bytes, _ := json.MarshalIndent(res, "", "\t")
			file, _ := os.OpenFile("AST_out.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			file.Write(bytes)
			file.Close()
		} else {
			fmt.Println("Parse result is not of type ast.Program")
		}
	}
}
