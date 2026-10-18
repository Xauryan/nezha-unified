// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteControlGroupRuleBody struct {
    position.Body
    // 访问控制组规则ID列表
	RuleLists []string 
}

func (s DeleteControlGroupRuleBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteControlGroupRuleBody) GoString() string {
	return s.String()
}

func (s DeleteControlGroupRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteControlGroupRuleBody) SetRuleLists(v []string) *DeleteControlGroupRuleBody {
	s.RuleLists = v
	return s
}


type DeleteControlGroupRuleBodyBuilder struct {
	s *DeleteControlGroupRuleBody
}

func NewDeleteControlGroupRuleBodyBuilder() *DeleteControlGroupRuleBodyBuilder {
	s := &DeleteControlGroupRuleBody{}
	b := &DeleteControlGroupRuleBodyBuilder{s: s}
	return b
}

func (b *DeleteControlGroupRuleBodyBuilder) RuleLists(v []string) *DeleteControlGroupRuleBodyBuilder {
	b.s.RuleLists = v
	return b
}

func (b *DeleteControlGroupRuleBodyBuilder) Build() *DeleteControlGroupRuleBody {
	return b.s
}


