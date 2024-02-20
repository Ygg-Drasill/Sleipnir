package main

import (
	. "github.com/Ygg-Drasill/Sleipnir/compiler/lexer"
	"os"
)

func main() {
	var filePath string = os.Args[1]
	lexer := NewLexerFromFile(filePath)
	lexer.FindTokens()
}
