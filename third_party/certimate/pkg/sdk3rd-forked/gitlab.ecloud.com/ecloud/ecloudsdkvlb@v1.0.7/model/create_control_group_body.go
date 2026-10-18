// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateControlGroupBody struct {
    position.Body
    // 访问控制规则地址类型（IPv4或IPv6）
	IpVersion *string `json:"ipVersion,omitempty"`
    // 访问控制规则IP地址
	IpAddress *string `json:"ipAddress,omitempty"`
    // 访问控制组名称，5~64个字符，以英文大小写字母开头，可包含数字、下划线（_）和连字符（-）；不填则自动设置默认值
	Name *string `json:"name,omitempty"`
    // 访问控制组描述
	Description *string `json:"description,omitempty"`
    // 访问控制组规则描述
	RuleDescription *string `json:"ruleDescription,omitempty"`
}

func (s CreateControlGroupBody) String() string {
	return utils.Beautify(s)
}

func (s CreateControlGroupBody) GoString() string {
	return s.String()
}

func (s CreateControlGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateControlGroupBody) SetIpVersion(v string) *CreateControlGroupBody {
	s.IpVersion = &v
	return s
}

func (s *CreateControlGroupBody) SetIpAddress(v string) *CreateControlGroupBody {
	s.IpAddress = &v
	return s
}

func (s *CreateControlGroupBody) SetName(v string) *CreateControlGroupBody {
	s.Name = &v
	return s
}

func (s *CreateControlGroupBody) SetDescription(v string) *CreateControlGroupBody {
	s.Description = &v
	return s
}

func (s *CreateControlGroupBody) SetRuleDescription(v string) *CreateControlGroupBody {
	s.RuleDescription = &v
	return s
}


type CreateControlGroupBodyBuilder struct {
	s *CreateControlGroupBody
}

func NewCreateControlGroupBodyBuilder() *CreateControlGroupBodyBuilder {
	s := &CreateControlGroupBody{}
	b := &CreateControlGroupBodyBuilder{s: s}
	return b
}

func (b *CreateControlGroupBodyBuilder) IpVersion(v string) *CreateControlGroupBodyBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *CreateControlGroupBodyBuilder) IpAddress(v string) *CreateControlGroupBodyBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *CreateControlGroupBodyBuilder) Name(v string) *CreateControlGroupBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateControlGroupBodyBuilder) Description(v string) *CreateControlGroupBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateControlGroupBodyBuilder) RuleDescription(v string) *CreateControlGroupBodyBuilder {
	b.s.RuleDescription = &v
	return b
}

func (b *CreateControlGroupBodyBuilder) Build() *CreateControlGroupBody {
	return b.s
}


