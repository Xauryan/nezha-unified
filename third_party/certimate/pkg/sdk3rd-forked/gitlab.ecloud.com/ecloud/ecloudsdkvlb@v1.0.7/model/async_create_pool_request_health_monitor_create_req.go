// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum string

// List of HealthType
const (
    AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnumHttp AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum = "HTTP"
    AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnumHttp10 AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum = "HTTP10"
    AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnumHttp11 AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum = "HTTP11"
    AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnumHttps AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum = "HTTPS"
    AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnumTcp AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum = "TCP"
    AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnumPing AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum = "PING"
    AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnumUdp AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum = "UDP"
)

type AsyncCreatePoolRequestHealthMonitorCreateReq struct {

    // 监听器的检查间隔（单位：秒）
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 期望HTTP状态代码，取值范围100~599
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 监听器的最大尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器绑定的检查方式为HTTP的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 监听器健康检查方式
	HealthType *AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum `json:"healthType,omitempty"`
    // 用于监控的URL
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 监听器的超时时间（单位：秒）
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
    // 是否是多AZ场景
	IsMultiAz *bool `json:"isMultiAz,omitempty"`
}

func (s AsyncCreatePoolRequestHealthMonitorCreateReq) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePoolRequestHealthMonitorCreateReq) GoString() string {
	return s.String()
}

func (s AsyncCreatePoolRequestHealthMonitorCreateReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetHealthDelay(v int32) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.HealthDelay = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetHealthExpectedCode(v string) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.HealthExpectedCode = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetHealthMaxRetries(v int32) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.HealthMaxRetries = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetHealthHttpMethod(v string) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.HealthHttpMethod = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetPoolId(v string) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.PoolId = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetHealthType(v AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.HealthType = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetHealthUrlPath(v string) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.HealthUrlPath = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetAdminStateUp(v bool) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.AdminStateUp = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetHealthTimeout(v int32) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.HealthTimeout = &v
	return s
}

func (s *AsyncCreatePoolRequestHealthMonitorCreateReq) SetIsMultiAz(v bool) *AsyncCreatePoolRequestHealthMonitorCreateReq {
	s.IsMultiAz = &v
	return s
}


type AsyncCreatePoolRequestHealthMonitorCreateReqBuilder struct {
	s *AsyncCreatePoolRequestHealthMonitorCreateReq
}

func NewAsyncCreatePoolRequestHealthMonitorCreateReqBuilder() *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	s := &AsyncCreatePoolRequestHealthMonitorCreateReq{}
	b := &AsyncCreatePoolRequestHealthMonitorCreateReqBuilder{s: s}
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) HealthDelay(v int32) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) HealthExpectedCode(v string) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) HealthMaxRetries(v int32) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) HealthHttpMethod(v string) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) PoolId(v string) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.PoolId = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) HealthType(v AsyncCreatePoolRequestHealthMonitorCreateReqHealthTypeEnum) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.HealthType = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) HealthUrlPath(v string) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) AdminStateUp(v bool) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) HealthTimeout(v int32) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) IsMultiAz(v bool) *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder {
	b.s.IsMultiAz = &v
	return b
}

func (b *AsyncCreatePoolRequestHealthMonitorCreateReqBuilder) Build() *AsyncCreatePoolRequestHealthMonitorCreateReq {
	return b.s
}


