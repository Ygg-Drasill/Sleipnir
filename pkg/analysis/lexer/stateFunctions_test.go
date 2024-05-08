package lexer

import (
	"pgregory.net/rapid"
	"testing"
)

func TestStateFunction(t *testing.T) {

	// To perform multi tests
	rapid.Check(t, func(t *rapid.T) {

		// Test for letters using matchNumbers
		genLetter := rapid.StringMatching("[ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvxyz]+").Draw(t, "genLetter")
		l := NewFromString(genLetter)
		v := matchNumbers(l)
		tokenListLength := len(l.tokenList)
		if v == nil {
			t.Fatalf("No tokens was made, Tokenlist: %v length: %d ", l.tokenList, tokenListLength)
		}

	})
}

/*
	for i, token := range l.tokenList {
		fmt.Printf("index:%d tokenType:%v tokenvalue:%v \n", i, token.Type, token.IDValue())
	}
*/
