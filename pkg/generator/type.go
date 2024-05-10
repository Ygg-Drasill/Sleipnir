package generator

import (
	"bytes"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/utils"
)

type Generator struct {
	outBuffer   *bytes.Buffer
	syntaxTree  *ast.Program
	context     *ast.ParseContext
	currentNode *ast.Node
	outNodeVars map[string]*ast.Junction
}

func New(tree *ast.Program, ctx *ast.ParseContext) *Generator {
	newGen := &Generator{
		outBuffer:   new(bytes.Buffer),
		syntaxTree:  tree,
		context:     ctx,
		outNodeVars: make(map[string]*ast.Junction),
	}
	newGen.memoOutVariables()
	return newGen
}

func (g *Generator) memoOutVariables() {
	for _, v := range g.syntaxTree.Connections {
		g.outNodeVars[utils.JunctionToKey(v.InJunction)] = &v.OutJunction
	}
}
