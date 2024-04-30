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
	Type       string    `json:"type"`
	AssigneeId Attribute `json:"assigneeId"`
	Expression Attribute `json:"expression"`
}

type Expression struct {
	FirstOperand  Attribute `json:"firstOperand"`
	SecondOperand Attribute `json:"secondOperand"`
	Operator      Attribute `json:"operator"`
}

type Program struct {
	Nodes       NodeList       `json:"nodes"`
	Connections ConnectionList `json:"connections"`
}

type Node struct {
	Id              string          `json:"id"`
	InDeclarations  DeclarationList `json:"inDeclarations"`
	OutDeclarations DeclarationList `json:"outDeclarations"`
	ProcStatements  StatementList   `json:"procStatements"`
}

type NodeVar struct {
	IoType string `json:"ioType'"`
	VarId  string `json:"varId"`
}

type Connection struct {
	OutId Junction `json:"outId"`
	InId  Junction `json:"inId"`
}

type Junction struct {
	NodeId string `json:"nodeId"`
	VarId  string `json:"varId"`
}

type IfStatement struct {
	Expression     Expression    `json:"expression"`
	BodyStatements StatementList `json:"bodyStatements"`
	ElseStatements StatementList `json:"elseStatements"`
}

type WhileStatement struct {
	Condition      Expression    `json:"condition"`
	BodyStatements StatementList `json:"bodyStatements"`
}

type Assignment struct {
	Expression Expression `json:"expression"`
}

type AssignmentStatement struct {
	Expression Expression `json:"expression"`
	Identifier string     `json:"identifier"`
	Assignment Assignment `json:"assignment"`
}
