package generator

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

type IdentifierConstraint interface {
	ast.NodeVar | ast.LocalVar
}

type Identifier interface {
	getLabel(id string) string
}

type NodeIdentifier struct {
	ast.NodeVar
}

type LocalIdentifier struct {
	ast.LocalVar
}

func (identifier NodeIdentifier) getLabel(nodeId string) string {
	return fmt.Sprintf("test\n")
}

func (identifier LocalIdentifier) getLabel(id string) string {
	return id
}
