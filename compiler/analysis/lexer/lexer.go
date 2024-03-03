package lexer

import (
	"os"
)

type Lexer struct {
	inputCode   string
	inputLength int
	tokenStart  int
	cursor      int
	tokens      chan Token
}

func NewLexerFromString(inputPath string) *Lexer {
	file, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	input := string(file)
	return &Lexer{
		inputCode:   input,
		inputLength: len(input),
		tokenStart:  0,
		cursor:      0,
		tokens:      make(chan Token),
	}
}

func (lexer *Lexer) FindTokens() []Token {
	tempTokens := make([]Token, 0)
	for range lexer.inputCode {
		if lexer.cursor > lexer.inputLength-1 {
			break
		}

		var token *Token
		lexer.MatchToken(MatchWhitespace)
		token = lexer.MatchToken(MatchPunctuation)
		if token != nil {
			tempTokens = append(tempTokens, *token)
			continue
		}
		token = lexer.MatchToken(MatchArrowOperator)
		if token != nil {
			tempTokens = append(tempTokens, *token)
			continue
		}
		token = lexer.MatchToken(MatchOperator)
		if token != nil {
			tempTokens = append(tempTokens, *token)
			continue
		}
		lexer.cursor++
		//run matchers ...
	}
	return tempTokens
}

func (lexer *Lexer) MatchToken(matcher TokenMatcher) *Token {
	characterCount, token := matcher(lexer)
	if characterCount > 0 {
		//lexer.serveToken(*token) Uncomment when tokens are consumed
		lexer.cursor += characterCount
		return token
	}
	return nil
}

func (lexer *Lexer) serveToken(token Token) {
	lexer.tokens <- token
}
