package ast

import (
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
)

func NewNodeVar(ioType Attribute, varId Attribute) (NodeVar, error) {
	ioTypeStr := ""
	if ioType != nil {
		ioTypeStr = string(ioType.(*token.Token).Lit)
	}
	varIdStr := string(varId.(*token.Token).Lit)
	return NodeVar{ioType: ioTypeStr, varId: varIdStr}, nil
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
