package analysis

import (
	"testing"

	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
)

func TestFindTokens(t *testing.T) {
	lexer := NewLexerFromString("../../testData/snippets/testFindTokens.ygl") //Todo: non-existent

	//stupid test
	gotTokens := lexer.FindTokens()
	wantedTokens := []Token{} //Todo: fill out array

	for i := range wantedTokens {
		got := gotTokens[i]
		want := wantedTokens[i]
		if got != want {
			t.Errorf("Got %q, expected %q", got, want)
		}
	}
}
