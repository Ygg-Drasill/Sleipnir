package lexer

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/compiler/analysis/utils"
	"strings"
)

type StateFunction func(lexer *Lexer) StateFunction

var nonTokenRunes = " \n\t\r"

func matchPreprocessor(lexer *Lexer) StateFunction {
	//TODO: do preprocessing
	return matchNonToken
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
		lexer.cursorBackup()
		lexer.cursorIgnore()

		if utils.IsLetter(currentRune) {
			return matchLetters
		}

		if utils.IsNumber(currentRune) {
			return matchNumbers
		}

		//TODO: logical operator
		if utils.IsOperator(currentRune) {
			lexer.cursorNext()
			lexer.serveToken(TokenOperator)
			return matchNonToken
		}

		if utils.IsPunctuation(currentRune) {
			lexer.cursorNext()
			lexer.serveToken(TokenPunctuation)
			return matchNonToken
		}

		//TODO: errorhandler
		fmt.Printf("Unrecognised token %c\n", currentRune)
		return nil
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
