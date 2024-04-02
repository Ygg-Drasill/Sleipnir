package lexer

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/compiler/analysis/utils"
	"strings"
)

type StateFunction func(lexer *Lexer) StateFunction

const nonTokenRunes string = " \t\r"
const newLineRunes string = "\n"
const connector string = "->"

const commentSingle = "//"
const commentMultiStart = "/*"
const commentMultiEnd = "*/"

func matchPreprocessor(lexer *Lexer) StateFunction {
	//TODO: do preprocessing
	return matchAny
}

func matchAny(lexer *Lexer) StateFunction {
	currentRune := lexer.cursorNext()
	if utils.IsLetter(currentRune) {
		return matchLetters
	}

	if utils.IsNumber(currentRune) {
		return matchNumbers
	}

	if strings.HasPrefix(lexer.inputCode[lexer.tokenStart:], connector) {
		lexer.cursorBackup()
		return matchConnector
	}

	//TODO: logical operator
	if utils.IsOperator(currentRune) {
		lexer.serveToken(TokenOperator)
		return matchNonToken
	}

	if utils.IsPunctuation(currentRune) {
		lexer.serveToken(TokenPunctuation)
		return matchNonToken
	}

	//TODO: errorhandler
	fmt.Printf("Unrecognised token %c\n", currentRune)
	return nil
}

func matchNonToken(lexer *Lexer) StateFunction {
	for {
		currentRune := lexer.cursorNext()
		if currentRune == EOF {
			return nil
		}
		if strings.ContainsRune(nonTokenRunes, currentRune) {
			continue
		}

		if strings.ContainsRune(newLineRunes, currentRune) {
			lexer.cursorRow++
			lexer.cursorCol = 1
			continue
		}

		lexer.cursorBackup()
		lexer.cursorIgnore()

		if strings.HasPrefix(lexer.inputCode[lexer.cursor:], commentSingle) {
			return matchCommentSingle
		}

		if strings.HasPrefix(lexer.inputCode[lexer.cursor:], commentMultiStart) {
			lexer.cursorJump(len(commentMultiStart))
			return matchCommentMulti
		}

		return matchAny
	}
}

func matchLetters(lexer *Lexer) StateFunction {
	for {
		currentRune := lexer.cursorNext()
		if utils.IsNumber(currentRune) {
			return matchIdentifier
		}
		if !utils.IsLetter(currentRune) {
			lexer.cursorBackup()
			return matchKeyword
		}
	}
}

func matchNumbers(lexer *Lexer) StateFunction {
	for {
		currentRune := lexer.cursorNext()
		if utils.IsLetter(currentRune) {
			//TODO: unexpected symbol
			return nil
		}

		if !utils.IsNumber(currentRune) {
			lexer.cursorBackup()
			lexer.serveToken(TokenLiteral)
			return matchNonToken
		}
	}
}

func matchIdentifier(lexer *Lexer) StateFunction {
	for {
		currentRune := lexer.cursorNext()
		if !utils.IsLetter(currentRune) && !utils.IsNumber(currentRune) {
			lexer.cursorBackup()
			lexer.serveToken(TokenIdentifier)
			return matchNonToken
		}
	}
}

func matchKeyword(lexer *Lexer) StateFunction {
	for _, keyword := range reservedKeywords {
		if strings.HasPrefix(lexer.inputCode[lexer.tokenStart:], keyword) {
			lexer.serveToken(TokenKeyword)
			return matchNonToken
		}
	}
	return matchIdentifier
}

func matchConnector(lexer *Lexer) StateFunction {
	lexer.cursorJump(len(connector))
	lexer.serveToken(TokenConnector)
	return matchNonToken
}

func matchCommentSingle(lexer *Lexer) StateFunction {
	for {
		currentRune := lexer.cursorNext()
		if currentRune == '\n' {
			break
		}
	}
	lexer.cursorIgnore()
	lexer.cursorBackup()
	return matchNonToken
}

func matchCommentMulti(lexer *Lexer) StateFunction {
	return matchNonToken
}
