package main

import (
	"fmt"
	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"os"
)

func main() {
	var filePath string = os.Args[1]
	lexer := NewLexerFromString(filePath)
	tokens := lexer.FindTokens()
	fmt.Printf("Tokens found %d, %q\n", len(tokens), tokens)
}
