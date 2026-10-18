// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetL7PolicyRespResponseBodyCompareTypeEnum string

// List of CompareType
const (
    GetL7PolicyRespResponseBodyCompareTypeEnumRegex GetL7PolicyRespResponseBodyCompareTypeEnum = "REGEX"
    GetL7PolicyRespResponseBodyCompareTypeEnumStartsWith GetL7PolicyRespResponseBodyCompareTypeEnum = "STARTS_WITH"
    GetL7PolicyRespResponseBodyCompareTypeEnumEqualTo GetL7PolicyRespResponseBodyCompareTypeEnum = "EQUAL_TO"
)

type GetL7PolicyRespResponseBody struct {

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
    // 匹配规格，取值如下： REGEX：正则表达式STARTS_WITH：以某字符串开始ENDS_WITH：以某字符串结束CONTAINS：包含 EQUAL_TO：等于
	CompareType *GetL7PolicyRespResponseBodyCompareTypeEnum `json:"compareType,omitempty"`
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

func (s GetL7PolicyRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetL7PolicyRespResponseBody) GoString() string {
	return s.String()
}

func (s GetL7PolicyRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetL7PolicyRespResponseBody) SetModifiedTime(v string) *GetL7PolicyRespResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetDescription(v string) *GetL7PolicyRespResponseBody {
	s.Description = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetL7PolicyDomainName(v string) *GetL7PolicyRespResponseBody {
	s.L7PolicyDomainName = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetIsMultiAz(v bool) *GetL7PolicyRespResponseBody {
	s.IsMultiAz = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetListenerId(v string) *GetL7PolicyRespResponseBody {
	s.ListenerId = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetCompareType(v GetL7PolicyRespResponseBodyCompareTypeEnum) *GetL7PolicyRespResponseBody {
	s.CompareType = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetDeleted(v bool) *GetL7PolicyRespResponseBody {
	s.Deleted = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetL7PolicyUrl(v string) *GetL7PolicyRespResponseBody {
	s.L7PolicyUrl = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetL7PolicyName(v string) *GetL7PolicyRespResponseBody {
	s.L7PolicyName = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetPoolId(v string) *GetL7PolicyRespResponseBody {
	s.PoolId = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetCreatedTime(v string) *GetL7PolicyRespResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetL7PolicyId(v string) *GetL7PolicyRespResponseBody {
	s.L7PolicyId = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetPosition(v int32) *GetL7PolicyRespResponseBody {
	s.Position = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetAdminStateUp(v bool) *GetL7PolicyRespResponseBody {
	s.AdminStateUp = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetMultiAzUuid(v string) *GetL7PolicyRespResponseBody {
	s.MultiAzUuid = &v
	return s
}

func (s *GetL7PolicyRespResponseBody) SetPoolName(v string) *GetL7PolicyRespResponseBody {
	s.PoolName = &v
	return s
}


type GetL7PolicyRespResponseBodyBuilder struct {
	s *GetL7PolicyRespResponseBody
}

func NewGetL7PolicyRespResponseBodyBuilder() *GetL7PolicyRespResponseBodyBuilder {
	s := &GetL7PolicyRespResponseBody{}
	b := &GetL7PolicyRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) ModifiedTime(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) Description(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) L7PolicyDomainName(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.L7PolicyDomainName = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) IsMultiAz(v bool) *GetL7PolicyRespResponseBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) ListenerId(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) CompareType(v GetL7PolicyRespResponseBodyCompareTypeEnum) *GetL7PolicyRespResponseBodyBuilder {
	b.s.CompareType = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) Deleted(v bool) *GetL7PolicyRespResponseBodyBuilder {
	b.s.Deleted = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) L7PolicyUrl(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.L7PolicyUrl = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) L7PolicyName(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.L7PolicyName = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) PoolId(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) CreatedTime(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) L7PolicyId(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) Position(v int32) *GetL7PolicyRespResponseBodyBuilder {
	b.s.Position = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) AdminStateUp(v bool) *GetL7PolicyRespResponseBodyBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) MultiAzUuid(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) PoolName(v string) *GetL7PolicyRespResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetL7PolicyRespResponseBodyBuilder) Build() *GetL7PolicyRespResponseBody {
	return b.s
}


