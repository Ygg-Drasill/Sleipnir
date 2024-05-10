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
	TemplateId      string          `json:"templateId"`
	InDeclarations  DeclarationList `json:"inDeclarations"`
	OutDeclarations DeclarationList `json:"outDeclarations"`
	ProcStatements  StatementList   `json:"procStatements"`
}

type NodeVar struct {
	Id           string `json:"id"`
	JunctionType string `json:"junctionType'"`
}

type LocalVar struct {
	Id string `json:"id"`
}

type Connection struct {
	OutJunction Junction `json:"outId"`
	InJunction  Junction `json:"inId"`
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

type AssignmentStatement struct {
	Identifier Attribute `json:"identifier"`
	Expression Attribute `json:"expression"`
}

type Template struct {
	Node
}

type TemplateUse struct {
	TemplateId string `json:"templateId"`
	NodeId     string `json:"nodeId"`
}
