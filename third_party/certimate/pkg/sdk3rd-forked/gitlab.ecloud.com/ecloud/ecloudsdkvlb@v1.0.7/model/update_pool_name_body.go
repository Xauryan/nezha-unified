// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdatePoolNameBody struct {
    position.Body
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
    // 后端服务器组名称为5~64位的英文、数字、下划线、中划线等的组合
	PoolName *string `json:"poolName,omitempty"`
}

func (s UpdatePoolNameBody) String() string {
	return utils.Beautify(s)
}

func (s UpdatePoolNameBody) GoString() string {
	return s.String()
}

func (s UpdatePoolNameBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdatePoolNameBody) SetPoolId(v string) *UpdatePoolNameBody {
	s.PoolId = &v
	return s
}

func (s *UpdatePoolNameBody) SetPoolName(v string) *UpdatePoolNameBody {
	s.PoolName = &v
	return s
}


type UpdatePoolNameBodyBuilder struct {
	s *UpdatePoolNameBody
}

func NewUpdatePoolNameBodyBuilder() *UpdatePoolNameBodyBuilder {
	s := &UpdatePoolNameBody{}
	b := &UpdatePoolNameBodyBuilder{s: s}
	return b
}

func (b *UpdatePoolNameBodyBuilder) PoolId(v string) *UpdatePoolNameBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *UpdatePoolNameBodyBuilder) PoolName(v string) *UpdatePoolNameBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *UpdatePoolNameBodyBuilder) Build() *UpdatePoolNameBody {
	return b.s
}


