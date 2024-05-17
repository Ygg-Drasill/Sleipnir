package generator

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/standardTemplates"
)

func (g *Generator) write(code string, args ...interface{}) {
	g.outBuffer.WriteString(fmt.Sprintf(code, args...))
}

// GenWrapper starts the generation process from the root Program node.
func (g *Generator) GenWrapper() (*bytes.Buffer, error) {
	err := g.gen(g.syntaxTree)
	return g.outBuffer, err
}

func (g *Generator) gen(node ast.Attribute) error {
	switch node := node.(type) {
	case *ast.Program:
		return g.genProgram(node)
	case *ast.Node:
		return g.genNode(node)
	case *ast.AssignmentStatement:
		return g.genAssignmentStmt(node)
	case *ast.Statement:
		return g.genStmt(node)
	case *ast.Expression:
		return g.genExpr(node)
	case ast.Expression:
		return g.genExpr(&node)
	case int64:
		return g.genInt(node)
	}
	return nil
}

func (g *Generator) genProgram(node *ast.Program) error {
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
				return errors.New(fmt.Sprintf("template does not exist for %s\n", n.TemplateId))
			}
			for _, outId := range template.Outputs {
				nodeId := n.Id
				g.write("(global $%s_%s (mut i32) (i32.const 0))\n", nodeId, outId)
			}
		}

		g.write("(global $%s_processed (mut i32) (i32.const 0))\n", n.Id)
	}

	for _, nodes := range node.Nodes {
		if err := g.gen(&nodes); err != nil {
			return err
		}
	}

	g.genRoots(rootNodes)
	g.write(")")
	return nil
}

var UnknownStatementError = errors.New("unknown statement")

func (g *Generator) genStmt(node *ast.Statement) error {
	if ifStatement, ok := (*node).(ast.IfStatement); ok {
		return g.genIfStatement(&ifStatement)
	}
	if assStmt, ok := (*node).(ast.AssignmentStatement); ok {
		return g.gen(&assStmt)
	}
	if decStmt, ok := (*node).(ast.Declaration); ok {
		return g.genDeclaration(&decStmt)
	}
	if isExitStmt(*node) {
		g.genExitStmt()
		return nil
	}
	return UnknownStatementError
}

func (g *Generator) genAssignmentStmt(node *ast.AssignmentStatement) error {

	_, isExpression := node.Expression.(ast.Expression)
	_, isLiteral := node.Expression.(int64)

	if isExpression || isLiteral {
		if err := g.gen(node.Expression); err != nil {
			return err
		}
	}

	if identifier, ok := g.isIdentifier(&node.Expression); ok {
		g.write("%s\n", identifier.toGetInstruction())
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

	if err := g.genAssignment(identifier); err != nil {
		return err
	}

	return nil
}

func (g *Generator) genDeclaration(node *ast.Declaration) error {
	if expr, ok := node.Expression.(ast.Expression); ok {
		if err := g.gen(&expr); err != nil {
			return err
		}
	}

	g.write("local.set $%s\n", node.AssigneeId)
	return nil
}

func (g *Generator) genAssignment(identifier Identifier) error {
	if nodeIdentifier, ok := identifier.(NodeInIdentifier); ok {
		g.write("%s\n", nodeIdentifier.toSetInstruction())
	}
	if localIdentifier, ok := identifier.(LocalIdentifier); ok {
		g.write("%s\n", localIdentifier.toSetInstruction())
	}
	return nil
}

func (g *Generator) genInt(val int64) error {
	g.write("i32.const %d\n", val)
	return nil
}

func (g *Generator) genRoots(rootNodes []ast.Node) {
	g.write("(func (export \"root\")\n")
	for _, root := range rootNodes {
		g.write("call $%s\n", root.Id)
	}
	g.write(")\n")
}
