// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreatePolicyBodyCompareTypeEnum string

// List of CompareType
const (
    AsyncCreatePolicyBodyCompareTypeEnumRegex AsyncCreatePolicyBodyCompareTypeEnum = "REGEX"
    AsyncCreatePolicyBodyCompareTypeEnumStartsWith AsyncCreatePolicyBodyCompareTypeEnum = "STARTS_WITH"
    AsyncCreatePolicyBodyCompareTypeEnumEqualTo AsyncCreatePolicyBodyCompareTypeEnum = "EQUAL_TO"
)
type AsyncCreatePolicyBodyHostNameCompareTypeEnum string

// List of HostNameCompareType
const (
    AsyncCreatePolicyBodyHostNameCompareTypeEnumRegex AsyncCreatePolicyBodyHostNameCompareTypeEnum = "REGEX"
    AsyncCreatePolicyBodyHostNameCompareTypeEnumStartsWith AsyncCreatePolicyBodyHostNameCompareTypeEnum = "STARTS_WITH"
    AsyncCreatePolicyBodyHostNameCompareTypeEnumEqualTo AsyncCreatePolicyBodyHostNameCompareTypeEnum = "EQUAL_TO"
)

type AsyncCreatePolicyBody struct {
    position.Body
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
    // URL的匹配规格，取值如下:  REGEX:正则匹配; STARTS_WITH:前缀匹配; EQUAL_TO:精准匹配 选择URL时必填
	CompareType *AsyncCreatePolicyBodyCompareTypeEnum `json:"compareType,omitempty"`
    // 七层转发策略URL，与转发策略域名需填写其中一个，也可都填写
	L7PolicyUrl *string `json:"l7PolicyUrl,omitempty"`
    // 域名的匹配规格，取值如下：REGEX：正则匹配；STARTS_WITH：前缀匹配； EQUAL_TO：精准匹配，选择域名时必填
	HostNameCompareType *AsyncCreatePolicyBodyHostNameCompareTypeEnum `json:"hostNameCompareType,omitempty"`
    // 七层转发策略名称，为5~64位的英文、数字、下划线中划线等的组合
	L7PolicyName *string `json:"l7PolicyName,omitempty"`
    // 后端服务组ID, 若为Null则新建后端服务器组
	PoolId *string `json:"poolId,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 七层转发策略域名，与转发策略URL需填写其中一个，也可都填写
	L7PolicyDomainName *string `json:"l7PolicyDomainName,omitempty"`
    // 七层转发策略在监听器中位置优先级
	Position *int32 `json:"position,omitempty"`
    // 后端服务组创建参数，若为Null则使用已有后端服务器组，poolId则必填
	PoolCreateReq *AsyncCreatePolicyRequestPoolCreateReq `json:"poolCreateReq,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
}

func (s AsyncCreatePolicyBody) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePolicyBody) GoString() string {
	return s.String()
}

func (s AsyncCreatePolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePolicyBody) SetListenerId(v string) *AsyncCreatePolicyBody {
	s.ListenerId = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetCompareType(v AsyncCreatePolicyBodyCompareTypeEnum) *AsyncCreatePolicyBody {
	s.CompareType = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetL7PolicyUrl(v string) *AsyncCreatePolicyBody {
	s.L7PolicyUrl = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetHostNameCompareType(v AsyncCreatePolicyBodyHostNameCompareTypeEnum) *AsyncCreatePolicyBody {
	s.HostNameCompareType = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetL7PolicyName(v string) *AsyncCreatePolicyBody {
	s.L7PolicyName = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetPoolId(v string) *AsyncCreatePolicyBody {
	s.PoolId = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetDescription(v string) *AsyncCreatePolicyBody {
	s.Description = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetL7PolicyDomainName(v string) *AsyncCreatePolicyBody {
	s.L7PolicyDomainName = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetPosition(v int32) *AsyncCreatePolicyBody {
	s.Position = &v
	return s
}

func (s *AsyncCreatePolicyBody) SetPoolCreateReq(v *AsyncCreatePolicyRequestPoolCreateReq) *AsyncCreatePolicyBody {
	s.PoolCreateReq = v
	return s
}

func (s *AsyncCreatePolicyBody) SetIsMultiAz(v bool) *AsyncCreatePolicyBody {
	s.IsMultiAz = &v
	return s
}


type AsyncCreatePolicyBodyBuilder struct {
	s *AsyncCreatePolicyBody
}

func NewAsyncCreatePolicyBodyBuilder() *AsyncCreatePolicyBodyBuilder {
	s := &AsyncCreatePolicyBody{}
	b := &AsyncCreatePolicyBodyBuilder{s: s}
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) ListenerId(v string) *AsyncCreatePolicyBodyBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) CompareType(v AsyncCreatePolicyBodyCompareTypeEnum) *AsyncCreatePolicyBodyBuilder {
	b.s.CompareType = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) L7PolicyUrl(v string) *AsyncCreatePolicyBodyBuilder {
	b.s.L7PolicyUrl = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) HostNameCompareType(v AsyncCreatePolicyBodyHostNameCompareTypeEnum) *AsyncCreatePolicyBodyBuilder {
	b.s.HostNameCompareType = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) L7PolicyName(v string) *AsyncCreatePolicyBodyBuilder {
	b.s.L7PolicyName = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) PoolId(v string) *AsyncCreatePolicyBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) Description(v string) *AsyncCreatePolicyBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) L7PolicyDomainName(v string) *AsyncCreatePolicyBodyBuilder {
	b.s.L7PolicyDomainName = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) Position(v int32) *AsyncCreatePolicyBodyBuilder {
	b.s.Position = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) PoolCreateReq(v *AsyncCreatePolicyRequestPoolCreateReq) *AsyncCreatePolicyBodyBuilder {
	b.s.PoolCreateReq = v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) IsMultiAz(v bool) *AsyncCreatePolicyBodyBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *AsyncCreatePolicyBodyBuilder) Build() *AsyncCreatePolicyBody {
	return b.s
}


