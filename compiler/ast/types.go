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
	Type       string
	AssigneeId Attribute
	Expression Expression
}

type Expression Attribute

type Program struct {
	Nodes       NodeList
	Connections ConnectionList
}

type Node struct {
	id              string
	inDeclarations  []StatementList
	outDeclarations []StatementList
	procStatements  []StatementList
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
