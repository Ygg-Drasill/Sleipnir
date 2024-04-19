package ast

import "github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"

//NodeVar : in "." VarId | out "." VarId | VarId  ;

func NewNodeVar(ioType Attribute, varId Attribute) (NodeVar, error) {
	varIdStr := string(varId.(*token.Token).Lit)
	ioTypeStr := string(ioType.(*token.Token).Lit)
	return NodeVar{IoType: ioTypeStr, VarId: varIdStr}, nil
}
