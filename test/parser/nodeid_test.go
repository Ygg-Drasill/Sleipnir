package parser

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/compiler/analysis/lexer"
	"github.com/Ygg-Drasill/Sleipnir/compiler/ast"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/parser"
	"os"
	"pgregory.net/rapid"
	"strings"
	"testing"
)

//

const alphabeticRunes = "abcdefghijklmnopqrstuvwxyz"

// TestNodeId is for testing that only valid node ids are accepted
func TestNodeId(t *testing.T) {
	formatFile, err := os.ReadFile("../data/samples/format/nodeId.ygl")
	format := string(formatFile)

	if err != nil {
		panic(err)
	}
	rapid.Check(t, func(t *rapid.T) {
		sampleNodeIdGen := rapid.SliceOfBytesMatching("[ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvxyz]+")
		sampleNodeId := string(sampleNodeIdGen.Draw(t, "Node id"))
		sample := fmt.Sprintf(format, sampleNodeId)
		l := lexer.NewLexerFromString(sample)
		tokens := l.FindTokens()
		scanner := lexer.NewScanner(tokens)
		p := parser.NewParser()
		p.Context = ast.NewParseContext()
		_, err := p.Parse(scanner)

		if len(sampleNodeId) < 1 && err == nil {
			t.Fatalf("parsed empty node id: %s", sampleNodeId)
		}

		if !strings.ContainsRune(strings.ToUpper(alphabeticRunes), rune(sampleNodeId[0])) && err == nil {
			t.Fatalf("parsed non-capitalised node id: %s", sampleNodeId)
		}
	})
}
