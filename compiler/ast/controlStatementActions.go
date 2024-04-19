package ast

func NewIfStatement(condition, bodyBlock, elseBlock Attribute) (IfStatement, error) {
	var conditionExpression Expression
	var bodyStatementList StatementList
	var elseStatementList StatementList

	conditionExpression = condition.(Expression)

	if bodyBlock != nil {
		bodyStatementList = bodyBlock.(StatementList)
	}
	if elseBlock != nil {
		elseStatementList = elseBlock.(StatementList)
	}

	return IfStatement{
		condition:      conditionExpression,
		bodyStatements: bodyStatementList,
		elseStatements: elseStatementList,
	}, nil
}

func NewWhileStatement(condition, bodyBlock Attribute) (WhileStatement, error) {
	conditionExpression := condition.(Expression)
	statementList := bodyBlock.(StatementList)
	return WhileStatement{
		condition:      conditionExpression,
		bodyStatements: statementList,
	}, nil
}
