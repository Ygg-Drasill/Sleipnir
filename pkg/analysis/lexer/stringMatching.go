package lexer

import "strings"

const alphanumericRunes = "abcdefghijklmnopqrstuvwxyz"

const numbers = "1234567890"

const operators = "&|!=+-*/%<>"

const punctuation = ";:.(){}[]"

func isLetter(runeToCheck rune) bool {
	lowercase := strings.ContainsRune(alphanumericRunes, runeToCheck)
	uppercase := strings.ContainsRune(strings.ToUpper(alphanumericRunes), runeToCheck)
	return lowercase || uppercase
}

func isNumber(runeToCheck rune) bool {
	return strings.ContainsRune(numbers, runeToCheck)
}

func isOperator(runeToCheck rune) bool {
	return strings.ContainsRune(operators, runeToCheck)
}

func isPunctuation(runeToCheck rune) bool {
	return strings.ContainsRune(punctuation, runeToCheck)
}
