package lexer

import (
	"os"
	"pgregory.net/rapid"
)

func testLexer(t rapid.T) {
	formatFile, err := os.ReadFile("../data/samples/valid/addition.ygl")
	format := string(formatFile)

	if err != nil {
		panic(err)
	}

	rapid.Check(&t, func(t *rapid.T) {

	})

}
