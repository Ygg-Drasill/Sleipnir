package standardTemplates

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/utils"
)

var Move standardTemplate = standardTemplate{
	Body: `global.get $%s
i32.wrap_i64
call $move`,
	Inputs: []string{"move"},
	FormatBody: func(t standardTemplate, nodeId string, nodeVarMap map[string]*ast.Junction) string {
		text := nodeVarMap[utils.JunctionKey(nodeId, t.Inputs[0])]
		return fmt.Sprintf(t.Body, mapVarJunctionVariable(text.NodeId, text.VarId))
	},
}