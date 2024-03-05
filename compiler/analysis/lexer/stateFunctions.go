package lexer

import "strings"

type StateFunction func(lexer *Lexer) StateFunction

var nonTokenRunes = " \n\t\r"

// TODO: instead, glide over alphanumeric word until punctuation or whitespace, then check inputCode[lexer.tokenStart : lexer.cursor] == keyword
func isKeyword(lexer *Lexer, keyword string) bool {
	currentRune := lexer.cursorNext()
	for _, r := range keyword[0:] {
		if r != currentRune {
			return false
		}
		currentRune = lexer.cursorNext()
	}
	return true
}

func matchPreprocessor(lexer *Lexer) StateFunction {
	return nil
}

func matchNonToken(lexer *Lexer) StateFunction {

	for {
		currentRune := lexer.cursorNext()
		if currentRune == EOF {
			break
		}
		if strings.ContainsRune(nonTokenRunes, currentRune) {
			continue
		}

		switch currentRune {
		case 'n':
			return matchNodeKeyword
		}
	}
	return nil
}

func matchNodeKeyword(lexer *Lexer) StateFunction {
	lexer.cursorBackup()
	if isKeyword(lexer, "node") {
		lexer.serveToken(TokenKeyword)
		return matchNonToken
	}
	return nil
}
