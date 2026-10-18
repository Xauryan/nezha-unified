// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupRuleResponseContent struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // IP版本（IPv4或IPv6）
	IpVersion *string `json:"ipVersion,omitempty"`
    // 是否删除
	IsDelete *bool `json:"isDelete,omitempty"`
    // 用户名
	Proposer *string `json:"proposer,omitempty"`
    // 客户ID
	CustomerId *string `json:"customerId,omitempty"`
    // IP地址
	IpAddress *string `json:"ipAddress,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 访问控制组规则描述
	Description *string `json:"description,omitempty"`
    // 修改人
	ModifiedBy *string `json:"modifiedBy,omitempty"`
    // 访问控制组规则ID
	Id *string `json:"id,omitempty"`
    // 控制组ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
}

func (s ListControlGroupRuleResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupRuleResponseContent) GoString() string {
	return s.String()
}

func (s ListControlGroupRuleResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupRuleResponseContent) SetModifiedTime(v string) *ListControlGroupRuleResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetIpVersion(v string) *ListControlGroupRuleResponseContent {
	s.IpVersion = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetIsDelete(v bool) *ListControlGroupRuleResponseContent {
	s.IsDelete = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetProposer(v string) *ListControlGroupRuleResponseContent {
	s.Proposer = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetCustomerId(v string) *ListControlGroupRuleResponseContent {
	s.CustomerId = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetIpAddress(v string) *ListControlGroupRuleResponseContent {
	s.IpAddress = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetCreatedTime(v string) *ListControlGroupRuleResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetDescription(v string) *ListControlGroupRuleResponseContent {
	s.Description = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetModifiedBy(v string) *ListControlGroupRuleResponseContent {
	s.ModifiedBy = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetId(v string) *ListControlGroupRuleResponseContent {
	s.Id = &v
	return s
}

func (s *ListControlGroupRuleResponseContent) SetControlGroupId(v string) *ListControlGroupRuleResponseContent {
	s.ControlGroupId = &v
	return s
}


type ListControlGroupRuleResponseContentBuilder struct {
	s *ListControlGroupRuleResponseContent
}

func NewListControlGroupRuleResponseContentBuilder() *ListControlGroupRuleResponseContentBuilder {
	s := &ListControlGroupRuleResponseContent{}
	b := &ListControlGroupRuleResponseContentBuilder{s: s}
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) ModifiedTime(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) IpVersion(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.IpVersion = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) IsDelete(v bool) *ListControlGroupRuleResponseContentBuilder {
	b.s.IsDelete = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) Proposer(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.Proposer = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) CustomerId(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.CustomerId = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) IpAddress(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.IpAddress = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) CreatedTime(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) Description(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) ModifiedBy(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.ModifiedBy = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) Id(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.Id = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) ControlGroupId(v string) *ListControlGroupRuleResponseContentBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *ListControlGroupRuleResponseContentBuilder) Build() *ListControlGroupRuleResponseContent {
	return b.s
}


