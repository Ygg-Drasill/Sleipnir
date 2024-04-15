package ast

import (
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
)

func NewConnectionList(connection Attribute) (ConnectionList, error) {
	return ConnectionList{connection.(Connection)}, nil
}

func AppendConnection(connectionList, connection Attribute) (ConnectionList, error) {
	return append(connectionList.(ConnectionList), connection.(Connection)), nil
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
