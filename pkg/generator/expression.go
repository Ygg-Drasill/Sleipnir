package generator

import (
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
	"log"
	"log/slog"
	"strconv"
)

func (g *Generator) genExpr(node *ast.Expression) string {
	g.genExprOperand(&node.FirstOperand)
	g.genExprOperand(&node.SecondOperand)

	exprOp := string(node.Operator.(*token.Token).Lit)
	secondOp, err := strconv.Atoi(string(node.Operator.(*token.Token).Lit))
	if err != nil {
		log.Fatalf("Error converting secondOp to int: %v", err.Error())
	}
	switch exprOp {
	case "+":
		g.write("i32.add\n")
		break
	case "-":
		g.write("i32.sub\n")
		break
	case "*":
		g.write("i32.mul\n")
		break
	case "/":
		// TODO: move to error handling to catch error before code generation
		if secondOp == 0 {
			log.Fatalf("Division by zero is not allowed")
		}
		g.write("i32.div_s\n")
		break
	case "%":
		g.write("i32.rem_s\n")
		break
	default:
		slog.Error("Failed to generate unknown operator")
	}
	return ""
}

func (g *Generator) genExprOperand(node *ast.Attribute) {
	if operand, ok := (*node).(ast.Expression); ok {
		g.genExpr(&operand)
	} else {
		g.genValue(node)
	}
}

func (g *Generator) genValue(node *ast.Attribute) string {
	if identifier, ok := g.isIdentifier(node); ok {
		g.genVarUsage(&identifier)
	}
	if val, ok := (*node).(int64); ok {
		g.genInt(val)
	}
	return ""
}

func (g *Generator) genVarUsage(node *Identifier) string {
	identifier, ok := (*node).(Identifier)
	if !ok {
		slog.Error("Failed to generate identifier")
	}

	label := identifier.getOperation()
	g.write("%s\n", label)
	return ""
}
