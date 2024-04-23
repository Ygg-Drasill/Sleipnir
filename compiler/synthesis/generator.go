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
	case *ast.Connection:
		return genConn(node, b)
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

		for _, conn := range node.Connections {
			connPtr := &conn
			gen(connPtr, b)
		}

		write(b, "\n)\n")
	}
	write(b, ")")
	return ""
}

func genNode(node *ast.Node, b *bytes.Buffer) string {
	value := node.Id
	write(b, "(func $%s", value)

	for _, inDec := range node.InDeclarations {
		inDecPtr := &inDec
		gen(inDecPtr, b)
	}
	for _, outDec := range node.OutDeclarations {
		outDecPtr := &outDec
		gen(outDecPtr, b)
	}
	return ""
}

func genConn(node *ast.Connection, b *bytes.Buffer) string {
	//var inParams, outLocals []string

	// Collect InId connections (parameters)
	//for _, conn := range node.InId.VarId {
	//	inParams = append(inParams, conn)
	//}
	//
	//// Collect OutId connections (locals)
	//for _, conn := range node.OutId {
	//	outLocals = append(outLocals, conn.VarId)
	//}
	//
	//// Write parameters for InId
	//for _, param := range inParams {
	//	write(b, " (param $%s i64)", param)
	//}
	//
	//// Write locals for OutId
	//for _, local := range outLocals {
	//	write(b, " (local $%s i64)", local)
	//}
	//
	return ""
}

func genDecLst(node *ast.DeclarationList, b *bytes.Buffer) string {

	value := node
	write(b, " (param $%s i64)", value)
	return ""
}
