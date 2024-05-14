package generator

import "github.com/Ygg-Drasill/Sleipnir/pkg/ast"

func (g *Generator) genIfStatement(statement *ast.IfStatement) string {
	g.genExpr(&statement.Expression)
	g.write("(if")

	g.write("(then")
	for _, stmt := range statement.ElseStatements {
		g.genStmt(&stmt)
	}
	g.write(")")

	if len(statement.ElseStatements) > 0 {
		g.write("(else")
		for _, stmt := range statement.ElseStatements {
			g.genStmt(&stmt)
		}
		g.write(")")
	}

	g.write(")")

	return ""
}
