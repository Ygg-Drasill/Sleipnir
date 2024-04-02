package ast

type Attribute interface{}

type (
	NodeList       []Node
	ConnectionList []Connection
)

type Program struct {
	Nodes       NodeList
	Connections ConnectionList
}

type Node struct {
	id string
}

type Connection struct {
	outId Junction
	inId  Junction
}

type Junction struct {
	nodeId string
	varId  string
}
