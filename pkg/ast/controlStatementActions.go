package ast

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"
)

func NewIfStatement(condition, bodyBlock, elseBlock Attribute) (*IfStatement, error) {
	var conditionExpression Expression
	var bodyStatementList StatementList
	var elseStatementList StatementList

	conditionExpression = condition.(Expression)

	op := string(conditionExpression.Operator.(*token.Token).Lit)
	if op == "+" || op == "-" || op == "/" || op == "*" || op == "%" {
		return nil, fmt.Errorf("if statement must contain condition")
	}
	if bodyBlock != nil {
		bodyStatementList = bodyBlock.(StatementList)
	}
	if elseBlock != nil {
		elseStatementList = elseBlock.(StatementList)
	}

	return &IfStatement{
		Expression:     conditionExpression,
		BodyStatements: bodyStatementList,
		ElseStatements: elseStatementList,
	}, nil
}
