package ast

type Attribute interface{}

type (
	NodeList       []Node
	ConnectionList []Connection
)

type Statement Attribute

type (
	StatementList   []Statement
	DeclarationList []Declaration
)

type Declaration struct {
	Type       string
	AssigneeId Attribute
	Expression Attribute
}

type Expression struct {
	FirstOperand  Attribute
	SecondOperand Attribute
	Operator      Attribute
}

type Program struct {
	Nodes       NodeList
	Connections ConnectionList
}

type Node struct {
	Id              string
	inDeclarations  DeclarationList
	outDeclarations DeclarationList
	procStatements  StatementList
}

type NodeVar struct {
	ioType string
	varId  string
	value  int64
}

type Connection struct {
	outId Junction
	inId  Junction
}

type Junction struct {
	nodeId string
	varId  string
}

type IfStatement struct {
	condition      Expression
	bodyStatements StatementList
	elseStatements StatementList
}

type WhileStatement struct {
	condition      Expression
	bodyStatements StatementList
}
