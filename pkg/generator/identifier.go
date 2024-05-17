package generator

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

type IdentifierConstraint interface {
	ast.NodeVar | ast.LocalVar
}

type Identifier interface {
	toGetInstruction() string
	toSetInstruction() string
}

type NodeInIdentifier struct {
	ast.NodeVar
	sourceJunction *ast.Junction
	nodeId         string
}

type NodeOutIdentifier struct {
	ast.NodeVar
	nodeId string
}

type LocalIdentifier struct {
	ast.LocalVar
}

func (identifier NodeInIdentifier) toGetInstruction() string {
	return fmt.Sprintf("global.get $%s_%s", identifier.sourceJunction.NodeId, identifier.sourceJunction.VarId)
}

func (identifier NodeOutIdentifier) toGetInstruction() string {
	return fmt.Sprintf("global.get $%s_%s", identifier.nodeId, identifier.Id)
}

func (identifier LocalIdentifier) toGetInstruction() string {
	return fmt.Sprintf("local.get $%s", identifier.Id)
}

func (identifier NodeInIdentifier) toSetInstruction() string {
	return fmt.Sprintf("global.set $%s_%s", identifier.nodeId, identifier.Id)
}

func (identifier NodeOutIdentifier) toSetInstruction() string {
	return fmt.Sprintf("global.set $%s_%s", identifier.nodeId, identifier.Id)
}

func (identifier LocalIdentifier) toSetInstruction() string {
	return fmt.Sprintf("local.set $%s", identifier.Id)
}
