// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePoolPath struct {
    position.Path
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s DeletePoolPath) String() string {
	return utils.Beautify(s)
}

func (s DeletePoolPath) GoString() string {
	return s.String()
}

func (s DeletePoolPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePoolPath) SetPoolId(v string) *DeletePoolPath {
	s.PoolId = &v
	return s
}


type DeletePoolPathBuilder struct {
	s *DeletePoolPath
}

func NewDeletePoolPathBuilder() *DeletePoolPathBuilder {
	s := &DeletePoolPath{}
	b := &DeletePoolPathBuilder{s: s}
	return b
}

func (b *DeletePoolPathBuilder) PoolId(v string) *DeletePoolPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *DeletePoolPathBuilder) Build() *DeletePoolPath {
	return b.s
}


