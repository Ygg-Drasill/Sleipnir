package synthesis

import (
	"bytes"
	"fmt"

	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

type Generator struct {
	outBuffer   *bytes.Buffer
	connections *ast.ConnectionList
}

func (g *Generator) write(code string, args ...interface{}) {
	g.outBuffer.WriteString(fmt.Sprintf(code, args...))
}

// GenWrapper starts the generation process from the root Program node.
func GenWrapper(p *ast.Program) *bytes.Buffer {
	b := new(bytes.Buffer)
	generator := &Generator{outBuffer: b, connections: &p.Connections}
	generator.gen(p)
	return b
}

func (g *Generator) gen(node ast.Attribute) string {
	switch node := node.(type) {
	case *ast.Program:
		return g.genProgram(node)
	case *ast.Node:
		return g.genNode(node)
	case *ast.DeclarationList:
		return g.genDecLst(node)
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
	for _, conn := range *g.connections {
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

	g.gen(node.ProcStatements)

	clear(connectionsMemo)
	connectionsMemo = make(map[string]bool)
	for _, conn := range *g.connections {
		if conn.OutId.NodeId == node.Id && !connectionsMemo[conn.InId.NodeId] {
			g.write("call $%s\n", conn.InId.NodeId)
			connectionsMemo[conn.InId.NodeId] = true
		}
	}
	g.write(")\n")
	return ""
}

func (g *Generator) genDecLst(node *ast.DeclarationList) string {

	value := node
	g.write(" (param $%s i64)", value)
	return ""
}
