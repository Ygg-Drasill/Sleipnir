package generator

import (
	"bytes"
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
	"log/slog"
)

func (g *Generator) write(code string, args ...interface{}) {
	g.outBuffer.WriteString(fmt.Sprintf(code, args...))
}

// GenWrapper starts the generation process from the root Program node.
func (g *Generator) GenWrapper() *bytes.Buffer {
	g.gen(g.syntaxTree)
	return g.outBuffer
}

func (g *Generator) gen(node ast.Attribute) string {
	switch node := node.(type) {
	case *ast.Program:
		return g.genProgram(node)
	case *ast.Node:
		return g.genNode(node)
	case *ast.Declaration:
		return g.genDecLst(node)
	case *ast.Statement:
		return g.genStmt(node)
	case *ast.AssignmentStatement:
		return g.genAssStmt(node)
	case *ast.Expression:
		return g.genExpr(node)
	}
	return ""
}

func (g *Generator) genProgram(node *ast.Program) string {

	g.write("(module\n")

	for _, n := range node.Nodes {
		for _, outDec := range n.OutDeclarations {
			outId := n.Id
			outAss := outDec.AssigneeId
			g.write("(global $%s_%s (mut i32) (i32.const 0))\n", outId, outAss)
		}

		g.write("(global $%s_processed (mut i32) (i32.const 0))\n", n.Id)
	}

	for _, nodes := range node.Nodes {
		g.gen(&nodes)
	}
	g.write(")")
	return ""
}

func (g *Generator) genNode(node *ast.Node) string {
	g.currentNode = node
	g.write("(func $%s\n", node.Id)

	inputs := make([]ast.Connection, 0)
	for _, conn := range g.syntaxTree.Connections {
		if conn.InJunction.NodeId != node.Id {
			continue
		}
		inputs = append(inputs, conn)
	}

	connectionsMemo := make(map[string]bool)
	for _, conn := range inputs {
		if connectionsMemo[conn.InJunction.NodeId] {
			break
		}
		g.write("global.get $%s_processed\n", conn.OutJunction.NodeId)
		g.write("(if (then nop) (else return))\n")
		connectionsMemo[conn.InJunction.NodeId] = true
	}

	for _, stmt := range node.ProcStatements {
		g.gen(&stmt)
	}
	for _, decStmt := range node.OutDeclarations {
		g.gen(&decStmt)
	}

	clear(connectionsMemo)
	connectionsMemo = make(map[string]bool)
	for _, conn := range g.syntaxTree.Connections {
		if conn.OutJunction.NodeId == node.Id && !connectionsMemo[conn.InJunction.NodeId] {
			g.write("call $%s\n", conn.InJunction.NodeId)
			connectionsMemo[conn.InJunction.NodeId] = true
		}
	}
	g.write(")\n")
	return ""
}

func (g *Generator) genStmt(node *ast.Statement) string {
	if assStmt, ok := (*node).(ast.AssignmentStatement); ok {
		g.gen(&assStmt)

	}
	return ""
}

func (g *Generator) genAssStmt(node *ast.AssignmentStatement) string {

	expr := node.Expression.(ast.Expression)
	g.gen(&expr)

	return ""
}

func (g *Generator) genExpr(node *ast.Expression) string {
	g.genExprOperand(&node.FirstOperand)
	g.genExprOperand(&node.SecondOperand)

	exprOp := string(node.Operator.(*token.Token).Lit)

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

func (g *Generator) genDecLst(node *ast.Declaration) string {
	if node.Expression != nil {
		g.write("i32.const %d\n", node.Expression)
	}

	g.write("global.set $%s_%s\n", g.currentNode.Id, node.AssigneeId)
	return ""
}
