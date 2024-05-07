package generator

import "github.com/Ygg-Drasill/Sleipnir/pkg/ast"

func (g *Generator) isIdentifier(attr *ast.Attribute) (Identifier, bool) {

	if i, ok := (*attr).(ast.NodeVar); ok {
		sourceJunction := g.outNodeVars[junctionKey(g.currentNode.Id, i.Id)]
		newIdentifier := NodeIdentifier{
			NodeVar:        i,
			sourceJunction: sourceJunction,
			nodeId:         g.currentNode.Id,
		}
		return newIdentifier, true
	}

	if i, ok := (*attr).(ast.LocalVar); ok {
		newIdentifier := LocalIdentifier{
			i,
		}
		return newIdentifier, true
	}

	return nil, false
}
