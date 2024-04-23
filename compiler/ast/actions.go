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
	return Program{
		Nodes:       nodes.(NodeList),
		Connections: connections.(ConnectionList),
	}, nil
}

func ParseInt(b []byte) (int64, error) {
	s := string(b)
	return strconv.ParseInt(s, 10, 64)
}

func toInt64(val Attribute) (int64, error) {
	switch v := val.(type) {
	case int64:
		return v, nil
	case *token.Token:
		return strconv.ParseInt(string(v.Lit), 10, 64)
	default:
		return 0, fmt.Errorf("unexpected type for val: %T", val)
	}
}
