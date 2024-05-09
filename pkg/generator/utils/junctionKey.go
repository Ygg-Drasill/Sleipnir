package utils

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/ast"
)

func JunctionToKey(junction ast.Junction) string {
	return JunctionKey(junction.NodeId, junction.VarId)
}

func JunctionKey(nodeId, varId string) string {
	return fmt.Sprintf("%s-%s", nodeId, varId)
}
