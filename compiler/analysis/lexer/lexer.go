package lexer

import (
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
	"os"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	inputCode     string
	inputLength   int
	tokenStart    int
	lastRuneWidth int
	cursor        int
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
	return rune
}

func (lexer *Lexer) cursorIgnore() {
	lexer.tokenStart = lexer.cursor
}

func (lexer *Lexer) cursorBackup() {
	lexer.cursor -= lexer.lastRuneWidth
}

func (lexer *Lexer) cursorPeek() rune {
	nextRune := lexer.cursorNext()
	lexer.cursorBackup()
	return nextRune
}

func (lexer *Lexer) cursorJump(length int) {
	lexer.cursor += length
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

	if tokenType == TokenConnector ||
		tokenType == TokenPunctuation ||
		tokenType == TokenKeyword {
		if value == ";" {
			tokType = token.TokMap.Type("stmtEnd")
		} else {
			tokType = token.TokMap.Type(value)
		}
	}

	if tokenType == TokenOperator {
		if value == "=" {
			tokType = token.TokMap.Type("assignOp")
		}

		if value == "==" || value == "!=" ||
			value == "<" || value == ">" {
			tokType = token.TokMap.Type("compOp")
		}

		if value == "&&" || value == "||" {
			tokType = token.TokMap.Type("logicOp")
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
		Pos:  token.Pos{},
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
