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
	if _, e := p.Parse(scanner); e != nil {
		fmt.Println(e.Error())
	}
	//fmt.Printf("Tokens found %d, %q\n", len(tokens), tokens)
}
