// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPoolRespResponseL7PolicyRespsCompareTypeEnum string

// List of CompareType
const (
    GetPoolRespResponseL7PolicyRespsCompareTypeEnumRegex GetPoolRespResponseL7PolicyRespsCompareTypeEnum = "REGEX"
    GetPoolRespResponseL7PolicyRespsCompareTypeEnumStartsWith GetPoolRespResponseL7PolicyRespsCompareTypeEnum = "STARTS_WITH"
    GetPoolRespResponseL7PolicyRespsCompareTypeEnumEndsWith GetPoolRespResponseL7PolicyRespsCompareTypeEnum = "ENDS_WITH"
    GetPoolRespResponseL7PolicyRespsCompareTypeEnumContains GetPoolRespResponseL7PolicyRespsCompareTypeEnum = "CONTAINS"
    GetPoolRespResponseL7PolicyRespsCompareTypeEnumEqualTo GetPoolRespResponseL7PolicyRespsCompareTypeEnum = "EQUAL_TO"
)
type GetPoolRespResponseL7PolicyRespsRuleTypeEnum string

// List of RuleType
const (
    GetPoolRespResponseL7PolicyRespsRuleTypeEnumHostName GetPoolRespResponseL7PolicyRespsRuleTypeEnum = "HOST_NAME"
    GetPoolRespResponseL7PolicyRespsRuleTypeEnumPath GetPoolRespResponseL7PolicyRespsRuleTypeEnum = "PATH"
    GetPoolRespResponseL7PolicyRespsRuleTypeEnumFileType GetPoolRespResponseL7PolicyRespsRuleTypeEnum = "FILE_TYPE"
    GetPoolRespResponseL7PolicyRespsRuleTypeEnumHeader GetPoolRespResponseL7PolicyRespsRuleTypeEnum = "HEADER"
    GetPoolRespResponseL7PolicyRespsRuleTypeEnumCookie GetPoolRespResponseL7PolicyRespsRuleTypeEnum = "COOKIE"
)

type GetPoolRespResponseL7PolicyResps struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 域名
	L7PolicyDomainName *string `json:"l7PolicyDomainName,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
    // 七层转发策略规则值，已弃用
	L7RuleValue *string `json:"l7RuleValue,omitempty"`
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // 匹配规格：REGEX（正则表达式匹配），STARTS_WITH（开头匹配），ENDS_WITH（结尾匹配），CONTAINS（包含匹配），EQUAL_TO（等于匹配）
	CompareType *GetPoolRespResponseL7PolicyRespsCompareTypeEnum `json:"compareType,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 七层转发策略URL
	L7PolicyUrl *string `json:"l7PolicyUrl,omitempty"`
    // 七层转发策略名称
	L7PolicyName *string `json:"l7PolicyName,omitempty"`
    // 七层转发策略规则类型，已弃用
	RuleType *GetPoolRespResponseL7PolicyRespsRuleTypeEnum `json:"ruleType,omitempty"`
    // 后端服务组ID
	PoolId *string `json:"poolId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 七层转发策略ID
	L7PolicyId *string `json:"l7PolicyId,omitempty"`
    // 七层转发策略的位置
	Position *int32 `json:"position,omitempty"`
    // 激活状态
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 多AZ场景下创建请求的requestId唯一标志为uuid
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 后端服务组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s GetPoolRespResponseL7PolicyResps) String() string {
	return utils.Beautify(s)
}

func (s GetPoolRespResponseL7PolicyResps) GoString() string {
	return s.String()
}

func (s GetPoolRespResponseL7PolicyResps) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolRespResponseL7PolicyResps) SetModifiedTime(v string) *GetPoolRespResponseL7PolicyResps {
	s.ModifiedTime = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetDescription(v string) *GetPoolRespResponseL7PolicyResps {
	s.Description = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetL7PolicyDomainName(v string) *GetPoolRespResponseL7PolicyResps {
	s.L7PolicyDomainName = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetIsMultiAz(v bool) *GetPoolRespResponseL7PolicyResps {
	s.IsMultiAz = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetL7RuleValue(v string) *GetPoolRespResponseL7PolicyResps {
	s.L7RuleValue = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetListenerId(v string) *GetPoolRespResponseL7PolicyResps {
	s.ListenerId = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetCompareType(v GetPoolRespResponseL7PolicyRespsCompareTypeEnum) *GetPoolRespResponseL7PolicyResps {
	s.CompareType = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetDeleted(v bool) *GetPoolRespResponseL7PolicyResps {
	s.Deleted = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetL7PolicyUrl(v string) *GetPoolRespResponseL7PolicyResps {
	s.L7PolicyUrl = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetL7PolicyName(v string) *GetPoolRespResponseL7PolicyResps {
	s.L7PolicyName = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetRuleType(v GetPoolRespResponseL7PolicyRespsRuleTypeEnum) *GetPoolRespResponseL7PolicyResps {
	s.RuleType = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetPoolId(v string) *GetPoolRespResponseL7PolicyResps {
	s.PoolId = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetCreatedTime(v string) *GetPoolRespResponseL7PolicyResps {
	s.CreatedTime = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetL7PolicyId(v string) *GetPoolRespResponseL7PolicyResps {
	s.L7PolicyId = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetPosition(v int32) *GetPoolRespResponseL7PolicyResps {
	s.Position = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetAdminStateUp(v bool) *GetPoolRespResponseL7PolicyResps {
	s.AdminStateUp = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetMultiAzUuid(v string) *GetPoolRespResponseL7PolicyResps {
	s.MultiAzUuid = &v
	return s
}

func (s *GetPoolRespResponseL7PolicyResps) SetPoolName(v string) *GetPoolRespResponseL7PolicyResps {
	s.PoolName = &v
	return s
}


type GetPoolRespResponseL7PolicyRespsBuilder struct {
	s *GetPoolRespResponseL7PolicyResps
}

func NewGetPoolRespResponseL7PolicyRespsBuilder() *GetPoolRespResponseL7PolicyRespsBuilder {
	s := &GetPoolRespResponseL7PolicyResps{}
	b := &GetPoolRespResponseL7PolicyRespsBuilder{s: s}
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) ModifiedTime(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) Description(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.Description = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) L7PolicyDomainName(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyDomainName = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) IsMultiAz(v bool) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) L7RuleValue(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7RuleValue = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) ListenerId(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) CompareType(v GetPoolRespResponseL7PolicyRespsCompareTypeEnum) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.CompareType = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) Deleted(v bool) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.Deleted = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) L7PolicyUrl(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyUrl = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) L7PolicyName(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyName = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) RuleType(v GetPoolRespResponseL7PolicyRespsRuleTypeEnum) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.RuleType = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) PoolId(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) CreatedTime(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) L7PolicyId(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) Position(v int32) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.Position = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) AdminStateUp(v bool) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) MultiAzUuid(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) PoolName(v string) *GetPoolRespResponseL7PolicyRespsBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetPoolRespResponseL7PolicyRespsBuilder) Build() *GetPoolRespResponseL7PolicyResps {
	return b.s
}


