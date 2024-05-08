package generator

import (
	"bytes"
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
	"log/slog"
	"reflect"
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
	fmt.Println(reflect.TypeOf(node))
	switch node := node.(type) {
	case *ast.Program:
		return g.genProgram(node)
	case *ast.Node:
		return g.genNode(node)
	case *ast.Statement:
		return g.genStmt(node)
	case *ast.AssignmentStatement:
		return g.genAssStmt(node)
	case *ast.Expression:
		return g.genExpr(node)
	case ast.Expression:
		return g.genExpr(&node)
	case int64:
		g.genInt(node)
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
	g.genNodeLocals(&node.ProcStatements)
	g.genNodeGlobals(&node.OutDeclarations)

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

func (g *Generator) genNodeLocals(statements *ast.StatementList) string {
	for _, statement := range *statements {
		dec, ok := statement.(ast.Declaration)
		if !ok {
			continue
		}
		g.write("(local $%s i32)\n", dec.AssigneeId)
	}
	return ""
}

func (g *Generator) genNodeGlobals(decList *ast.DeclarationList) string {
	for _, assignment := range *decList {
		value := 0
		if v, ok := assignment.Expression.(int64); assignment.Expression != nil && ok {
			value = int(v)
		}

		g.write("(global.set $%s_%s (i32.const %d))\n", g.currentNode.Id, assignment.AssigneeId, value)
	}

	return ""
}

func (g *Generator) genStmt(node *ast.Statement) string {
	if assStmt, ok := (*node).(ast.AssignmentStatement); ok {
		g.gen(&assStmt)

	}
	if decStmt, ok := (*node).(ast.Declaration); ok {
		g.genDeclaration(&decStmt)
	}
	return ""
}

func (g *Generator) genAssStmt(node *ast.AssignmentStatement) string {

	if expr, ok := node.Expression.(ast.Expression); ok {
		g.gen(&expr)
	}

	if identifier, ok := g.isIdentifier(&node.Expression); ok {
		g.write("%s\n", identifier.getOperation())
	}

	var identifier Identifier
	if nodeVar, ok := (*node).Identifier.(ast.NodeVar); ok {
		identifier = NodeInIdentifier{
			NodeVar:        nodeVar,
			sourceJunction: nil,
			nodeId:         g.currentNode.Id,
		}
	}

	if localVar, ok := (*node).Identifier.(ast.LocalVar); ok {
		identifier = LocalIdentifier{localVar}
	}

	g.genAssignment(identifier)

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

func (g *Generator) genDeclaration(node *ast.Declaration) string {
	if expr, ok := node.Expression.(ast.Expression); ok {
		g.gen(&expr)
	}

	g.write("local.set $%s\n", node.AssigneeId)
	return ""
}

func (g *Generator) genAssignment(identifier Identifier) string {
	if nodeIdentifier, ok := identifier.(NodeInIdentifier); ok {
		g.write("%s\n", nodeIdentifier.setOperation())
	}
	if localIdentifier, ok := identifier.(LocalIdentifier); ok {
		g.write("%s\n", localIdentifier.setOperation())
	}
	return ""
}

func (g *Generator) genInt(val int64) string {
	g.write("i32.const %d\n", val)
	return ""
}
