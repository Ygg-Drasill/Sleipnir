package ast

type varKey string
type nodeKey string
type symbolTable map[varKey]*Variable

func (table *symbolTable) AddVariable(id string) {
	(*table)[varKey(id)] = newVariable("int")
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

func newSymbolTable() symbolTable {
	return make(symbolTable)
}

func NewParseContext() ParseContext {
	return ParseContext{
		Nodes:        make(map[nodeKey]*NodeContext),
		CurrentScope: newSymbolTable(),
		CurrentNode:  newNodeContext(),
	}
}

func newNodeContext() *NodeContext {
	return &NodeContext{
		InVariables:   newSymbolTable(),
		OutVariables:  newSymbolTable(),
		ProcVariables: newSymbolTable(),
	}
}

type ParseContext struct {
	Nodes        map[nodeKey]*NodeContext
	CurrentNode  *NodeContext
	CurrentScope symbolTable
}

func (ctx *ParseContext) AddNodeContext(nodeName nodeKey, nodeContext *NodeContext) {
	ctx.Nodes[nodeName] = nodeContext
}

func (ctx *ParseContext) NewScope() {
	ctx.CurrentScope = newSymbolTable()
}

func (ctx *ParseContext) NewNodeScope() {
	ctx.CurrentNode = newNodeContext()
}

func (ctx *ParseContext) BabushkaPopScopeIn() {
	ctx.CurrentNode.InVariables = ctx.CurrentScope
	ctx.NewScope()
}

func (ctx *ParseContext) BabushkaPopScopeOut() {
	ctx.CurrentNode.OutVariables = ctx.CurrentScope
	ctx.NewScope()
}

func (ctx *ParseContext) BabushkaPopScopeProc() {
	ctx.CurrentNode.ProcVariables = ctx.CurrentScope
	ctx.NewScope()
}

func (ctx *ParseContext) BabushkaPopScopeNode(nodeId string) {
	ctx.Nodes[nodeKey(nodeId)] = ctx.CurrentNode
	ctx.NewNodeScope()
}
