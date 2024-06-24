package standardTemplates

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

type bodyFormatFunction func(template standardTemplate, nodeId string, nodeVarMap map[string]*ast.Junction) string

type standardTemplate struct {
	Body       string
	Outputs    []string
	Inputs     []string
	FormatBody bodyFormatFunction
}

var StandardTemplates = map[string]*standardTemplate{
	"Print": &Print,
	"Add":   &Add,
	"Sub":   &Sub,
	"Mul":   &Mul,
	"Div":   &Div,
	"Mod":   &Mod,
	"Move":  &Move,
}

func mapVarJunctionVariable(nodeId, varId string) string {
	return fmt.Sprintf("%s_%s", nodeId, varId)
}
