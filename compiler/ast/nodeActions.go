package ast

import (
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
)

func AppendNode(nodeList, node Attribute) (NodeList, error) {
	return append(nodeList.(NodeList), node.(Node)), nil
}

func NewNodeList(node Attribute) (NodeList, error) {
	return NodeList{node.(Node)}, nil
}

func NewNode(context, node, in, out, process Attribute) (Node, error) {
	var inDeclarations DeclarationList
	var outDeclarations DeclarationList
	var processStatements StatementList

	if process != nil {
		processStatements = process.(StatementList)
	}

	nodeId := string(node.(*token.Token).Lit)
	ctx := context.(ParseContext)
	ctx.BabushkaPopScopeNode(nodeId)

	return Node{
		Id:              nodeId,
		InDeclarations:  inDeclarations,
		OutDeclarations: outDeclarations,
		ProcStatements:  processStatements,
	}, nil
}

func NewScopeIn(context, declarationList Attribute) (Attribute, error) {
	ctx := context.(ParseContext)
	ctx.BabushkaPopScopeIn()
	return nil, nil
}

func NewScopeOut(context Attribute) (Attribute, error) {
	ctx := context.(ParseContext)
	ctx.BabushkaPopScopeOut()
	return nil, nil
}

func NewScopeProc(context Attribute) (Attribute, error) {
	ctx := context.(ParseContext)
	ctx.BabushkaPopScopeProc()
	return nil, nil
}
