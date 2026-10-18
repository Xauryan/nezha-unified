// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateTCAclBody struct {
    position.Body
    // 访问控制规则地址类型（IPv4或IPv6）
	IpVersion *string `json:"ipVersion,omitempty"`
    // 访问控制规则IP地址
	IpAddress *string `json:"ipAddress,omitempty"`
    // 访问控制组名称，5~64个字符，以英文大小写字母开头，可包含数字、下划线（_）和连字符（-）；不填则自动设置默认值
	Name *string `json:"name,omitempty"`
    // 访问控制组描述
	Description *string `json:"description,omitempty"`
    // 访问控制组ID
	Id *string `json:"id,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 访问控制组描述
	RuleDescription *string `json:"ruleDescription,omitempty"`
    // 边缘云字段
	Vaz *string `json:"vaz,omitempty"`
}

func (s CreateTCAclBody) String() string {
	return utils.Beautify(s)
}

func (s CreateTCAclBody) GoString() string {
	return s.String()
}

func (s CreateTCAclBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateTCAclBody) SetIpVersion(v string) *CreateTCAclBody {
	s.IpVersion = &v
	return s
}

func (s *CreateTCAclBody) SetIpAddress(v string) *CreateTCAclBody {
	s.IpAddress = &v
	return s
}

func (s *CreateTCAclBody) SetName(v string) *CreateTCAclBody {
	s.Name = &v
	return s
}

func (s *CreateTCAclBody) SetDescription(v string) *CreateTCAclBody {
	s.Description = &v
	return s
}

func (s *CreateTCAclBody) SetId(v string) *CreateTCAclBody {
	s.Id = &v
	return s
}

func (s *CreateTCAclBody) SetRegion(v string) *CreateTCAclBody {
	s.Region = &v
	return s
}

func (s *CreateTCAclBody) SetRuleDescription(v string) *CreateTCAclBody {
	s.RuleDescription = &v
	return s
}

func (s *CreateTCAclBody) SetVaz(v string) *CreateTCAclBody {
	s.Vaz = &v
	return s
}


type CreateTCAclBodyBuilder struct {
	s *CreateTCAclBody
}

func NewCreateTCAclBodyBuilder() *CreateTCAclBodyBuilder {
	s := &CreateTCAclBody{}
	b := &CreateTCAclBodyBuilder{s: s}
	return b
}

func (b *CreateTCAclBodyBuilder) IpVersion(v string) *CreateTCAclBodyBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *CreateTCAclBodyBuilder) IpAddress(v string) *CreateTCAclBodyBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *CreateTCAclBodyBuilder) Name(v string) *CreateTCAclBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateTCAclBodyBuilder) Description(v string) *CreateTCAclBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateTCAclBodyBuilder) Id(v string) *CreateTCAclBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *CreateTCAclBodyBuilder) Region(v string) *CreateTCAclBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *CreateTCAclBodyBuilder) RuleDescription(v string) *CreateTCAclBodyBuilder {
	b.s.RuleDescription = &v
	return b
}

func (b *CreateTCAclBodyBuilder) Vaz(v string) *CreateTCAclBodyBuilder {
	b.s.Vaz = &v
	return b
}

func (b *CreateTCAclBodyBuilder) Build() *CreateTCAclBody {
	return b.s
}


