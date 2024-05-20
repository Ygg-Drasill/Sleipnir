package ast

import (
	"errors"
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
)

//NodeVar : in "." Id | out "." Id | Id  ;

func NewNodeVar(ioType, varId Attribute) (NodeVar, error) {
	varIdStr := string(varId.(*token.Token).Lit)
	ioTypeStr := string(ioType.(*token.Token).Lit)
	return NodeVar{JunctionType: ioTypeStr, Id: varIdStr}, nil
}

func NewLocalVar(id Attribute) (*LocalVar, error) {
	idToken, ok := id.(*token.Token)
	varName := string(idToken.Lit)
	err := checkLocal(varName)
	if !ok {
		return nil, errors.New("identifier expected")
	}
	return &LocalVar{Id: varName}, err
}

var LocalDeclarationError = errors.New("local declaration error")

func checkLocal(varName string) error {
	var err error
	if varName == "out" || varName == "in" {
		err = errors.Join(LocalDeclarationError, fmt.Errorf("the variable name %s is not allowed in a process",
			varName,
		))
	}

	return err
}
