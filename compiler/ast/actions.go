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

func NewNode(context, node, in, out, process Attribute) (Node, error) {
	nodeId := string(node.(*token.Token).Lit)
	ctx := context.(ParseContext)
	ctx.BabushkaPopScopeNode(nodeId)

	return Node{
		id:              nodeId,
		inDeclarations:  nil,
		outDeclarations: nil,
		procStatements:  nil,
	}, nil
}

func NewScopeIn(context Attribute) (Attribute, error) {
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

func NewDeclarationEmpty(varType, varId, context Attribute) (Declaration, error) {
	ctx := context.(ParseContext)
	varTypeStr := string(varType.(*token.Token).Lit)
	varIdStr := string(varId.(*token.Token).Lit)
	ctx.CurrentScope.AddVariable(varIdStr)

	return Declaration{
		Type:       varTypeStr,
		AssigneeId: varIdStr,
		Expression: nil,
	}, nil
}

func NewDeclaration(varType, varId, expression, context Attribute) (Declaration, error) {
	ctx := context.(ParseContext)
	varIdStr := string(varId.(*token.Token).Lit)
	varTypeStr := string(varType.(*token.Token).Lit)
	ctx.CurrentScope.AddVariable(varIdStr)

	return Declaration{
		Type:       varTypeStr,
		AssigneeId: varIdStr,
		Expression: expression,
	}, nil
}

func NewDeclarationList(declaration Attribute) (DeclarationList, error) {
	firstDeclaration := declaration.(Declaration)
	return DeclarationList{firstDeclaration}, nil
}

func AppendDeclaration(declarationList, declaration Attribute) (DeclarationList, error) {
	return append(declarationList.(DeclarationList), declaration.(Declaration)), nil
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
