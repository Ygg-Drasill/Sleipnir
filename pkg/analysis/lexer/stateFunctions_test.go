package lexer

import (
	"pgregory.net/rapid"
	"testing"
)

func TestStateFunction_matchNumber(t *testing.T) {

	// To perform multi tests
	rapid.Check(t, func(t *rapid.T) {

		// Test for letters using matchNumbers
		genLetter := rapid.StringMatching("[A-Za-z]+").Draw(t, "genLetter")
		lLetter := NewFromString(genLetter)
		matchNumbers(lLetter)
		if len(lLetter.tokenList) > 0 {
			t.Fatalf("Tokens were made, Tokenlist: %v ", lLetter.tokenList)
		}

		// Test for Numbers using matchNumbers
		genNumber := rapid.StringMatching("[0-9]+").Draw(t, "genNumber")
		lNumber := NewFromString(genNumber)
		matchNumbers(lNumber)
		if len(lNumber.tokenList) == 0 {
			t.Fatalf("No tokens were made, Tokenlist: %v ", lNumber.tokenList)
		}
	})
}
