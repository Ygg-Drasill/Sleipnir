package standardTemplates

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
	"github.com/Ygg-Drasill/Sleipnir/pkg/generator/utils"
)

var SetMemory = standardTemplate{
	Body: `%s
i32.wrap_i64
call $_set`,
	Inputs: []string{"value"},
	FormatBody: func(t standardTemplate, nodeId string, nodeVarMap map[string]*ast.Junction) string {
		value := nodeVarMap[utils.JunctionKey(nodeId, t.Inputs[0])]
		return fmt.Sprintf(t.Body, mapVarJunctionVariableGet(value))
	},
}

var GetMemory = standardTemplate{
	Body: `call $_get
i64.extend_i32_s
%s`,
	Outputs: []string{"value"},
	FormatBody: func(t standardTemplate, nodeId string, nodeVarMap map[string]*ast.Junction) string {
		return fmt.Sprintf(t.Body, mapVarJunctionVariableSet(nodeId, t.Outputs[0]))
	},
}
