package generator

import (
	"bytes"
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
	"log/slog"

	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
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
	case *ast.DeclarationList:
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
		nodePtr := &nodes
		g.gen(nodePtr)
	}
	g.write(")")
	return ""
}

func (g *Generator) genNode(node *ast.Node) string {
	g.write("(func $%s\n", node.Id)

	inputs := make([]ast.Connection, 0)
	for _, conn := range g.syntaxTree.Connections {
		if conn.InId.NodeId != node.Id {
			continue
		}
		inputs = append(inputs, conn)
	}

	connectionsMemo := make(map[string]bool)
	for _, conn := range inputs {
		if connectionsMemo[conn.InId.NodeId] {
			break
		}
		g.write("global.get $%s_processed\n", conn.OutId.NodeId)
		g.write("(if (then nop) (else return))\n")
		connectionsMemo[conn.InId.NodeId] = true
	}

	for _, stmt := range node.ProcStatements {
		stmtPtr := &stmt
		g.gen(stmtPtr)
	}

	clear(connectionsMemo)
	connectionsMemo = make(map[string]bool)
	for _, conn := range g.syntaxTree.Connections {
		if conn.OutId.NodeId == node.Id && !connectionsMemo[conn.InId.NodeId] {
			g.write("call $%s\n", conn.InId.NodeId)
			connectionsMemo[conn.InId.NodeId] = true
		}
	}
	g.write(")\n")
	return ""
}

func (g *Generator) genStmt(node *ast.Statement) string {
	if assStmt, ok := (*node).(ast.AssignmentStatement); ok {
		assStmtPtr := &assStmt
		g.gen(assStmtPtr)

	}
	return ""
}

func (g *Generator) genAssStmt(node *ast.AssignmentStatement) string {

	expr := node.Expression.(ast.Expression)
	exprPtr := &expr
	g.gen(exprPtr)

	return ""
}

func (g *Generator) genExpr(node *ast.Expression) string {
	if exprFirstOp, ok := node.FirstOperand.(ast.Expression); ok {
		exprFirstOpPtr := &exprFirstOp
		g.genExpr(exprFirstOpPtr)
	} else if node.FirstOperand != nil {
		if attr, ok := node.FirstOperand.(ast.Attribute); ok {
			g.genFactor(&attr)
		} else {
			slog.Error("Expected an attribute")
		}
	}

	if exprSecondOp, ok := node.SecondOperand.(ast.Expression); ok {
		exprSecondOpPtr := &exprSecondOp
		g.genExpr(exprSecondOpPtr)
	} else if node.SecondOperand != nil {
		if attr, ok := node.SecondOperand.(ast.Attribute); ok {
			g.genFactor(&attr)
		} else {
			slog.Error("Expected an attribute")
		}
	}

	exprOp := string(node.Operator.(*token.Token).Lit)

	if exprOp == "+" {
		g.write("i32.add\n")
	} else if exprOp == "-" {
		g.write("i32.sub\n")
	} else if exprOp == "*" {
		g.write("i32.mul\n")
	} else if exprOp == "/" {
		g.write("i32.div_s\n")
	} else {
		slog.Error("Unknown operator")
	}

	return ""
}

func (g *Generator) genFactor(node *ast.Attribute) string {
	if factor, ok := (*node).(ast.NodeVar); ok {
		g.write("%s\n", factor.Id)
	}
	return ""
}

func (g *Generator) genDecLst(node *ast.DeclarationList) string {

	value := node
	g.write(" (param $%s i64)", value)
	return ""
}
