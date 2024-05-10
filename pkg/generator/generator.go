package generator

import (
	"bytes"
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/standardTemplates"
	"log"
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
	rootNodes := make([]ast.Node, 0)

	g.write("(module\n")
	g.write("%s\n", imports)

	for _, n := range node.Nodes {
		if g.isRoot(&n) {
			rootNodes = append(rootNodes, n)
		}

		for _, outDec := range n.OutDeclarations {
			nodeId := n.Id
			varId := outDec.AssigneeId
			g.write("(global $%s_%s (mut i32) (i32.const 0))\n", nodeId, varId)
		}

		if len(n.TemplateId) > 0 {
			template := standardTemplates.StandardTemplates[n.TemplateId]
			if template == nil {
				log.Fatalf("template does not exist for %s\n", n.TemplateId)
			}
			for _, outId := range template.Outputs {
				nodeId := n.Id
				g.write("(global $%s_%s (mut i32) (i32.const 0))\n", nodeId, outId)
			}
		}

		g.write("(global $%s_processed (mut i32) (i32.const 0))\n", n.Id)
	}

	for _, nodes := range node.Nodes {
		g.gen(&nodes)
	}

	g.genRoots(rootNodes)
	g.write(")")
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

func (g *Generator) genRoots(rootNodes []ast.Node) {
	g.write("(func (export \"root\")\n")
	for _, root := range rootNodes {
		g.write("call $%s\n", root.Id)
	}
	g.write(")\n")
}
