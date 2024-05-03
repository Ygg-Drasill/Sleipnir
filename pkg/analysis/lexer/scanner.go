package lexer

import "github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"

type Scanner struct {
	tokens []*token.Token
}

func NewScanner(tokens []*token.Token) *Scanner {
	return &Scanner{tokens: tokens}
}

func (s *Scanner) Scan() *token.Token {
	t := s.tokens[0]
	s.tokens = s.tokens[1:]
	return t
}
