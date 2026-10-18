// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListPoolRespResponseL7PolicyRespsCompareTypeEnum string

// List of CompareType
const (
    ListPoolRespResponseL7PolicyRespsCompareTypeEnumRegex ListPoolRespResponseL7PolicyRespsCompareTypeEnum = "REGEX"
    ListPoolRespResponseL7PolicyRespsCompareTypeEnumStartsWith ListPoolRespResponseL7PolicyRespsCompareTypeEnum = "STARTS_WITH"
    ListPoolRespResponseL7PolicyRespsCompareTypeEnumEndsWith ListPoolRespResponseL7PolicyRespsCompareTypeEnum = "ENDS_WITH"
    ListPoolRespResponseL7PolicyRespsCompareTypeEnumContains ListPoolRespResponseL7PolicyRespsCompareTypeEnum = "CONTAINS"
    ListPoolRespResponseL7PolicyRespsCompareTypeEnumEqualTo ListPoolRespResponseL7PolicyRespsCompareTypeEnum = "EQUAL_TO"
)
type ListPoolRespResponseL7PolicyRespsRuleTypeEnum string

// List of RuleType
const (
    ListPoolRespResponseL7PolicyRespsRuleTypeEnumHostName ListPoolRespResponseL7PolicyRespsRuleTypeEnum = "HOST_NAME"
    ListPoolRespResponseL7PolicyRespsRuleTypeEnumPath ListPoolRespResponseL7PolicyRespsRuleTypeEnum = "PATH"
    ListPoolRespResponseL7PolicyRespsRuleTypeEnumFileType ListPoolRespResponseL7PolicyRespsRuleTypeEnum = "FILE_TYPE"
    ListPoolRespResponseL7PolicyRespsRuleTypeEnumHeader ListPoolRespResponseL7PolicyRespsRuleTypeEnum = "HEADER"
    ListPoolRespResponseL7PolicyRespsRuleTypeEnumCookie ListPoolRespResponseL7PolicyRespsRuleTypeEnum = "COOKIE"
)

type ListPoolRespResponseL7PolicyResps struct {

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
    // 匹配规格：REGEX（正则表达式匹配），STARTS（开头匹配）_WITH，ENDS_WITH（结尾匹配），CONTAINS（包含匹配），EQUAL_TO（等于匹配）
	CompareType *ListPoolRespResponseL7PolicyRespsCompareTypeEnum `json:"compareType,omitempty"`
    // 是否删除
	Deleted *bool `json:"deleted,omitempty"`
    // 七层转发策略URL
	L7PolicyUrl *string `json:"l7PolicyUrl,omitempty"`
    // 七层转发策略名称
	L7PolicyName *string `json:"l7PolicyName,omitempty"`
    // 七层转发策略规则类型，已弃用
	RuleType *ListPoolRespResponseL7PolicyRespsRuleTypeEnum `json:"ruleType,omitempty"`
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
    // 双AZ场景下创建请求的requestId唯一标志为UUID
	MultiAzUuid *string `json:"multiAzUuid,omitempty"`
    // 后端服务组名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ListPoolRespResponseL7PolicyResps) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespResponseL7PolicyResps) GoString() string {
	return s.String()
}

func (s ListPoolRespResponseL7PolicyResps) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespResponseL7PolicyResps) SetModifiedTime(v string) *ListPoolRespResponseL7PolicyResps {
	s.ModifiedTime = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetDescription(v string) *ListPoolRespResponseL7PolicyResps {
	s.Description = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetL7PolicyDomainName(v string) *ListPoolRespResponseL7PolicyResps {
	s.L7PolicyDomainName = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetIsMultiAz(v bool) *ListPoolRespResponseL7PolicyResps {
	s.IsMultiAz = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetL7RuleValue(v string) *ListPoolRespResponseL7PolicyResps {
	s.L7RuleValue = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetListenerId(v string) *ListPoolRespResponseL7PolicyResps {
	s.ListenerId = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetCompareType(v ListPoolRespResponseL7PolicyRespsCompareTypeEnum) *ListPoolRespResponseL7PolicyResps {
	s.CompareType = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetDeleted(v bool) *ListPoolRespResponseL7PolicyResps {
	s.Deleted = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetL7PolicyUrl(v string) *ListPoolRespResponseL7PolicyResps {
	s.L7PolicyUrl = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetL7PolicyName(v string) *ListPoolRespResponseL7PolicyResps {
	s.L7PolicyName = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetRuleType(v ListPoolRespResponseL7PolicyRespsRuleTypeEnum) *ListPoolRespResponseL7PolicyResps {
	s.RuleType = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetPoolId(v string) *ListPoolRespResponseL7PolicyResps {
	s.PoolId = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetCreatedTime(v string) *ListPoolRespResponseL7PolicyResps {
	s.CreatedTime = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetL7PolicyId(v string) *ListPoolRespResponseL7PolicyResps {
	s.L7PolicyId = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetPosition(v int32) *ListPoolRespResponseL7PolicyResps {
	s.Position = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetAdminStateUp(v bool) *ListPoolRespResponseL7PolicyResps {
	s.AdminStateUp = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetMultiAzUuid(v string) *ListPoolRespResponseL7PolicyResps {
	s.MultiAzUuid = &v
	return s
}

func (s *ListPoolRespResponseL7PolicyResps) SetPoolName(v string) *ListPoolRespResponseL7PolicyResps {
	s.PoolName = &v
	return s
}


type ListPoolRespResponseL7PolicyRespsBuilder struct {
	s *ListPoolRespResponseL7PolicyResps
}

func NewListPoolRespResponseL7PolicyRespsBuilder() *ListPoolRespResponseL7PolicyRespsBuilder {
	s := &ListPoolRespResponseL7PolicyResps{}
	b := &ListPoolRespResponseL7PolicyRespsBuilder{s: s}
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) ModifiedTime(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) Description(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.Description = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) L7PolicyDomainName(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyDomainName = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) IsMultiAz(v bool) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) L7RuleValue(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7RuleValue = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) ListenerId(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) CompareType(v ListPoolRespResponseL7PolicyRespsCompareTypeEnum) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.CompareType = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) Deleted(v bool) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.Deleted = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) L7PolicyUrl(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyUrl = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) L7PolicyName(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyName = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) RuleType(v ListPoolRespResponseL7PolicyRespsRuleTypeEnum) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.RuleType = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) PoolId(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) CreatedTime(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) L7PolicyId(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) Position(v int32) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.Position = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) AdminStateUp(v bool) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) MultiAzUuid(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.MultiAzUuid = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) PoolName(v string) *ListPoolRespResponseL7PolicyRespsBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListPoolRespResponseL7PolicyRespsBuilder) Build() *ListPoolRespResponseL7PolicyResps {
	return b.s
}


