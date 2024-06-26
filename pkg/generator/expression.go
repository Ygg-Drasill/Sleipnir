package generator

import (
	"errors"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
	"log"
)

func (g *Generator) genExpr(node *ast.Expression) error {
	var err error
	err = g.genExprOperand(&node.FirstOperand)
	if err != nil {
		return err
	}
	err = g.genExprOperand(&node.SecondOperand)
	if err != nil {
		return err
	}
	opToken := node.Operator.(*token.Token)
	tokenLit := string(opToken.Lit)
	if tokenLit == "/" || tokenLit == "%" {
		err := handleZeroDivision(node)

		if err != nil {
			return errors.Join(errors.New(opToken.Pos.String()), err)
		}
	}

	g.genInstruction(opToken)
	return nil
}

func (g *Generator) genExprOperand(node *ast.Attribute) error {
	var err error
	if operand, ok := (*node).(ast.Expression); ok {
		err = g.genExpr(&operand)
	} else {
		err = g.genValue(node)
	}

	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) genValue(node *ast.Attribute) error {
	identifier, err := g.isIdentifier(node)
	if err != nil {
		return err
	}
	if identifier != nil {
		g.genVarUsage(&identifier)
	} else {
		if val, ok := (*node).(int64); ok {
			err = g.genInt(val)
		}
	}
	return err
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
		g.write("i64.add\n")
		break
	case "-":
		g.write("i64.sub\n")
		break
	case "*":
		g.write("i64.mul\n")
		break
	case "/":
		g.write("i64.div_s\n")
		break
	case "%":
		g.write("i64.rem_s\n")
		break
	case ">":
		g.write("i64.gt_s\n")
		break
	case ">=":
		g.write("i64.ge_s\n")
		break
	case "<":
		g.write("i64.lt_s\n")
		break
	case "<=":
		g.write("i64.le_s\n")
		break
	case "==":
		g.write("i64.eq_s\n")
		break
	default:
		log.Fatalf("%s: Failed to generate unknown operator %s", opToken.Pos.String(), exprOp)
	}
}

var ZeroDivisionError = errors.New("zero division is not allowed")
var ZeroModuloError = errors.New("zero modulo is not allowed")

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
	if string(opToken.Lit) == "%" && secondOperand == 0 {
		return ZeroModuloError
	}
	return nil
}
