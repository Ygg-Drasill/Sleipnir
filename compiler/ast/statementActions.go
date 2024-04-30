package ast

import (
	"errors"
	"github.com/Ygg-Drasill/Sleipnir/compiler/gocc/token"
	"go/ast"
)

func NewDeclaration(context, varType, varId, expression Attribute) (Declaration, error) {
	ctx := context.(ParseContext)
	varIdStr := string(varId.(*token.Token).Lit)
	varTypeStr := string(varType.(*token.Token).Lit)
	ctx.CurrentScope.AddVariable(varIdStr)

	if !(varTypeStr == "int" || varTypeStr == "bool" || varTypeStr == "string") {
		return Declaration{}, errors.New("Invalid variable type: " + varTypeStr)
	}

	return Declaration{
		Type:       varTypeStr,
		AssigneeId: varIdStr,
		Expression: expression,
	}, nil
}

func NewDeclarationList() (DeclarationList, error) {
	return DeclarationList{}, nil
}

func AppendDeclaration(declarationList, declaration Attribute) (DeclarationList, error) {
	return append(declarationList.(DeclarationList), declaration.(Declaration)), nil
}

func NewStatementList(statement Attribute) (StatementList, error) {
	firstStatement := statement.(Statement)
	return StatementList{firstStatement}, nil
}

func AppendStatement(statementList, statement Attribute) (StatementList, error) {
	return append(statementList.(StatementList), statement.(Statement)), nil
}

func NewAssignmentStatement(assignment Attribute) (AssignmentStatement, error) {
	return AssignmentStatement{
		Expression: Expression{
			FirstOperand:  nil,
			SecondOperand: nil,
			Operator:      nil,
		},
		Identifier: "",
		Assignment: ,
	}, nil
}
