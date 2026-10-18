// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListPoolRespResponseHealthMonitorRespHealthTypeEnum string

// List of HealthType
const (
    ListPoolRespResponseHealthMonitorRespHealthTypeEnumHttp ListPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTP"
    ListPoolRespResponseHealthMonitorRespHealthTypeEnumHttps ListPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTPS"
    ListPoolRespResponseHealthMonitorRespHealthTypeEnumTcp ListPoolRespResponseHealthMonitorRespHealthTypeEnum = "TCP"
    ListPoolRespResponseHealthMonitorRespHealthTypeEnumPing ListPoolRespResponseHealthMonitorRespHealthTypeEnum = "PING"
    ListPoolRespResponseHealthMonitorRespHealthTypeEnumUdp ListPoolRespResponseHealthMonitorRespHealthTypeEnum = "UDP"
    ListPoolRespResponseHealthMonitorRespHealthTypeEnumHttp10 ListPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTP10"
    ListPoolRespResponseHealthMonitorRespHealthTypeEnumHttp11 ListPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTP11"
)

type ListPoolRespResponseHealthMonitorResp struct {

    // 健康监控时间间隔（单位：秒）
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 最大尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器绑定的检查方式为HTTP的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 健康检查的资源ID
	HealthId *string `json:"healthId,omitempty"`
    // 健康监控类型
	HealthType *ListPoolRespResponseHealthMonitorRespHealthTypeEnum `json:"healthType,omitempty"`
    // 用于监控的URL，必须以“/”打头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 最大等待时间（单位：秒）
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
}

func (s ListPoolRespResponseHealthMonitorResp) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespResponseHealthMonitorResp) GoString() string {
	return s.String()
}

func (s ListPoolRespResponseHealthMonitorResp) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthDelay(v int32) *ListPoolRespResponseHealthMonitorResp {
	s.HealthDelay = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthExpectedCode(v string) *ListPoolRespResponseHealthMonitorResp {
	s.HealthExpectedCode = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthMaxRetries(v int32) *ListPoolRespResponseHealthMonitorResp {
	s.HealthMaxRetries = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthHttpMethod(v string) *ListPoolRespResponseHealthMonitorResp {
	s.HealthHttpMethod = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthId(v string) *ListPoolRespResponseHealthMonitorResp {
	s.HealthId = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthType(v ListPoolRespResponseHealthMonitorRespHealthTypeEnum) *ListPoolRespResponseHealthMonitorResp {
	s.HealthType = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthUrlPath(v string) *ListPoolRespResponseHealthMonitorResp {
	s.HealthUrlPath = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetAdminStateUp(v bool) *ListPoolRespResponseHealthMonitorResp {
	s.AdminStateUp = &v
	return s
}

func (s *ListPoolRespResponseHealthMonitorResp) SetHealthTimeout(v int32) *ListPoolRespResponseHealthMonitorResp {
	s.HealthTimeout = &v
	return s
}


type ListPoolRespResponseHealthMonitorRespBuilder struct {
	s *ListPoolRespResponseHealthMonitorResp
}

func NewListPoolRespResponseHealthMonitorRespBuilder() *ListPoolRespResponseHealthMonitorRespBuilder {
	s := &ListPoolRespResponseHealthMonitorResp{}
	b := &ListPoolRespResponseHealthMonitorRespBuilder{s: s}
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthDelay(v int32) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthExpectedCode(v string) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthMaxRetries(v int32) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthHttpMethod(v string) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthId(v string) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthId = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthType(v ListPoolRespResponseHealthMonitorRespHealthTypeEnum) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthType = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthUrlPath(v string) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) AdminStateUp(v bool) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) HealthTimeout(v int32) *ListPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *ListPoolRespResponseHealthMonitorRespBuilder) Build() *ListPoolRespResponseHealthMonitorResp {
	return b.s
}


