package ast

import "github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"

func NewProgram(nodes, connections Attribute) (Program, error) {
	return Program{
		Nodes:       nodes.(NodeList),
		Connections: connections.(ConnectionList),
	}, nil
}

func NewNodeList(node Attribute) (NodeList, error) {
	return NodeList{node.(Node)}, nil
}

func AppendNode(nodeList, node Attribute) (NodeList, error) {
	return append(nodeList.(NodeList), node.(Node)), nil
}

func NewConnectionList(connection Attribute) (ConnectionList, error) {
	return ConnectionList{connection.(Connection)}, nil
}

func AppendConnection(connectionList, connection Attribute) (ConnectionList, error) {
	return append(connectionList.(ConnectionList), connection.(Connection)), nil
}

func NewNode(node, in, out, process Attribute) (Node, error) {
	return Node{id: string(node.(*token.Token).Lit)}, nil
}

func NewConnection(out, in Attribute) (Connection, error) {
	return Connection{outId: out.(Junction), inId: in.(Junction)}, nil
}

func NewJunction(nodeId, varId Attribute) (Junction, error) {
	varIdToken, ok := varId.(*token.Token)
	var varIdStr string

	if ok {
		varIdStr = string(varIdToken.Lit)
	} else {
		varIdStr = ""
	}
	return Junction{
		nodeId: string(nodeId.(*token.Token).Lit),
		varId:  varIdStr,
	}, nil
}

func NewDeclaration(varId, expression Attribute) (Declaration, error) {
	return Declaration{
		Assignee:   varId.(string),
		Expression: expression,
	}, nil
}

func NewDeclarationList(varId, expression Attribute) (DeclarationList, error) {
	return DeclarationList{Declaration{
		Assignee:   varId.(string),
		Expression: expression,
	}}, nil
}
