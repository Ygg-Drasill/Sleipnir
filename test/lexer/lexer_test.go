package lexer

import (
	"github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"github.com/Ygg-Drasill/Sleipnir/compiler/ast"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"pgregory.net/rapid"
	"testing"
)

func TestLexer(t *testing.T) {

	rapid.Check(t, func(t *rapid.T) {

		//variable generated as input for the lexer
		gen := rapid.StringMatching(`([a-zA-Z]{4})([ ])([a-zA-Z]{4})([ ])([{]{1})([ ])([}]{1})`).Draw(t, "gen")
		//fmt.Println(gen)

		// Creation of the lexer with the generated input
		l := lexer.NewLexerFromString(gen)
		tokens := l.FindTokens()
		//fmt.Printf("%s", tokens)

		// Checks if there is any tokens
		var tokenLen = len(tokens)
		if tokenLen < 1 {
			t.Fatalf("No tokens made")
		}

		//The scanner is created
		scanner := lexer.NewScanner(tokens)

		// Asserts for expected and actual token

		// Do not know if this is relevant for the test for the lexer
		// Parser
		p := parser.NewParser()
		p.Context = ast.NewParseContext()
		_, err := p.Parse(scanner)
		if len(gen) < 1 && err == nil {
			t.Fatalf("parsed empty token?: %s", gen)
		}

	})
}
