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

		//formatFile, err := os.ReadFile("../data/samples/valid/addition.ygl")
		//format := string(formatFile)
		//if err != nil { panic(err)	}
		//rapid.Check(&t, func(t *rapid.T) {	})

		//variable generated as input for the lexer
		gen := rapid.StringMatching(`([a-zA-Z]+{4}) ([ ]) ([a-zA-Z]+{4})
										 ([ ])
										 ([a-zA-Z]+{1}) ([ ]) ([a-zA-Z]+{1})`).Draw(t, "gen")

		// Creation of the lexer with the generated input
		l := lexer.NewLexerFromString(gen)
		tokens := l.FindTokens()
		var b = len(tokens)
		if b < 1 {
			t.Fatalf("")
		}

		scanner := lexer.NewScanner(tokens)

		// Asserts for expected and actual token

		// Parser
		p := parser.NewParser()
		p.Context = ast.NewParseContext()
		_, err := p.Parse(scanner)

		if len(gen) < 1 && err == nil {
			t.Fatalf("parsed empty token?: %s", gen)
		}

	})
}
