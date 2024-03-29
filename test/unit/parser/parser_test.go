package parser

import (
	. "github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"testing"
)

func TestParseNode(t *testing.T) {>
	lexer := NewLexerFromString("../../testData/snippets/testFindTokens.ygl")
	tokens := lexer.FindTokens()
	p := parser.NewParser()
	p.Parse(NewScanner(tokens))
}