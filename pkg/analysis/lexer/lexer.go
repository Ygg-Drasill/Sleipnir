package lexer

import (
	"os"
	"unicode"
	"unicode/utf8"

	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
)

type Lexer struct {
	inputCode     string
	inputLength   int
	tokenStart    int
	lastRuneWidth int
	cursor        int
	cursorCol     int
	cursorRow     int
	tokens        chan *token.Token
	tokenList     []*token.Token
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
		cursorCol:   1,
		cursorRow:   1,
		tokens:      make(chan *token.Token),
		tokenList:   make([]*token.Token, 0),
	}
}

func (lexer *Lexer) cursorNext() (rune rune) {
	if lexer.cursor > lexer.inputLength {
		return -1
	}
	rune, lexer.lastRuneWidth = utf8.DecodeRuneInString(lexer.inputCode[lexer.cursor:])
	lexer.cursor += lexer.lastRuneWidth
	lexer.cursorCol++
	return rune
}

func (lexer *Lexer) cursorIgnore() {
	lexer.tokenStart = lexer.cursor
}

func (lexer *Lexer) cursorBackup() {
	lexer.cursor -= lexer.lastRuneWidth
	lexer.cursorCol--
}

func (lexer *Lexer) cursorPeek() rune {
	nextRune := lexer.cursorNext()
	lexer.cursorBackup()
	return nextRune
}

func (lexer *Lexer) cursorJump(length int) {
	lexer.cursor += length
	lexer.cursorCol += length
}

func (lexer *Lexer) Position() token.Pos {
	return token.Pos{
		Line:   lexer.cursorRow,
		Column: lexer.cursorCol - (lexer.cursor - lexer.tokenStart),
	}
}

func (lexer *Lexer) serveToken(tokenType TokenType) {
	value := lexer.inputCode[lexer.tokenStart:lexer.cursor]
	var tokType token.Type

	if tokenType == TokenIdentifier {
		if unicode.IsUpper(rune(value[0])) {
			tokType = token.TokMap.Type("nodeId")
		} else {
			tokType = token.TokMap.Type("varId")
		}
	}

	if tokenType == TokenPunctuation {
		tokType = token.TokMap.Type(PunctuationMap[value])
	}

	if tokenType == TokenConnector ||
		tokenType == TokenKeyword {
		tokType = token.TokMap.Type(value)
	}

	if tokenType == TokenOperator {
		tokType = token.TokMap.Type(OperatorMap[value])

		if value == "=" {
			tokType = token.TokMap.Type("assign")
		}
	}

	if tokenType == TokenLiteral {
		tokType = token.TokMap.Type("int64")
	}

	if tokenType == TokenEOF {
		tokType = token.TokMap.Type("␚")
	}

	//TODO: dont serve empty tokens
	//lexer.tokens <- token
	lexer.tokenList = append(lexer.tokenList, &token.Token{
		Type: tokType,
		Lit:  []byte(value),
		Pos:  lexer.Position(),
	})
	lexer.tokenStart = lexer.cursor
}

func (lexer *Lexer) FindTokens() []*token.Token {
	for matchState := matchNonToken; matchState != nil; {
		matchState = matchState(lexer)
	}
	lexer.serveToken(TokenEOF)
	close(lexer.tokens)
	return lexer.tokenList
}
