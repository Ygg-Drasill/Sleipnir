package ast

import (
	"errors"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
)

//NodeVar : in "." Id | out "." Id | Id  ;

func NewNodeVar(ioType Attribute, varId Attribute) (NodeVar, error) {
	varIdStr := string(varId.(*token.Token).Lit)
	ioTypeStr := string(ioType.(*token.Token).Lit)
	return NodeVar{JunctionType: ioTypeStr, Node: nil, Id: varIdStr}, nil
}

func NewLocalVar(id Attribute) (*LocalVar, error) {
	idToken, ok := id.(*token.Token)
	if !ok {
		return nil, errors.New("identifier expected")
	}
	return &LocalVar{Id: string(idToken.Lit)}, nil
}
