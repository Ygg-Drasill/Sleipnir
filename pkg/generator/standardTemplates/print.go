package standardTemplates

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/utils"
)

var Print standardTemplate = standardTemplate{
	Body: `%s
i32.wrap_i64
call $_log`,
	Inputs: []string{"text"},
	FormatBody: func(t standardTemplate, nodeId string, nodeVarMap map[string]*ast.Junction) string {
		text := nodeVarMap[utils.JunctionKey(nodeId, t.Inputs[0])]
		return fmt.Sprintf(t.Body, mapVarJunctionVariableGet(text))
	},
}
