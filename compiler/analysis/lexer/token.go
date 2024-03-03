package lexer

import "fmt"

type TokenType = int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenIdentifier
	TokenKeyword
	TokenOperator
	TokenPunctuation
	TokenLiteral
)

type Token struct {
	tokenType TokenType
	value     string
}

func NewToken(tokenType TokenType, value string) *Token {
	return &Token{
		tokenType: tokenType,
		value:     value,
	}
}

func (token *Token) ToString() string {
	switch token.tokenType {
	case TokenError:
		return token.value
	}

	if len(token.value) > 10 {
		return fmt.Sprintf("%.10q...", token.value)
	}
	return fmt.Sprintf("%q", token.value)
}
