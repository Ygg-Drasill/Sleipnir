package generator

import "github.com/Ygg-Drasill/Sleipnir/pkg/ast"

func (g *Generator) genIfStatement(statement *ast.IfStatement) error {
	if err := g.genExpr(&statement.Expression); err != nil {
		return err
	}
	g.write("(if\n")

	g.write("(then\n")
	for _, stmt := range statement.BodyStatements {
		if err := g.genStmt(&stmt); err != nil {
			return err
		}
	}
	g.write(")\n")

	if len(statement.ElseStatements) > 0 {
		g.write("(else\n")
		for _, stmt := range statement.ElseStatements {
			if err := g.genStmt(&stmt); err != nil {
				return err
			}
		}
		g.write(")\n")
	}

	g.write(")\n")

	return nil
}
