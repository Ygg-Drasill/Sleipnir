package ast

type varKey string
type nodeKey string
type symbolTable map[varKey]*Variable

type ParseContext struct {
	Nodes        map[nodeKey]*NodeContext
	CurrentScope symbolTable
	CurrentNode  *NodeContext
}

type NodeContext struct {
	InVariables   symbolTable
	OutVariables  symbolTable
	ProcVariables symbolTable
}

type Variable struct {
	Type  string
	Value []byte
}

func NewParseContext() ParseContext {
	return ParseContext{
		Nodes:        make(map[nodeKey]*NodeContext),
		CurrentScope: nil,
		CurrentNode:  nil,
	}
}

func NewNodeContext() *NodeContext {
	return &NodeContext{
		InVariables:   make(symbolTable),
		OutVariables:  make(symbolTable),
		ProcVariables: make(symbolTable),
	}
}

func (pc *ParseContext) AddNodeContext(nodeName nodeKey, nodeContext *NodeContext) {
	pc.Nodes[nodeName] = nodeContext
}

func (pc *ParseContext) NewScope() {
	pc.CurrentScope = make(symbolTable)
}

func (pc *ParseContext) NewNodeScope() {
	pc.CurrentNode = NewNodeContext()
}
