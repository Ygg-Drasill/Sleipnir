package generator

import (
	"bytes"
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

type Generator struct {
	outBuffer   *bytes.Buffer
	syntaxTree  *ast.Program
	context     *ast.ParseContext
	currentNode *ast.Node
	outNodeVars map[string]ast.Junction
}

func New(tree *ast.Program, ctx *ast.ParseContext) *Generator {
	newGen := &Generator{
		outBuffer:   new(bytes.Buffer),
		syntaxTree:  tree,
		context:     ctx,
		outNodeVars: make(map[string]ast.Junction),
	}
	newGen.memoOutVariables()
	return newGen
}

func (g *Generator) memoOutVariables() {
	for _, v := range g.syntaxTree.Connections {
		g.outNodeVars[junctionToKey(v.InJunction)] = v.OutJunction
	}
}

func junctionToKey(junction ast.Junction) string {
	return junctionKey(junction.NodeId, junction.VarId)
}

func junctionKey(nodeId, varId string) string {
	return fmt.Sprintf("%s-%s", nodeId, varId)
}
