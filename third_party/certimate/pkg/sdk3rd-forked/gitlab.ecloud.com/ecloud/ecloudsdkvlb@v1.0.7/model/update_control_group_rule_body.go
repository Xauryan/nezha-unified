// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateControlGroupRuleBody struct {
    position.Body
    // 访问控制组规则描述
	Description *string `json:"description,omitempty"`
    // 控制组规则ID
	Id *string `json:"id,omitempty"`
    // 所属控制组ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
}

func (s UpdateControlGroupRuleBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateControlGroupRuleBody) GoString() string {
	return s.String()
}

func (s UpdateControlGroupRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateControlGroupRuleBody) SetDescription(v string) *UpdateControlGroupRuleBody {
	s.Description = &v
	return s
}

func (s *UpdateControlGroupRuleBody) SetId(v string) *UpdateControlGroupRuleBody {
	s.Id = &v
	return s
}

func (s *UpdateControlGroupRuleBody) SetControlGroupId(v string) *UpdateControlGroupRuleBody {
	s.ControlGroupId = &v
	return s
}


type UpdateControlGroupRuleBodyBuilder struct {
	s *UpdateControlGroupRuleBody
}

func NewUpdateControlGroupRuleBodyBuilder() *UpdateControlGroupRuleBodyBuilder {
	s := &UpdateControlGroupRuleBody{}
	b := &UpdateControlGroupRuleBodyBuilder{s: s}
	return b
}

func (b *UpdateControlGroupRuleBodyBuilder) Description(v string) *UpdateControlGroupRuleBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateControlGroupRuleBodyBuilder) Id(v string) *UpdateControlGroupRuleBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *UpdateControlGroupRuleBodyBuilder) ControlGroupId(v string) *UpdateControlGroupRuleBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *UpdateControlGroupRuleBodyBuilder) Build() *UpdateControlGroupRuleBody {
	return b.s
}


