package generator

import (
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/standardTemplates"
)

func (g *Generator) genNode(node *ast.Node) error {
	userTemplate := g.context.Templates[node.TemplateId]
	standardTemplate := standardTemplates.StandardTemplates[node.TemplateId]
	outDeclarations := node.OutDeclarations
	process := node.ProcStatements
	var err error

	//don't generate template definition
	if g.context.Templates[node.Id] != nil {
		return nil
	}

	if userTemplate != nil {
		outDeclarations = userTemplate.OutDeclarations
		process = userTemplate.ProcStatements
	}

	g.currentNode = node
	g.write("(func $%s\n", node.Id)
	err = g.genNodeLocals(&process)
	err = g.genNodeGlobals(&outDeclarations)

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
		g.write("global.get $%s_processed (if (then nop) (else return))\n", conn.OutJunction.NodeId)
		connectionsMemo[conn.OutJunction.NodeId] = true
	}
	g.write("i32.const 1\n")
	g.write("(global.set $%s_processed)\n", g.currentNode.Id)

	//if template is standard
	if standardTemplate != nil {
		g.write("%s\n", standardTemplate.FormatBody(*standardTemplate, node.Id, g.outNodeVars))
	}

	for _, stmt := range process {
		if err := g.gen(&stmt); err != nil {
			return err
		}
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
	return err
}

func (g *Generator) genNodeLocals(statements *ast.StatementList) error {
	for _, statement := range *statements {
		dec, ok := statement.(ast.Declaration)
		if !ok {
			continue
		}
		g.write("(local $%s i32)\n", dec.AssigneeId)
	}
	return nil

}

func (g *Generator) genNodeGlobals(decList *ast.DeclarationList) error {
	for _, assignment := range *decList {
		value := 0
		if v, ok := assignment.Expression.(int64); assignment.Expression != nil && ok {
			value = int(v)
		}

		g.write("(global.set $%s_%s (i32.const %d))\n", g.currentNode.Id, assignment.AssigneeId, value)
	}

	return nil
}

func (g *Generator) isRoot(node *ast.Node) bool {
	isRoot := true
	for _, conn := range g.syntaxTree.Connections {
		if conn.InJunction.NodeId == node.Id {
			isRoot = false
		}
	}
	return isRoot && g.context.Templates[node.Id] == nil
}
