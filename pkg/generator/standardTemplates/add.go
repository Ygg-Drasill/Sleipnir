package standardTemplates

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/utils"
)

var Add standardTemplate = standardTemplate{
	Body: `global.get $%s
global.get $%s
i64.add
global.set $%s`,
	Inputs:  []string{"a", "b"},
	Outputs: []string{"result"},
	FormatBody: func(t standardTemplate, nodeId string, varMap map[string]*ast.Junction) string {
		a := varMap[utils.JunctionKey(nodeId, "a")]
		b := varMap[utils.JunctionKey(nodeId, "b")]
		return fmt.Sprintf(t.Body,
			mapVarJunctionVariable(a.NodeId, a.VarId),
			mapVarJunctionVariable(b.NodeId, b.VarId),
			mapVarJunctionVariable(nodeId, t.Outputs[0]),
		)
	},
}
