// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncUpdateRequestHealthMonitorUpdateReq struct {

    // 监听器的检查间隔（单位：秒）
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 期望HTTP状态代码，取值范围100~599
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 监听器的最大尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器绑定的检查方式为http的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 用于监控的URL
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 监听器的超时时间（单位：秒）
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
}

func (s AsyncUpdateRequestHealthMonitorUpdateReq) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateRequestHealthMonitorUpdateReq) GoString() string {
	return s.String()
}

func (s AsyncUpdateRequestHealthMonitorUpdateReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateRequestHealthMonitorUpdateReq) SetHealthDelay(v int32) *AsyncUpdateRequestHealthMonitorUpdateReq {
	s.HealthDelay = &v
	return s
}

func (s *AsyncUpdateRequestHealthMonitorUpdateReq) SetHealthExpectedCode(v string) *AsyncUpdateRequestHealthMonitorUpdateReq {
	s.HealthExpectedCode = &v
	return s
}

func (s *AsyncUpdateRequestHealthMonitorUpdateReq) SetHealthMaxRetries(v int32) *AsyncUpdateRequestHealthMonitorUpdateReq {
	s.HealthMaxRetries = &v
	return s
}

func (s *AsyncUpdateRequestHealthMonitorUpdateReq) SetHealthHttpMethod(v string) *AsyncUpdateRequestHealthMonitorUpdateReq {
	s.HealthHttpMethod = &v
	return s
}

func (s *AsyncUpdateRequestHealthMonitorUpdateReq) SetHealthUrlPath(v string) *AsyncUpdateRequestHealthMonitorUpdateReq {
	s.HealthUrlPath = &v
	return s
}

func (s *AsyncUpdateRequestHealthMonitorUpdateReq) SetHealthTimeout(v int32) *AsyncUpdateRequestHealthMonitorUpdateReq {
	s.HealthTimeout = &v
	return s
}


type AsyncUpdateRequestHealthMonitorUpdateReqBuilder struct {
	s *AsyncUpdateRequestHealthMonitorUpdateReq
}

func NewAsyncUpdateRequestHealthMonitorUpdateReqBuilder() *AsyncUpdateRequestHealthMonitorUpdateReqBuilder {
	s := &AsyncUpdateRequestHealthMonitorUpdateReq{}
	b := &AsyncUpdateRequestHealthMonitorUpdateReqBuilder{s: s}
	return b
}

func (b *AsyncUpdateRequestHealthMonitorUpdateReqBuilder) HealthDelay(v int32) *AsyncUpdateRequestHealthMonitorUpdateReqBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *AsyncUpdateRequestHealthMonitorUpdateReqBuilder) HealthExpectedCode(v string) *AsyncUpdateRequestHealthMonitorUpdateReqBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *AsyncUpdateRequestHealthMonitorUpdateReqBuilder) HealthMaxRetries(v int32) *AsyncUpdateRequestHealthMonitorUpdateReqBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *AsyncUpdateRequestHealthMonitorUpdateReqBuilder) HealthHttpMethod(v string) *AsyncUpdateRequestHealthMonitorUpdateReqBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *AsyncUpdateRequestHealthMonitorUpdateReqBuilder) HealthUrlPath(v string) *AsyncUpdateRequestHealthMonitorUpdateReqBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *AsyncUpdateRequestHealthMonitorUpdateReqBuilder) HealthTimeout(v int32) *AsyncUpdateRequestHealthMonitorUpdateReqBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *AsyncUpdateRequestHealthMonitorUpdateReqBuilder) Build() *AsyncUpdateRequestHealthMonitorUpdateReq {
	return b.s
}


