package analysis

import (
	"bytes"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
	"testing"

	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
)

func compareToken(a *token.Token, b *token.Token) bool {
	if a.Type != b.Type {
		return false
	}

	if bytes.Compare(a.Lit, b.Lit) != 0 {
		return false
	}

	if a.Line != b.Line {
		return false
	}

	if a.Column != b.Column {
		return false
	}

	return true
}

func TestFindTokens(t *testing.T) {
	lexer := NewLexerFromString("../../testData/snippets/testFindTokens.ygl") //Todo: non-existent

	//stupid test
	gotTokens := lexer.FindTokens()
	wantedTokens := []*token.Token{
		{token.TokMap.Type("node"), []byte("node"), token.Pos{Line: 1, Column: 1}},
		{token.TokMap.Type("nodeId"), []byte("A"), token.Pos{Line: 1, Column: 6}},
		{token.TokMap.Type("{"), []byte("{"), token.Pos{Line: 1, Column: 8}},
		{token.TokMap.Type("}"), []byte("}"), token.Pos{Line: 3, Column: 1}},
		{token.TokMap.Type("node"), []byte("node"), token.Pos{Line: 5, Column: 1}},
		{token.TokMap.Type("nodeId"), []byte("B"), token.Pos{Line: 5, Column: 6}},
		{token.TokMap.Type("{"), []byte("{"), token.Pos{Line: 5, Column: 8}},
		{token.TokMap.Type("}"), []byte("}"), token.Pos{Line: 7, Column: 1}},
	}

	for i := range wantedTokens {
		got := gotTokens[i]
		want := wantedTokens[i]
		if !compareToken(got, want) {
			t.Errorf("Got %q, expected %q", got, want)
		}
	}
}
