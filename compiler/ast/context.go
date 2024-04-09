package ast

type varKey string
type nodeKey string
type symbolTable map[varKey]*Variable

type ParseContext struct {
	Nodes        map[nodeKey]*NodeContext
	CurrentNode  *NodeContext
	CurrentScope symbolTable
}

type NodeContext struct {
	InVariables   symbolTable
	OutVariables  symbolTable
	ProcVariables symbolTable
}

type Variable struct {
	Type string
}

func newVariable(variableType string) *Variable {
	return &Variable{
		Type: variableType,
	}
}

func NewParseContext() ParseContext {
	return ParseContext{
		Nodes:        make(map[nodeKey]*NodeContext),
		CurrentScope: make(symbolTable),
		CurrentNode:  NewNodeContext(),
	}
}

func NewNodeContext() *NodeContext {
	return &NodeContext{
		InVariables:   make(symbolTable),
		OutVariables:  make(symbolTable),
		ProcVariables: make(symbolTable),
	}
}

func (ctx ParseContext) AddNodeContext(nodeName nodeKey, nodeContext *NodeContext) {
	ctx.Nodes[nodeName] = nodeContext
}

func (ctx ParseContext) NewScope() {
	ctx.CurrentScope = make(symbolTable)
}

func (ctx ParseContext) NewNodeScope() {
	ctx.CurrentNode = NewNodeContext()
}

func (table symbolTable) AddVariable(id string) {
	table[varKey(id)] = newVariable("int")
}
