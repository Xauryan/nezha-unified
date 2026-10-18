// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPoolRespResponseHealthMonitorRespHealthTypeEnum string

// List of HealthType
const (
    GetPoolRespResponseHealthMonitorRespHealthTypeEnumHttp GetPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTP"
    GetPoolRespResponseHealthMonitorRespHealthTypeEnumHttps GetPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTPS"
    GetPoolRespResponseHealthMonitorRespHealthTypeEnumTcp GetPoolRespResponseHealthMonitorRespHealthTypeEnum = "TCP"
    GetPoolRespResponseHealthMonitorRespHealthTypeEnumPing GetPoolRespResponseHealthMonitorRespHealthTypeEnum = "PING"
    GetPoolRespResponseHealthMonitorRespHealthTypeEnumUdp GetPoolRespResponseHealthMonitorRespHealthTypeEnum = "UDP"
    GetPoolRespResponseHealthMonitorRespHealthTypeEnumHttp10 GetPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTP10"
    GetPoolRespResponseHealthMonitorRespHealthTypeEnumHttp11 GetPoolRespResponseHealthMonitorRespHealthTypeEnum = "HTTP11"
)

type GetPoolRespResponseHealthMonitorResp struct {

    // 健康监控时间间隔（单位：秒）
	HealthDelay *int32 `json:"healthDelay,omitempty"`
    // 期望HTTP状态代码
	HealthExpectedCode *string `json:"healthExpectedCode,omitempty"`
    // 监听器最大尝试次数
	HealthMaxRetries *int32 `json:"healthMaxRetries,omitempty"`
    // 监听器绑定的检查方式为HTTP的用于检查的方法，取值范围：GET，POST，HEAD
	HealthHttpMethod *string `json:"healthHttpMethod,omitempty"`
    // 健康检查的资源ID
	HealthId *string `json:"healthId,omitempty"`
    // 健康监控类型
	HealthType *GetPoolRespResponseHealthMonitorRespHealthTypeEnum `json:"healthType,omitempty"`
    // 用于监控的URL，必须以“/”打头
	HealthUrlPath *string `json:"healthUrlPath,omitempty"`
    // 健康检查开关
	AdminStateUp *bool `json:"adminStateUp,omitempty"`
    // 最大等待时间（单位：秒）
	HealthTimeout *int32 `json:"healthTimeout,omitempty"`
}

func (s GetPoolRespResponseHealthMonitorResp) String() string {
	return utils.Beautify(s)
}

func (s GetPoolRespResponseHealthMonitorResp) GoString() string {
	return s.String()
}

func (s GetPoolRespResponseHealthMonitorResp) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthDelay(v int32) *GetPoolRespResponseHealthMonitorResp {
	s.HealthDelay = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthExpectedCode(v string) *GetPoolRespResponseHealthMonitorResp {
	s.HealthExpectedCode = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthMaxRetries(v int32) *GetPoolRespResponseHealthMonitorResp {
	s.HealthMaxRetries = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthHttpMethod(v string) *GetPoolRespResponseHealthMonitorResp {
	s.HealthHttpMethod = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthId(v string) *GetPoolRespResponseHealthMonitorResp {
	s.HealthId = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthType(v GetPoolRespResponseHealthMonitorRespHealthTypeEnum) *GetPoolRespResponseHealthMonitorResp {
	s.HealthType = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthUrlPath(v string) *GetPoolRespResponseHealthMonitorResp {
	s.HealthUrlPath = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetAdminStateUp(v bool) *GetPoolRespResponseHealthMonitorResp {
	s.AdminStateUp = &v
	return s
}

func (s *GetPoolRespResponseHealthMonitorResp) SetHealthTimeout(v int32) *GetPoolRespResponseHealthMonitorResp {
	s.HealthTimeout = &v
	return s
}


type GetPoolRespResponseHealthMonitorRespBuilder struct {
	s *GetPoolRespResponseHealthMonitorResp
}

func NewGetPoolRespResponseHealthMonitorRespBuilder() *GetPoolRespResponseHealthMonitorRespBuilder {
	s := &GetPoolRespResponseHealthMonitorResp{}
	b := &GetPoolRespResponseHealthMonitorRespBuilder{s: s}
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthDelay(v int32) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthDelay = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthExpectedCode(v string) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthExpectedCode = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthMaxRetries(v int32) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthMaxRetries = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthHttpMethod(v string) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthHttpMethod = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthId(v string) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthId = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthType(v GetPoolRespResponseHealthMonitorRespHealthTypeEnum) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthType = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthUrlPath(v string) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthUrlPath = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) AdminStateUp(v bool) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.AdminStateUp = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) HealthTimeout(v int32) *GetPoolRespResponseHealthMonitorRespBuilder {
	b.s.HealthTimeout = &v
	return b
}

func (b *GetPoolRespResponseHealthMonitorRespBuilder) Build() *GetPoolRespResponseHealthMonitorResp {
	return b.s
}


