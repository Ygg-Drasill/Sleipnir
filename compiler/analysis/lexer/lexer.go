package lexer

import "os"

type Lexer struct {
	inputCode  string
	tokenStart int
	cursor     int
	tokens     chan Token
}

func NewLexerFromFile(inputPath string) *Lexer {
	file, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	return &Lexer{
		inputCode:  string(file),
		tokenStart: 0,
		cursor:     0,
		tokens:     make(chan Token),
	}
}

func (lexer *Lexer) FindTokens() {
	for _, b := range lexer.inputCode {
		print(string(b))
	}
}

func (lexer *Lexer) serveToken(token Token) {
	lexer.tokens <- token
}
