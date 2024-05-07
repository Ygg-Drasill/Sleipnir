package generator

import "github.com/Ygg-Drasill/Sleipnir/pkg/ast"

func isIdentifier(attr *ast.Attribute) (Identifier, bool) {
	if i, ok := (*attr).(ast.NodeVar); ok {
		newIdentifier := NodeIdentifier{
			i,
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
