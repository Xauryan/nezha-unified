// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAliAclBody struct {
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

func (s CreateAliAclBody) String() string {
	return utils.Beautify(s)
}

func (s CreateAliAclBody) GoString() string {
	return s.String()
}

func (s CreateAliAclBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAliAclBody) SetIpVersion(v string) *CreateAliAclBody {
	s.IpVersion = &v
	return s
}

func (s *CreateAliAclBody) SetIpAddress(v string) *CreateAliAclBody {
	s.IpAddress = &v
	return s
}

func (s *CreateAliAclBody) SetName(v string) *CreateAliAclBody {
	s.Name = &v
	return s
}

func (s *CreateAliAclBody) SetDescription(v string) *CreateAliAclBody {
	s.Description = &v
	return s
}

func (s *CreateAliAclBody) SetId(v string) *CreateAliAclBody {
	s.Id = &v
	return s
}

func (s *CreateAliAclBody) SetRegion(v string) *CreateAliAclBody {
	s.Region = &v
	return s
}

func (s *CreateAliAclBody) SetRuleDescription(v string) *CreateAliAclBody {
	s.RuleDescription = &v
	return s
}

func (s *CreateAliAclBody) SetVaz(v string) *CreateAliAclBody {
	s.Vaz = &v
	return s
}


type CreateAliAclBodyBuilder struct {
	s *CreateAliAclBody
}

func NewCreateAliAclBodyBuilder() *CreateAliAclBodyBuilder {
	s := &CreateAliAclBody{}
	b := &CreateAliAclBodyBuilder{s: s}
	return b
}

func (b *CreateAliAclBodyBuilder) IpVersion(v string) *CreateAliAclBodyBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *CreateAliAclBodyBuilder) IpAddress(v string) *CreateAliAclBodyBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *CreateAliAclBodyBuilder) Name(v string) *CreateAliAclBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *CreateAliAclBodyBuilder) Description(v string) *CreateAliAclBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateAliAclBodyBuilder) Id(v string) *CreateAliAclBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *CreateAliAclBodyBuilder) Region(v string) *CreateAliAclBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *CreateAliAclBodyBuilder) RuleDescription(v string) *CreateAliAclBodyBuilder {
	b.s.RuleDescription = &v
	return b
}

func (b *CreateAliAclBodyBuilder) Vaz(v string) *CreateAliAclBodyBuilder {
	b.s.Vaz = &v
	return b
}

func (b *CreateAliAclBodyBuilder) Build() *CreateAliAclBody {
	return b.s
}


