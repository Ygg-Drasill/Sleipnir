package generator

import (
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/standardTemplates"
)

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
		connectionsMemo[conn.OutJunction.NodeId] = true
	}
	g.write("i32.const 1\n")
	g.write("(global.set $%s_processed)\n", g.currentNode.Id)

	if len(node.TemplateId) > 0 {
		template := standardTemplates.StandardTemplates[node.TemplateId]
		if template != nil {
			g.write("%s\n", template.FormatBody(*template, node.Id, g.outNodeVars))
		}
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

func (g *Generator) isRoot(node *ast.Node) bool {
	isRoot := true
	for _, conn := range g.syntaxTree.Connections {
		if conn.InJunction.NodeId == node.Id {
			isRoot = false
		}
	}
	return isRoot
}
