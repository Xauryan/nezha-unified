// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePolicyAsyncPath struct {
    position.Path
    // 七层转发策略ID
	L7PolicyId *string `json:"l7PolicyId,omitempty"`
}

func (s DeletePolicyAsyncPath) String() string {
	return utils.Beautify(s)
}

func (s DeletePolicyAsyncPath) GoString() string {
	return s.String()
}

func (s DeletePolicyAsyncPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePolicyAsyncPath) SetL7PolicyId(v string) *DeletePolicyAsyncPath {
	s.L7PolicyId = &v
	return s
}


type DeletePolicyAsyncPathBuilder struct {
	s *DeletePolicyAsyncPath
}

func NewDeletePolicyAsyncPathBuilder() *DeletePolicyAsyncPathBuilder {
	s := &DeletePolicyAsyncPath{}
	b := &DeletePolicyAsyncPathBuilder{s: s}
	return b
}

func (b *DeletePolicyAsyncPathBuilder) L7PolicyId(v string) *DeletePolicyAsyncPathBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *DeletePolicyAsyncPathBuilder) Build() *DeletePolicyAsyncPath {
	return b.s
}


