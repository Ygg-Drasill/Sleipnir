package generator

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

type IdentifierConstraint interface {
	ast.NodeVar | ast.LocalVar
}

type Identifier interface {
	getOperation() string
	setOperation() string
}

type NodeIdentifier struct {
	ast.NodeVar
	sourceJunction ast.Junction
	nodeId         string
}

type LocalIdentifier struct {
	ast.LocalVar
}

func (identifier NodeIdentifier) getOperation() string {
	return fmt.Sprintf("global.get $%s_%s", identifier.sourceJunction.NodeId, identifier.sourceJunction.VarId)
}

func (identifier LocalIdentifier) getOperation() string {
	return fmt.Sprintf("local.get $%s", identifier.Id)
}

func (identifier NodeIdentifier) setOperation() string {
	return fmt.Sprintf("global.set $%s_%s", identifier.nodeId, identifier.Id)
}

func (identifier LocalIdentifier) setOperation() string {
	return fmt.Sprintf("local.set $%s", identifier.Id)
}
