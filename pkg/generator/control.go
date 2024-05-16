package generator

import "github.com/Ygg-Drasill/Sleipnir/pkg/ast"

func (g *Generator) genIfStatement(statement *ast.IfStatement) string {
	g.genExpr(&statement.Expression)
	g.write("(if\n")

	g.write("(then\n")
	for _, stmt := range statement.BodyStatements {
		g.genStmt(&stmt)
	}
	g.write(")\n")

	if len(statement.ElseStatements) > 0 {
		g.write("(else\n")
		for _, stmt := range statement.ElseStatements {
			g.genStmt(&stmt)
		}
		g.write(")\n")
	}

	g.write(")\n")

	return ""
}
