package synthesis

import (
	"bytes"
	"fmt"

	"github.com/Ygg-Drasill/Sleipnir/compiler/ast"
)

type Generator struct {
	outBuffer   *bytes.Buffer
	connections *ast.ConnectionList
}

func (g *Generator) write(code string, args ...interface{}) {
	g.outBuffer.WriteString(fmt.Sprintf(code, args...))
}

// GenWrapper starts the generation process from the root Program node.
func GenWrapper(p *ast.Program) bytes.Buffer {
	var b bytes.Buffer
	generator := &Generator{outBuffer: &b, connections: &p.Connections}
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

	nodeHasRun := make(map[string]bool)
	for _, n := range node.Nodes {
		outId := n.Id
		if !nodeHasRun[outId] {
			g.write("(global $%s_Has_Run (mut i32) (i32.const 0))\n", outId)
			nodeHasRun[outId] = true
		}
	}

	for _, n := range node.Nodes {
		for _, outDec := range n.OutDeclarations {
			outId := n.Id
			outAss := outDec.AssigneeId
			g.write("(global $%s.%s (mut i32) (i32.const 0))\n", outId, outAss)
		}
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

	g.gen(node.ProcStatements)

	connectedNodes := make(map[string]bool)
	for _, conn := range *g.connections {
		if conn.InId.NodeId == node.Id && !connectedNodes[conn.InId.NodeId] {
			g.write("call $%s\n", conn.OutId.NodeId)
			connectedNodes[conn.InId.NodeId] = true
		}
	}

	for _, conn := range inputs {
		g.write("global.get $%s.%s ;; Connection %s -> %s\n", conn.OutId.NodeId, conn.OutId.VarId, conn.OutId.NodeId, conn.InId.NodeId)
	}

	// TODO: Make return/drop if node has no connection
	g.write(";; might need drop/return\n")
	g.write(")\n")
	return ""
}

func (g *Generator) genDecLst(node *ast.DeclarationList) string {

	value := node
	g.write(" (param $%s i64)", value)
	return ""
}
