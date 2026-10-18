// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBatchHealthStatusRespResponseBody struct {

    // 主节点状态
	AdminUpState *bool `json:"adminUpState,omitempty"`
    // 后端服务器的底层状态，取值如下： up：正常 down：异常
	HealthStatus *string `json:"healthStatus,omitempty"`
    // 后端服务器ID
	Id *string `json:"id,omitempty"`
}

func (s GetBatchHealthStatusRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetBatchHealthStatusRespResponseBody) GoString() string {
	return s.String()
}

func (s GetBatchHealthStatusRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBatchHealthStatusRespResponseBody) SetAdminUpState(v bool) *GetBatchHealthStatusRespResponseBody {
	s.AdminUpState = &v
	return s
}

func (s *GetBatchHealthStatusRespResponseBody) SetHealthStatus(v string) *GetBatchHealthStatusRespResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *GetBatchHealthStatusRespResponseBody) SetId(v string) *GetBatchHealthStatusRespResponseBody {
	s.Id = &v
	return s
}


type GetBatchHealthStatusRespResponseBodyBuilder struct {
	s *GetBatchHealthStatusRespResponseBody
}

func NewGetBatchHealthStatusRespResponseBodyBuilder() *GetBatchHealthStatusRespResponseBodyBuilder {
	s := &GetBatchHealthStatusRespResponseBody{}
	b := &GetBatchHealthStatusRespResponseBodyBuilder{s: s}
	return b
}

func (b *GetBatchHealthStatusRespResponseBodyBuilder) AdminUpState(v bool) *GetBatchHealthStatusRespResponseBodyBuilder {
	b.s.AdminUpState = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBodyBuilder) HealthStatus(v string) *GetBatchHealthStatusRespResponseBodyBuilder {
	b.s.HealthStatus = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBodyBuilder) Id(v string) *GetBatchHealthStatusRespResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetBatchHealthStatusRespResponseBodyBuilder) Build() *GetBatchHealthStatusRespResponseBody {
	return b.s
}


