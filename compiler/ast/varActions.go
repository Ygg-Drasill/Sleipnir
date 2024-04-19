package ast

import "github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"

//NodeVar : in "." varId | out "." varId | varId  ;

func NewNodeVar(ioType Attribute, varId Attribute) (NodeVar, error) {
	varIdStr := string(varId.(*token.Token).Lit)
	ioTypeStr := string(ioType.(*token.Token).Lit)
	return NodeVar{ioType: ioTypeStr, varId: varIdStr}, nil
}
