package ast

import "github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"

func UseTemplate(nodeId, templateId Attribute) (Node, error) {
	return Node{
		TemplateId: string(templateId.(*token.Token).Lit),
		Id:         string(nodeId.(*token.Token).Lit),
	}, nil
}
