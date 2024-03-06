package utils

import "strings"

const alphanumericRunes = "abcdefghijklmnopqrstuvwxyz"

func IsLetter(runeToCheck rune) bool {
	lowercase := strings.ContainsRune(alphanumericRunes, runeToCheck)
	uppercase := strings.ContainsRune(strings.ToUpper(alphanumericRunes), runeToCheck)
	return lowercase || uppercase
}

const numbers = "1234567890"

func IsNumber(runeToCheck rune) bool {
	return strings.ContainsRune(numbers, runeToCheck)
}
