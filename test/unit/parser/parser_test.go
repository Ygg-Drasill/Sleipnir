package parser

import (
	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"testing"
)

func TestParseNode(t *testing.T) {
	lexer := NewLexerFromString("../../testData/snippets/testFindTokens.ygl")
	tokens := lexer.FindTokens()
	p := parser.NewParser()
	res, err := p.Parse(NewScanner(tokens))

	if err != nil {
		t.Errorf("Failed to parse snippet)")
	}

	if res == nil {
		t.Errorf("Parser did not ouput result")
	}
}
