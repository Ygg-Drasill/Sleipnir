package lexer

import (
	"fmt"
	"pgregory.net/rapid"
	"testing"
)

func TestStateFunction(t *testing.T) {

	rapid.Check(t, func(t *rapid.T) {

		//gen := rapid.StringMatching("[ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvxyz,*¨?`´<>|€$£@!(=){}1234567890]+").Draw(t, "gen")
		gen := rapid.StringMatching("[ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvxyz1234567890]+").Draw(t, "gen")
		l := NewFromString(gen)
		matchLetters(l)

		fmt.Printf("%v \n", l.tokenList)

		// matchNonToken returnes: matchCommentSingle, matchCommentMulti and matchAny
		//matchNonTokenvariable := matchNonToken(l)

	})
}

//The scanner is created
//_ = compiler.NewScanner(tokens)
//fmt.Printf("%s %s \n", tokens, scanner)
// Asserts for expected and actual token

// Do not know if this is relevant for the test for the lexer
// Parser
/*
	p := parser.NewParser()
	p.Context = ast.NewParseContext()
	_, err := p.Parse(scanner)
	if len(gen) < 1 && err == nil {
		t.Fatalf("parsed empty token?: %s", gen)
	}
*/
