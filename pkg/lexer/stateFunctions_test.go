package lexer

import (
	"fmt"
	"pgregory.net/rapid"
	"testing"
)

func TestStateFunction_matchNumber(t *testing.T) {

	rapid.Check(t, func(t *rapid.T) {

		// Test for letters using matchNumbers
		genLetter := rapid.StringMatching("[A-Za-z]+").Draw(t, "genLetter")
		lLetter := NewFromString(genLetter)
		matchNumbers(lLetter)
		if len(lLetter.tokenList) > 0 {
			t.Fatalf("TestStateFunction_matchNumber: Tokens were made, Tokenlist: %v ", lLetter.tokenList)
		}

		// Test for Numbers using matchNumbers
		genNumber := rapid.StringMatching("[0-9]+").Draw(t, "genNumber")
		lNumber := NewFromString(genNumber)
		matchNumbers(lNumber)
		if len(lNumber.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchNumber: No tokens were made, Tokenlist: %v ", lNumber.tokenList)
		}
	})
}

func TestStateFunction_matchIdentifier(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		//Failed characters: _^¨
		//Test for matchIdentifier
		genIdentifier := rapid.StringMatching(`[[:alnum:]]+`).Draw(t, "genIdentifier")
		lIdentifier := NewFromString(genIdentifier)
		//fmt.Printf("Generated string: %v \n", genIdentifier)
		matchIdentifier(lIdentifier)
		//fmt.Printf("lIdentifier: %v len: %v \n", lIdentifier.tokenList, len(lIdentifier.tokenList))
		if len(lIdentifier.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchIdentifier: Tokens were made, Tokenlist: %v ", lIdentifier.tokenList)
		}
	})
}

// !DO NOT WORK!
func TestStateFunction_matchLetters(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		genLetters := rapid.StringMatching(`node hey {}`).Draw(t, "genLetters")
		fmt.Printf("%v \n", genLetters)
		lLetter := NewFromString(genLetters)
		matchLetters(lLetter)

		fmt.Printf("lLetter: %v len: %v \n", lLetter.tokenList, len(lLetter.tokenList))

		if len(lLetter.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchLetters: Tokens were made, Tokenlist: %v ", lLetter.tokenList)
		}
	})
}

func TestStateFunction_matchKeyword(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

		//Test for matchKeyword for keywords
		genKeyword := rapid.SampledFrom(reservedKeywords).Draw(t, "genKeyword")
		lKeyword := NewFromString(genKeyword)
		matchKeyword(lKeyword)
		if len(lKeyword.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchIdentifier: No tokens were made, Tokenlist: %v ", lKeyword.tokenList)
		}

		//Test for matchKeyword for non-keyword
		genNonKeyword := rapid.StringMatching(`[[:alnum:]]+`).Draw(t, "genNonKeyword")
		lNonKeyword := NewFromString(genNonKeyword)
		matchKeyword(lNonKeyword)
		if len(lNonKeyword.tokenList) > 0 {
			t.Fatalf("TestStateFunction_matchIdentifier: Tokens were made, Tokenlist: %v ", lNonKeyword.tokenList)
		}
	})
}

func TestStateFunction_matchConnector(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

	})
}

func TestStateFunction_matchCommentSingle(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

	})
}

func TestStateFunction_matchCommentMulti(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

	})
}

func TestStateFunction_matchAny(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {

	})
}
