package generator

import (
	"errors"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
	"log"
)

func (g *Generator) genExpr(node *ast.Expression) string {
	g.genExprOperand(&node.FirstOperand)
	g.genExprOperand(&node.SecondOperand)
	opToken := node.Operator.(*token.Token)
	if string(opToken.Lit) == "/" {
		err := handleZeroDivision(node)

		if err != nil {
			log.Fatal(errors.Join(errors.New(opToken.Pos.String()), err))
		}
	}

	g.genInstruction(opToken)
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
		log.Fatalf("Failed to generate identifier")
	}

	label := identifier.toGetInstruction()
	g.write("%s\n", label)
	return ""
}

func (g *Generator) genInstruction(opToken *token.Token) {
	exprOp := string(opToken.Lit)
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
		g.write("i32.div_s\n")
		break
	case "%":
		g.write("i32.rem_s\n")
		break
	case ">":
		g.write("i32.gt_s\n")
		break
	case ">=":
		g.write("i32.ge_s\n")
		break
	case "<":
		g.write("i32.lt_s\n")
		break
	case "<=":
		g.write("i32.le_s\n")
		break
	case "==":
		g.write("i32.eq_s\n")
		break
	default:
		log.Fatalf("%s: Failed to generate unknown operator %s", opToken.Pos.String(), exprOp)
	}
}

var ZeroDivisionError = errors.New("zero division is not allowed")

func handleZeroDivision(expr *ast.Expression) error {
	opToken, isToken := expr.Operator.(*token.Token)
	if !isToken {
		return nil
	}
	secondOperand, isInt := (*expr).SecondOperand.(int64)
	if !isInt {
		return nil
	}
	if string(opToken.Lit) == "/" && secondOperand == 0 {
		return ZeroDivisionError
	}
	return nil
}
