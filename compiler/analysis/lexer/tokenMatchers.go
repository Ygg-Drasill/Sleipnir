package lexer

type TokenMatcher func(lexer *Lexer) (int, *Token)

var whitespaceCharacters = [3]rune{
	' ',
	'\n',
	'\t',
}

func isWhitespace(r rune) bool {
	for _, v := range whitespaceCharacters {
		if r == v {
			return true
		}
	}
	return false
}

var MatchWhitespace TokenMatcher = func(lexer *Lexer) (int, *Token) {
	for isWhitespace(rune(lexer.inputCode[lexer.cursor])) {
		lexer.cursor++
	}
	return 0, nil
}

var MatchPunctuation TokenMatcher = func(lexer *Lexer) (int, *Token) {
	value := lexer.inputCode[lexer.cursor]
	if value == ';' || value == '.' ||
		value == '(' || value == ')' ||
		value == '{' || value == '}' ||
		value == '[' || value == ']' {
		return 1, NewToken(TokenPunctuation, string(value))
	}
	return 0, nil
}

var MatchArrowOperator TokenMatcher = func(lexer *Lexer) (int, *Token) {
	if lexer.inputCode[lexer.cursor] != '-' || lexer.cursor+2 > lexer.inputLength-1 {
		return 0, nil
	}
	value := lexer.inputCode[lexer.cursor : lexer.cursor+2]
	if value == "->" {
		return 2, NewToken(TokenOperator, value)
	}
	return 0, nil
}

var MatchOperator TokenMatcher = func(lexer *Lexer) (int, *Token) {
	value := lexer.inputCode[lexer.cursor]
	if value == '+' || value == '-' || value == '*' || value == '/' {
		return 1, NewToken(TokenOperator, string(value))
	}
	return 0, nil
}
