package main

import (
	"fmt"
	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"os"
)

func main() {
	var filePath string = os.Args[1]
	lexer := NewLexerFromString(filePath)
	tokens := lexer.FindTokens()
	scanner := NewScanner(tokens)
	p := parser.NewParser()
	if res, e := p.Parse(scanner); e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(res)
	}
}
