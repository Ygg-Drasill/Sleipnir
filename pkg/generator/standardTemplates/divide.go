package standardTemplates

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/utils"
)

var Divide standardTemplate = standardTemplate{
	Body: `%s
%s
i64.div_s
%s`,
	Inputs:  []string{"a", "b"},
	Outputs: []string{"result"},
	FormatBody: func(t standardTemplate, nodeId string, varMap map[string]*ast.Junction) string {
		a := varMap[utils.JunctionKey(nodeId, "a")]
		b := varMap[utils.JunctionKey(nodeId, "b")]
		return fmt.Sprintf(t.Body,
			mapVarJunctionVariableGet(a),
			mapVarJunctionVariableGet(b),
			mapVarJunctionVariableSet(nodeId, t.Outputs[0]),
		)
	},
}
