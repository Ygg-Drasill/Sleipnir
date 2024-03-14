package analysis

import (
	"testing"

	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
)

func TestFindTokens(t *testing.T) {
	lexer := NewLexerFromString("../../testData/snippets/testFindTokens.ygl") //Todo: non-existent

	//stupid test
	gotTokens := lexer.FindTokens()
	wantedTokens := []Token{
		NewToken(TokenKeyword, "node"),
		NewToken(TokenIdentifier, "A"),
		NewToken(TokenPunctuation, "{"),
		NewToken(TokenPunctuation, "}"),
		NewToken(TokenKeyword, "node"),
		NewToken(TokenIdentifier, "B"),
		NewToken(TokenPunctuation, "{"),
		NewToken(TokenPunctuation, "}"),
	}

	for i := range wantedTokens {
		got := gotTokens[i]
		want := wantedTokens[i]
		if got != want {
			t.Errorf("Got %q, expected %q", got, want)
		}
	}
}
