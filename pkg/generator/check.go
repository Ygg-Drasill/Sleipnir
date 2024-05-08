package generator

import (
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"log/slog"
)

func (g *Generator) isIdentifier(attr *ast.Attribute) (Identifier, bool) {
	if i, ok := (*attr).(ast.NodeVar); ok {
		var newIdentifier Identifier

		if i.JunctionType == "out" {
			newIdentifier = NodeOutIdentifier{
				NodeVar: i,
				nodeId:  g.currentNode.Id,
			}
		}

		if i.JunctionType == "in" {
			sourceJunction := g.outNodeVars[junctionKey(g.currentNode.Id, i.Id)]

			if sourceJunction == nil {
				slog.Error("Reference to undeclared node variable", junctionKey(g.currentNode.Id, i.Id))
				return nil, false
			}

			newIdentifier = NodeInIdentifier{
				NodeVar:        i,
				sourceJunction: sourceJunction,
				nodeId:         g.currentNode.Id,
			}
		}
		return newIdentifier, true
	}

	if i, ok := (*attr).(*ast.LocalVar); ok {
		newIdentifier := LocalIdentifier{
			*i,
		}
		return newIdentifier, true
	}

	return nil, false
}
