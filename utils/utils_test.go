package utils

import (
	"testing"
)

func TestValidateValidFilePath(t *testing.T) {
	correctPath := "./example.ygl"

	isValid, err := ValidateYglFilePath(correctPath)
	if err != nil {
		t.Fatalf("\"%s\" threw an error: %s", correctPath, err.Error())
	}
	if !isValid {
		t.Errorf("\"%s\" returned invalid, expected valid", correctPath)
	}
}

func TestValidateInvalidFilePath(t *testing.T) {
	invalidPath := "./example.html"

	isValid, err := ValidateYglFilePath(invalidPath)
	if err == nil {
		t.Errorf("\"%s\" did not return an error", invalidPath)
	}
	if isValid {
		t.Errorf("\"%s\" returned valid, expected invalid", invalidPath)
	}
}
