package lexer

import (
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
		matchIdentifier(lIdentifier)
		if len(lIdentifier.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchIdentifier: Tokens were made, Tokenlist: %v ", lIdentifier.tokenList)
		}
	})
}

func TestStateFunction_matchLetters(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Test for letters using matchLetters
		genLetters := rapid.StringMatching(`[[:alpha:]]+`).Draw(t, "genLetters")
		lLetter := NewFromString(genLetters)
		matchLetters(lLetter)
		if len(lLetter.tokenList) > 0 {
			t.Fatalf("TestStateFunction_matchLetters: Tokens were made, Tokenlist: %v ", lLetter.tokenList)
		}
		// Test for letters using matchLetters & matchIdentifier
		matchIdentifier(lLetter)
		if len(lLetter.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchLetters_matchIdentifier: No tokens were made, Tokenlist: %v ", lLetter.tokenList)
		}
		// Test for letters using matchLetters & matchKeyword
		matchKeyword(lLetter)
		if len(lLetter.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchLetters_matchKeyword: No tokens were made, Tokenlist: %v ", lLetter.tokenList)
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
			t.Fatalf("TestStateFunction_matchKeyword: No tokens were made, Tokenlist: %v ", lKeyword.tokenList)
		}
		//Test for matchKeyword for non-keyword
		genNonKeyword := rapid.StringMatching(`[[:alnum:]]+`).Draw(t, "genNonKeyword")
		lNonKeyword := NewFromString(genNonKeyword)
		matchKeyword(lNonKeyword)
		if len(lNonKeyword.tokenList) > 0 {
			t.Fatalf("TestStateFunction_matchKeyword: Tokens were made, Tokenlist: %v ", lNonKeyword.tokenList)
		}
	})
}

func TestStateFunction_matchConnector(t *testing.T) {
	/*		Notes for match connector:
			- You can make the connection symbol infinitly long fx. "--------->
			- Any combination of -> is accepted fx. >--->-, >>->>- or --    	*/
	rapid.Check(t, func(t *rapid.T) {
		// Test for matchConnector for connector
		genMatchConnector := rapid.StringMatching(`([->]{2})+`).Draw(t, "genMatchConnector")
		lMatchConnector := NewFromString(genMatchConnector)
		matchConnector(lMatchConnector)
		if len(lMatchConnector.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchConnector: Tokens were made, Tokenlist: %v ", lMatchConnector.tokenList)
		}
	})
}

func TestStateFunction_matchCommentSingle(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		//Test for matchCommentSingle for single comment
		genMatchCommentSingle := rapid.StringMatching(`\n//`).Draw(t, "genMatchCommentSingle")
		lMatchCommentSingle := NewFromString(genMatchCommentSingle)
		matchConnector(lMatchCommentSingle)
		if len(lMatchCommentSingle.tokenList) == 0 {
			t.Fatalf("TestStateFunction_matchCommentSingle: Tokens were made, Tokenlist: %v ", lMatchCommentSingle.tokenList)
		}
	})
}

func TestStateFunction_matchCommentMulti(t *testing.T) {
	// Test for Multi comment is not perform, since the statement function is not finished
	rapid.Check(t, func(t *rapid.T) {})
}

func TestStateFunction_matchNonToken(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		//Test matchNonToken
		genMatchNonToken := rapid.StringMatching(`[[:word:]//[:punct:]]+`).Draw(t, "genMatchNonToken")
		lMatchNonToken := NewFromString(genMatchNonToken)
		matchNonToken(lMatchNonToken)
		if len(lMatchNonToken.tokenList) > 0 {
			t.Fatalf("TestStateFunction_matchNonToken: Tokens were made, Tokenlist: %v ", lMatchNonToken.tokenList)
		}
	})
}

func TestStateFunction_matchAny(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		//Note: Not really sure how to test it properly
		//Test MatchAny for large input
		genMatchAny := rapid.StringMatching(`[[:graph:]]+`).Draw(t, "genMatchAny")
		lMatchAny := NewFromString(genMatchAny)
		matchAny(lMatchAny)
		//Test MatchAny for unrecognised token
		genMatchAny_unrecognised := rapid.StringMatching(`\D`).Draw(t, "genMatchAny_unrecognised")
		lMatchAny_unrecognised := NewFromString(genMatchAny_unrecognised)
		matchAny(lMatchAny_unrecognised)

	})
}
