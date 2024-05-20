package generator

import (
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
)

func (g *Generator) genExitStmt() {
	g.write("return\n")
}

func isExitStmt(node ast.Statement) bool {
	tok, ok := node.(*token.Token)
	return ok && string(tok.Lit) == "exit"
}
