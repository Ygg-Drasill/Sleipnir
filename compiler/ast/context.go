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

func copyNodeContext(src *NodeContext) *NodeContext {
	newNode := newNodeContext()
	moveSymbolTable(src.InVariables, newNode.InVariables)
	moveSymbolTable(src.OutVariables, newNode.OutVariables)
	moveSymbolTable(src.ProcVariables, newNode.ProcVariables)
	return newNode
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
	moveSymbolTable(ctx.CurrentScope, ctx.CurrentNode.InVariables)
}

func (ctx *ParseContext) BabushkaPopScopeOut() {
	moveSymbolTable(ctx.CurrentScope, ctx.CurrentNode.OutVariables)
}

func (ctx *ParseContext) BabushkaPopScopeProc() {
	moveSymbolTable(ctx.CurrentScope, ctx.CurrentNode.ProcVariables)
}

func (ctx *ParseContext) BabushkaPopScopeNode(nodeId string) {
	ctx.Nodes[nodeKey(nodeId)] = copyNodeContext(ctx.CurrentNode)
	ctx.NewNodeScope()

}
