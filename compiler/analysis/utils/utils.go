package utils

import "strings"

const alphabeticRunes = "abcdefghijklmnopqrstuvwxyz"

const numbers = "1234567890"

const operators = "&|!=+-*/%<>"

const punctuation = ";:.(){}[]"

func IsLetter(runeToCheck rune) bool {
	lowercase := strings.ContainsRune(alphabeticRunes, runeToCheck)
	uppercase := strings.ContainsRune(strings.ToUpper(alphabeticRunes), runeToCheck)
	return lowercase || uppercase
}

func IsNumber(runeToCheck rune) bool {
	return strings.ContainsRune(numbers, runeToCheck)
}

func IsOperator(runeToCheck rune) bool {
	return strings.ContainsRune(operators, runeToCheck)
}

func IsPunctuation(runeToCheck rune) bool {
	return strings.ContainsRune(punctuation, runeToCheck)
}
