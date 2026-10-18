// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePoolAsyncPath struct {
    position.Path
    // 后端服务器组ID
	PoolId *string `json:"poolId,omitempty"`
}

func (s DeletePoolAsyncPath) String() string {
	return utils.Beautify(s)
}

func (s DeletePoolAsyncPath) GoString() string {
	return s.String()
}

func (s DeletePoolAsyncPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePoolAsyncPath) SetPoolId(v string) *DeletePoolAsyncPath {
	s.PoolId = &v
	return s
}


type DeletePoolAsyncPathBuilder struct {
	s *DeletePoolAsyncPath
}

func NewDeletePoolAsyncPathBuilder() *DeletePoolAsyncPathBuilder {
	s := &DeletePoolAsyncPath{}
	b := &DeletePoolAsyncPathBuilder{s: s}
	return b
}

func (b *DeletePoolAsyncPathBuilder) PoolId(v string) *DeletePoolAsyncPathBuilder {
	b.s.PoolId = &v
	return b
}

func (b *DeletePoolAsyncPathBuilder) Build() *DeletePoolAsyncPath {
	return b.s
}


