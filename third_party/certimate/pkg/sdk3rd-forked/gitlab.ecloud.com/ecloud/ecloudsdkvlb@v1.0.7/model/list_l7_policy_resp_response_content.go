// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListL7PolicyRespResponseContentCompareTypeEnum string

// List of CompareType
const (
    ListL7PolicyRespResponseContentCompareTypeEnumRegex ListL7PolicyRespResponseContentCompareTypeEnum = "REGEX"
    ListL7PolicyRespResponseContentCompareTypeEnumStartsWith ListL7PolicyRespResponseContentCompareTypeEnum = "STARTS_WITH"
    ListL7PolicyRespResponseContentCompareTypeEnumEqualTo ListL7PolicyRespResponseContentCompareTypeEnum = "EQUAL_TO"
)

type ListL7PolicyRespResponseContent struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 七层转发策略域名
	L7PolicyDomainName *string `json:"l7PolicyDomainName,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 匹配规格，取值如下： REGEX：正则表达式 STARTS_WITH：以某字符串开始 ENDS_WITH：以某字符串结束 CONTAINS：包含 EQUAL_TO：等于
	CompareType *ListL7PolicyRespResponseContentCompareTypeEnum `json:"compareType,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 七层转发策略规则输入的URL
	L7PolicyUrl *string `json:"l7PolicyUrl,omitempty"`
    // 七层转发策略名称
	L7PolicyName *string `json:"l7PolicyName,omitempty"`
    // 后端服务组ID
	PoolId *string `json:"poolId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 七层转发策略ID
	L7PolicyId *string `json:"l7PolicyId,omitempty"`
    // 七层转发策略的优先级的位置
	Position *int32 `json:"position,omitempty"`
    // 转发策略激活状态
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 后端服务组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ListL7PolicyRespResponseContent) String() string {
	return utils.Beautify(s)
}

func (s ListL7PolicyRespResponseContent) GoString() string {
	return s.String()
}

func (s ListL7PolicyRespResponseContent) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListL7PolicyRespResponseContent) SetModifiedTime(v string) *ListL7PolicyRespResponseContent {
	s.ModifiedTime = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetDescription(v string) *ListL7PolicyRespResponseContent {
	s.Description = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetL7PolicyDomainName(v string) *ListL7PolicyRespResponseContent {
	s.L7PolicyDomainName = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetIsMultiAz(v bool) *ListL7PolicyRespResponseContent {
	s.IsMultiAz = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetListenerId(v string) *ListL7PolicyRespResponseContent {
	s.ListenerId = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetCompareType(v ListL7PolicyRespResponseContentCompareTypeEnum) *ListL7PolicyRespResponseContent {
	s.CompareType = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetDeleted(v bool) *ListL7PolicyRespResponseContent {
	s.Deleted = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetL7PolicyUrl(v string) *ListL7PolicyRespResponseContent {
	s.L7PolicyUrl = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetL7PolicyName(v string) *ListL7PolicyRespResponseContent {
	s.L7PolicyName = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetPoolId(v string) *ListL7PolicyRespResponseContent {
	s.PoolId = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetCreatedTime(v string) *ListL7PolicyRespResponseContent {
	s.CreatedTime = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetL7PolicyId(v string) *ListL7PolicyRespResponseContent {
	s.L7PolicyId = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetPosition(v int32) *ListL7PolicyRespResponseContent {
	s.Position = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetAdminStateUp(v bool) *ListL7PolicyRespResponseContent {
	s.AdminStateUp = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetMultiAzUuid(v string) *ListL7PolicyRespResponseContent {
	s.MultiAzUuid = &v
	return s
}

func (s *ListL7PolicyRespResponseContent) SetPoolName(v string) *ListL7PolicyRespResponseContent {
	s.PoolName = &v
	return s
}


type ListL7PolicyRespResponseContentBuilder struct {
	s *ListL7PolicyRespResponseContent
}

func NewListL7PolicyRespResponseContentBuilder() *ListL7PolicyRespResponseContentBuilder {
	s := &ListL7PolicyRespResponseContent{}
	b := &ListL7PolicyRespResponseContentBuilder{s: s}
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) ModifiedTime(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) Description(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.Description = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) L7PolicyDomainName(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.L7PolicyDomainName = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) IsMultiAz(v bool) *ListL7PolicyRespResponseContentBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) ListenerId(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) CompareType(v ListL7PolicyRespResponseContentCompareTypeEnum) *ListL7PolicyRespResponseContentBuilder {
	b.s.CompareType = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) Deleted(v bool) *ListL7PolicyRespResponseContentBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) L7PolicyUrl(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.L7PolicyUrl = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) L7PolicyName(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.L7PolicyName = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) PoolId(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) CreatedTime(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) L7PolicyId(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) Position(v int32) *ListL7PolicyRespResponseContentBuilder {
	b.s.Position = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) AdminStateUp(v bool) *ListL7PolicyRespResponseContentBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) MultiAzUuid(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) PoolName(v string) *ListL7PolicyRespResponseContentBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListL7PolicyRespResponseContentBuilder) Build() *ListL7PolicyRespResponseContent {
	return b.s
}


