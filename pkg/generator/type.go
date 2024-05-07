package generator

import (
	"bytes"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

type Generator struct {
	outBuffer   *bytes.Buffer
	syntaxTree  *ast.Program
	context     *ast.ParseContext
	currentNode *ast.Node
}

func New(tree *ast.Program, ctx *ast.ParseContext) *Generator {
	return &Generator{
		outBuffer:  new(bytes.Buffer),
		syntaxTree: tree,
		context:    ctx,
	}
}
