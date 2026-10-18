// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetL7PolicyRespPath struct {
    position.Path
    // 七层转发策略ID
	L7PolicyId *string `json:"l7PolicyId,omitempty"`
}

func (s GetL7PolicyRespPath) String() string {
	return utils.Beautify(s)
}

func (s GetL7PolicyRespPath) GoString() string {
	return s.String()
}

func (s GetL7PolicyRespPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetL7PolicyRespPath) SetL7PolicyId(v string) *GetL7PolicyRespPath {
	s.L7PolicyId = &v
	return s
}


type GetL7PolicyRespPathBuilder struct {
	s *GetL7PolicyRespPath
}

func NewGetL7PolicyRespPathBuilder() *GetL7PolicyRespPathBuilder {
	s := &GetL7PolicyRespPath{}
	b := &GetL7PolicyRespPathBuilder{s: s}
	return b
}

func (b *GetL7PolicyRespPathBuilder) L7PolicyId(v string) *GetL7PolicyRespPathBuilder {
	b.s.L7PolicyId = &v
	return b
}

func (b *GetL7PolicyRespPathBuilder) Build() *GetL7PolicyRespPath {
	return b.s
}


