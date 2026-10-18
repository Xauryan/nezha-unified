// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBatchHealthStatusRespPath struct {
    position.Path
    // 负载均衡后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s GetBatchHealthStatusRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetBatchHealthStatusRespPath) GoString() string {
	return s.String()
}

func (s GetBatchHealthStatusRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBatchHealthStatusRespPath) SetPoolId(v string) *GetBatchHealthStatusRespPath {
	s.PoolId = &v
	return s
}


type GetBatchHealthStatusRespPathBuilder struct {
	s *GetBatchHealthStatusRespPath
}

func NewGetBatchHealthStatusRespPathBuilder() *GetBatchHealthStatusRespPathBuilder {
	s := &GetBatchHealthStatusRespPath{}
	b := &GetBatchHealthStatusRespPathBuilder{s: s}
	return b
}

func (b *GetBatchHealthStatusRespPathBuilder) PoolId(v string) *GetBatchHealthStatusRespPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetBatchHealthStatusRespPathBuilder) Build() *GetBatchHealthStatusRespPath {
	return b.s
}


