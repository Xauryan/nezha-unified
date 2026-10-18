// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetHealthStatusRespPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 负载均衡后端服务器ID
	MemberId *string `json:"memberId,omitempty"`
}

func (s GetHealthStatusRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetHealthStatusRespPath) GoString() string {
	return s.String()
}

func (s GetHealthStatusRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetHealthStatusRespPath) SetPoolId(v string) *GetHealthStatusRespPath {
	s.PoolId = &v
	return s
}

func (s *GetHealthStatusRespPath) SetMemberId(v string) *GetHealthStatusRespPath {
	s.MemberId = &v
	return s
}


type GetHealthStatusRespPathBuilder struct {
	s *GetHealthStatusRespPath
}

func NewGetHealthStatusRespPathBuilder() *GetHealthStatusRespPathBuilder {
	s := &GetHealthStatusRespPath{}
	b := &GetHealthStatusRespPathBuilder{s: s}
	return b
}

func (b *GetHealthStatusRespPathBuilder) PoolId(v string) *GetHealthStatusRespPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetHealthStatusRespPathBuilder) MemberId(v string) *GetHealthStatusRespPathBuilder {
	b.s.MemberId = &v
	return b
}

func (b *GetHealthStatusRespPathBuilder) Build() *GetHealthStatusRespPath {
	return b.s
}


