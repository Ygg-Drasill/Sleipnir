package ast

type Attribute interface{}

type (
	NodeList       []Node
	ConnectionList []Connection
)

type Statement interface {
}

type (
	StatementList   []Statement
	DeclarationList []Declaration
)

type Declaration struct {
	Type       string     `json:"type"`
	AssigneeId Attribute  `json:"assigneeId"`
	Expression Expression `json:"expression"`
}

type Expression Attribute

type Program struct {
	Nodes       NodeList       `json:"nodes"`
	Connections ConnectionList `json:"connections"`
}

type Node struct {
	Id              string          `json:"id"`
	InDeclarations  []StatementList `json:"inDeclarations"`
	OutDeclarations []StatementList `json:"outDeclarations"`
	ProcStatements  []StatementList `json:"procStatements"`
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
	VarId  string `json:"VarId"`
}
