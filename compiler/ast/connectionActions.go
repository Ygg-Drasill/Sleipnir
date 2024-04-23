package ast

import (
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
)

func NewConnectionList() (ConnectionList, error) {
	return ConnectionList{}, nil
}

func AppendConnection(connectionList, connection Attribute) (ConnectionList, error) {
	return append(connectionList.(ConnectionList), connection.(Connection)), nil
}

func NewConnection(out, in Attribute) (Connection, error) {
	return Connection{OutId: out.(Junction), InId: in.(Junction)}, nil
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
		NodeId: string(nodeId.(*token.Token).Lit),
		VarId:  varIdStr,
	}, nil
}
