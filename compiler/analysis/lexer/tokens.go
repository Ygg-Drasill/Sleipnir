package lexer

import "fmt"

type TokenType = int

type Token struct {
	tokenType TokenType
	value     string
}

const (
	TokenError TokenType = iota
	TokenEOF
)

func (token *Token) String() string {
	switch token.tokenType {
	case TokenError:
		return token.value
	}

	if len(token.value) > 10 {
		return fmt.Sprintf("%.10q...", token.value)
	}
	return fmt.Sprintf("%q", token.value)
}
