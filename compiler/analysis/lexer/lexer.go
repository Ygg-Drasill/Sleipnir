package lexer

import (
	"os"
	"unicode/utf8"
)

type Lexer struct {
	inputCode     string
	inputLength   int
	tokenStart    int
	lastRuneWidth int
	cursor        int
	tokens        chan Token
	tokenList     []Token
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
		tokenList:   make([]Token, 0),
	}
}

func (lexer *Lexer) cursorNext() (rune rune) {
	if lexer.cursor > lexer.inputLength {
		return -1
	}
	rune, lexer.lastRuneWidth = utf8.DecodeRuneInString(lexer.inputCode[lexer.cursor:])
	lexer.cursor += lexer.lastRuneWidth
	return rune
}

func (lexer *Lexer) cursorIgnore() {
	lexer.tokenStart = lexer.cursor
}

func (lexer *Lexer) cursorBackup() {
	lexer.cursor -= lexer.lastRuneWidth
}

func (lexer *Lexer) cursorPeek() rune {
	rune := lexer.cursorNext()
	lexer.cursorBackup()
	return rune
}

func (lexer *Lexer) serveToken(token Token) {
	lexer.tokens <- token
}

func (lexer *Lexer) FindTokens() []Token {
	tempTokens := make([]Token, 0)
	for matchState := matchNonToken; matchState != nil; {
		matchState = matchState(lexer)
	}
	close(lexer.tokens)
	return tempTokens
}
