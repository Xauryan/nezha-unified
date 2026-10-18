// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum string

// List of HealthType
const (
    AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnumHttp10 AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum = "HTTP10"
    AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnumHttp11 AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum = "HTTP11"
    AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnumHttps AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum = "HTTPS"
    AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnumTcp AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum = "TCP"
    AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnumPing AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum = "PING"
)

type AsyncCreatePolicyRequestHealthMonitorCreateReq struct {

    // 监听器的检查间隔（单位：秒）
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 监听器的尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器绑定的检查方式为HTTP的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 后端服务器组ID，若为Null，则新建后端服务器组，poolCreateReq则必填
	PoolId *string `json:"poolId,omitempty"`
    // 监听器健康检查方式
	HealthType *AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum `json:"healthType,omitempty"`
    // 用于监控的URL
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 监听器的超时时间（单位：秒）
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
}

func (s AsyncCreatePolicyRequestHealthMonitorCreateReq) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePolicyRequestHealthMonitorCreateReq) GoString() string {
	return s.String()
}

func (s AsyncCreatePolicyRequestHealthMonitorCreateReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetHealthDelay(v int32) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.HealthDelay = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetHealthExpectedCode(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.HealthExpectedCode = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetHealthMaxRetries(v int32) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.HealthMaxRetries = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetHealthHttpMethod(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.HealthHttpMethod = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetPoolId(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.PoolId = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetHealthType(v AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.HealthType = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetHealthUrlPath(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.HealthUrlPath = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetAdminStateUp(v bool) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.AdminStateUp = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetHealthTimeout(v int32) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.HealthTimeout = &v
	return s
}

func (s *AsyncCreatePolicyRequestHealthMonitorCreateReq) SetIsMultiAz(v bool) *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	s.IsMultiAz = &v
	return s
}


type AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder struct {
	s *AsyncCreatePolicyRequestHealthMonitorCreateReq
}

func NewAsyncCreatePolicyRequestHealthMonitorCreateReqBuilder() *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	s := &AsyncCreatePolicyRequestHealthMonitorCreateReq{}
	b := &AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder{s: s}
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) HealthDelay(v int32) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) HealthExpectedCode(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) HealthMaxRetries(v int32) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) HealthHttpMethod(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) PoolId(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.PoolId = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) HealthType(v AsyncCreatePolicyRequestHealthMonitorCreateReqHealthTypeEnum) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.HealthType = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) HealthUrlPath(v string) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) AdminStateUp(v bool) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) HealthTimeout(v int32) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) IsMultiAz(v bool) *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *AsyncCreatePolicyRequestHealthMonitorCreateReqBuilder) Build() *AsyncCreatePolicyRequestHealthMonitorCreateReq {
	return b.s
}


