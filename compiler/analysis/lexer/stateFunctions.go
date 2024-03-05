package lexer

import "strings"

type StateFunction func(lexer *Lexer) StateFunction

var nonTokenRunes = " \n\t\r"

var matchPreprocessor StateFunction = func(lexer *Lexer) StateFunction {
	return nil
}

var matchNonToken StateFunction = func(lexer *Lexer) StateFunction {
	for {
		for nextRune := lexer.cursorNext(); nextRune != EOF; {
			if !strings.ContainsRune(nonTokenRunes, nextRune) {
				break
			}
		}
		//switch state :)
		break
	}
	return nil
}

var matchNode StateFunction = func(lexer *Lexer) StateFunction {
	return nil
}
