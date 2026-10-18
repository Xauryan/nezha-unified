// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPoolRespPath struct {
    position.Path
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s GetPoolRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetPoolRespPath) GoString() string {
	return s.String()
}

func (s GetPoolRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolRespPath) SetPoolId(v string) *GetPoolRespPath {
	s.PoolId = &v
	return s
}


type GetPoolRespPathBuilder struct {
	s *GetPoolRespPath
}

func NewGetPoolRespPathBuilder() *GetPoolRespPathBuilder {
	s := &GetPoolRespPath{}
	b := &GetPoolRespPathBuilder{s: s}
	return b
}

func (b *GetPoolRespPathBuilder) PoolId(v string) *GetPoolRespPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetPoolRespPathBuilder) Build() *GetPoolRespPath {
	return b.s
}


