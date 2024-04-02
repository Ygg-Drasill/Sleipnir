package ast

import "github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"

func NewProgram(nodes, connections interface{}) (Program, error) {
	return Program{
		Nodes:       nodes.(NodeList),
		Connections: connections.(ConnectionList),
	}, nil
}

func NewNodeList(node interface{}) (NodeList, error) {
	return NodeList{node.(Node)}, nil
}

func AppendNode(nodeList, node interface{}) (NodeList, error) {
	return append(nodeList.(NodeList), node.(Node)), nil
}

func NewConnectionList(connection interface{}) (ConnectionList, error) {
	return ConnectionList{connection.(Connection)}, nil
}

func AppendConnection(connectionList, connection interface{}) (ConnectionList, error) {
	return append(connectionList.(ConnectionList), connection.(Connection)), nil
}

func NewNode(node, in, out, process interface{}) (Node, error) {
	return Node{id: string(node.(*token.Token).Lit)}, nil
}

func NewConnection(out, in interface{}) (Connection, error) {
	return Connection{outId: out.(Junction), inId: in.(Junction)}, nil
}

func NewJunction(nodeId, varId interface{}) (Junction, error) {
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
