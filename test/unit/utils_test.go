package unit

import (
	"github.com/Ygg-Drasill/Sleipnir/utils"
	"testing"
)

func TestValidateValidFilePath(t *testing.T) {
	correctPath := "../testData/example.ygl"

	isValid, err := utils.ValidateYglFilePath(correctPath)
	if err != nil {
		t.Fatalf("\"%s\" threw an error: %s", correctPath, err.Error())
	}
	if !isValid {
		t.Errorf("\"%s\" returned invalid, expected valid", correctPath)
	}
}

func TestValidateInvalidFilePath(t *testing.T) {
	invalidPath := "../testData/example.html"

	isValid, err := utils.ValidateYglFilePath(invalidPath)
	if err == nil {
		t.Errorf("\"%s\" did not return an error", invalidPath)
	}
	if isValid {
		t.Errorf("\"%s\" returned valid, expected invalid", invalidPath)
	}
}
