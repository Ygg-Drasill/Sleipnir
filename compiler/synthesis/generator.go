package synthesis

import (
	"bytes"
	"fmt"

	"github.com/Ygg-Drasill/Sleipnir/compiler/ast"
)

func write(b *bytes.Buffer, code string, args ...interface{}) {
	b.WriteString(fmt.Sprintf(code, args...))
}

func GenWrapper(p *ast.Program) bytes.Buffer {
	var b bytes.Buffer
	gen(p, &b)
	return b
}

func gen(node ast.Attribute, b *bytes.Buffer) string {
	switch node := node.(type) {
	case *ast.Program:
		return genProgram(node, b)
	case *ast.Node:
		return genNode(node, b)
	case *ast.DeclarationList:
		return genDecLst(node, b)
	}
	return ""
}

func genProgram(node *ast.Program, b *bytes.Buffer) string {

	write(b, "(module\n")

	for _, nodes := range node.Nodes {
		nodePtr := &nodes
		gen(nodePtr, b)

		write(b, "\n)\n")
	}
	write(b, ")")
	return ""
}

func genNode(node *ast.Node, b *bytes.Buffer) string {
	value := node.Id
	write(b, "(func $%s", value)

	for _, inDec := range node.InDeclarations {
		inAss := inDec.AssigneeId

		write(b, " (param $%s i32)", inAss)
	}
	for _, outDec := range node.OutDeclarations {
		outAss := outDec.AssigneeId

		write(b, " (local $%s i32)", outAss)
	}
	return ""
}

func genConn(node *ast.Junction, b *bytes.Buffer) string {
	return ""
}

func genDecLst(node *ast.DeclarationList, b *bytes.Buffer) string {

	value := node
	write(b, " (param $%s i64)", value)
	return ""
}
