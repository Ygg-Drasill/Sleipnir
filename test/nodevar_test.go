package test

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/compiler"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/parser"
	"github.com/Ygg-Drasill/Sleipnir/pkg/lexer"
	"os"
	"pgregory.net/rapid"
	"strings"
	"testing"
)

func TestNodeVar(t *testing.T) {
	formatFile, err := os.ReadFile("./samples/format/nodeVar.ygl")
	format := string(formatFile)

	if err != nil {
		panic(err)
	}
	rapid.Check(t, func(t *rapid.T) {
		sampleNodeVarGen := rapid.SliceOfBytesMatching("[a-z][A-Za-z]+")
		sampleNodeVar := string(sampleNodeVarGen.Draw(t, "Node variable"))
		sample := fmt.Sprintf(format, sampleNodeVar)
		l := lexer.NewFromString(sample)
		tokens := l.FindTokens()
		scanner := compiler.NewScanner(tokens)
		p := parser.NewParser()
		p.Context = ast.NewParseContext()
		_, err := p.Parse(scanner)

		if len(sampleNodeVar) < 1 && err == nil {
			t.Fatalf("parsed empty node id: %s", sampleNodeVar)
		}

		if !strings.ContainsRune(strings.ToUpper(alphabeticRunes), rune(sampleNodeVar[0])) && err == nil {
			t.Fatalf("parsed non-capitalised node id: %s", sampleNodeVar)
		}
	})
}
