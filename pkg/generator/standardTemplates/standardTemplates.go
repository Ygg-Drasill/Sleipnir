package standardTemplates

import "github.com/Ygg-Drasill/Sleipnir/pkg/ast"

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
}
