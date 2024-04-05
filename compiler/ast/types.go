package ast

type Attribute interface{}

type (
	NodeList       []Node
	ConnectionList []Connection
)

type (
	StatementList   []Statement
	Statement       Attribute
	DeclarationList []Declaration
)

type Declaration struct {
	Assignee   Attribute
	Expression Attribute
}

type Program struct {
	Nodes       NodeList
	Connections ConnectionList
}

type Node struct {
	id    string
	value int64
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
