package generator

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/utils"
)

func (g *Generator) isIdentifier(attr *ast.Attribute) (Identifier, error) {
	if i, ok := (*attr).(ast.NodeVar); ok {
		var newIdentifier Identifier

		if i.JunctionType == "out" {
			newIdentifier = NodeOutIdentifier{
				NodeVar: i,
				nodeId:  g.currentNode.Id,
			}
		}

		if i.JunctionType == "in" {
			sourceJunction := g.outNodeVars[utils.JunctionKey(g.currentNode.Id, i.Id)]

			if sourceJunction == nil {
				return nil, fmt.Errorf("reference to undeclared node variable %s", utils.JunctionKey(g.currentNode.Id, i.Id))
			}

			newIdentifier = NodeInIdentifier{
				NodeVar:        i,
				sourceJunction: sourceJunction,
				nodeId:         g.currentNode.Id,
			}
		}
		return newIdentifier, nil
	}

	if i, ok := (*attr).(*ast.LocalVar); ok {
		newIdentifier := LocalIdentifier{
			*i,
		}
		return newIdentifier, nil
	}

	return nil, nil
}
