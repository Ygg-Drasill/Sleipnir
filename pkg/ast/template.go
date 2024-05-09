package ast

import "github.com/Ygg-Drasill/Sleipnir/pkg/gocc/token"

func UseTemplate(nodeId, templateId Attribute) (TemplateUse, error) {
	return TemplateUse{
		TemplateId: string(templateId.(*token.Token).Lit),
		NodeId:     string(nodeId.(*token.Token).Lit),
	}, nil
}
