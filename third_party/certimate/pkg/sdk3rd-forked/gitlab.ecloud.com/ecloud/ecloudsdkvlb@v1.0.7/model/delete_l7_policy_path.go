// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteL7PolicyPath struct {
    position.Path
    // 七层转发策略ID
	L7PolicyId *string `json:"l7PolicyId,omitempty"`
}

func (s DeleteL7PolicyPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteL7PolicyPath) GoString() string {
	return s.String()
}

func (s DeleteL7PolicyPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteL7PolicyPath) SetL7PolicyId(v string) *DeleteL7PolicyPath {
	s.L7PolicyId = &v
	return s
}


type DeleteL7PolicyPathBuilder struct {
	s *DeleteL7PolicyPath
}

func NewDeleteL7PolicyPathBuilder() *DeleteL7PolicyPathBuilder {
	s := &DeleteL7PolicyPath{}
	b := &DeleteL7PolicyPathBuilder{s: s}
	return b
}

func (b *DeleteL7PolicyPathBuilder) L7PolicyId(v string) *DeleteL7PolicyPathBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *DeleteL7PolicyPathBuilder) Build() *DeleteL7PolicyPath {
	return b.s
}


