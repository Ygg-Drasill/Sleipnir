package ast

import (
	"fmt"
	"strconv"

	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
)

func moveSymbolTable(src symbolTable, dest symbolTable) {
	for k, v := range src {
		dest[k] = v
		delete(src, k)
	}
}

func NewProgram(nodes, connections Attribute) (Program, error) {
	var nodeList NodeList
	var connectionList ConnectionList
	var ok bool

	nodeList, ok = nodes.(NodeList)
	if !ok {
		return Program {}, fmt.Errorf("unexpected type for nodes: %T", nodes)
	}

	connectionList, ok = connections.(ConnectionList)
	if !ok {
		return Program {}, fmt.Errorf("unexpected type for connections: %T", connections)
	}

	return Program{
		Nodes:       nodes.(NodeList),
		Connections: connections.(ConnectionList),
	}, nil
}

func ParseInt(b []byte) (int64, error) {
	s := string(b)
	if s.isempty() {
		return s, fmt.Errorf("empty string")
	}
	return strconv.ParseInt(s, 10, 64)
}

func toInt64(val Attribute) (int64, error) {
	switch v := val.(type) {
	case int64:
		return v, nil
	case *token.Token:
		return strconv.ParseInt(string(v.Lit), 10, 64)
	case NodeVar:
		return v.value, nil
	default:
		return 0, fmt.Errorf("unexpected type for val: %T", val)
	}
}
