// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateControlGroupRuleBody struct {
    position.Body
    // 地址类型（IPv4或IPv6）
	IpVersion *string `json:"ipVersion,omitempty"`
    // IP地址
	IpAddress *string `json:"ipAddress,omitempty"`
    // 访问控制组规则描述
	Description *string `json:"description,omitempty"`
    // 所属控制组ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
}

func (s CreateControlGroupRuleBody) String() string {
	return utils.Beautify(s)
}

func (s CreateControlGroupRuleBody) GoString() string {
	return s.String()
}

func (s CreateControlGroupRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateControlGroupRuleBody) SetIpVersion(v string) *CreateControlGroupRuleBody {
	s.IpVersion = &v
	return s
}

func (s *CreateControlGroupRuleBody) SetIpAddress(v string) *CreateControlGroupRuleBody {
	s.IpAddress = &v
	return s
}

func (s *CreateControlGroupRuleBody) SetDescription(v string) *CreateControlGroupRuleBody {
	s.Description = &v
	return s
}

func (s *CreateControlGroupRuleBody) SetControlGroupId(v string) *CreateControlGroupRuleBody {
	s.ControlGroupId = &v
	return s
}


type CreateControlGroupRuleBodyBuilder struct {
	s *CreateControlGroupRuleBody
}

func NewCreateControlGroupRuleBodyBuilder() *CreateControlGroupRuleBodyBuilder {
	s := &CreateControlGroupRuleBody{}
	b := &CreateControlGroupRuleBodyBuilder{s: s}
	return b
}

func (b *CreateControlGroupRuleBodyBuilder) IpVersion(v string) *CreateControlGroupRuleBodyBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *CreateControlGroupRuleBodyBuilder) IpAddress(v string) *CreateControlGroupRuleBodyBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *CreateControlGroupRuleBodyBuilder) Description(v string) *CreateControlGroupRuleBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateControlGroupRuleBodyBuilder) ControlGroupId(v string) *CreateControlGroupRuleBodyBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *CreateControlGroupRuleBodyBuilder) Build() *CreateControlGroupRuleBody {
	return b.s
}


