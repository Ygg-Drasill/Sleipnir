package ast

import (
	"fmt"
	"strconv"

	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
)

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
	switch v := node.(type) {
	case *token.Token:
		return Node{id: string(v.Lit)}, nil
	case NodeVar:
		return Node{id: v.varId, value: v.value}, nil
	default:
		return Node{}, fmt.Errorf("unknown type for node: %T", v)
	}
}

func NewNodeVar(ioType Attribute, varId Attribute) (NodeVar, error) {
	ioTypeStr := ""
	if ioType != nil {
		ioTypeStr = string(ioType.(*token.Token).Lit)
	}
	varIdStr := string(varId.(*token.Token).Lit)
	return NodeVar{ioType: ioTypeStr, varId: varIdStr}, nil
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
	varIdStr := string(varId.(*token.Token).Lit)
	expressionStr := string(expression.(*token.Token).Lit)
	return DeclarationList{Declaration{
		Assignee:   varIdStr,
		Expression: expressionStr,
	}}, nil
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
	case NodeVar:
		return v.value, nil
	default:
		return 0, fmt.Errorf("unexpected type for val: %T", val)
	}
}

func Add(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 + intVal2, nil
}

func Sub(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 - intVal2, nil
}

func Mul(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 * intVal2, nil
}

func Div(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 / intVal2, nil
}
