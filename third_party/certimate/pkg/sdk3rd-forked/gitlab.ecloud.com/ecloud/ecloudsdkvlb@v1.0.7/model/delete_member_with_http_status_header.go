// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteMemberWithHttpStatusHeader struct {
    position.Header
    // 请求头部,region值
	Region *string `json:"region,omitempty"`
}

func (s DeleteMemberWithHttpStatusHeader) String() string {
	return utils.Beautify(s)
}

func (s DeleteMemberWithHttpStatusHeader) GoString() string {
	return s.String()
}

func (s DeleteMemberWithHttpStatusHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteMemberWithHttpStatusHeader) SetRegion(v string) *DeleteMemberWithHttpStatusHeader {
	s.Region = &v
	return s
}


type DeleteMemberWithHttpStatusHeaderBuilder struct {
	s *DeleteMemberWithHttpStatusHeader
}

func NewDeleteMemberWithHttpStatusHeaderBuilder() *DeleteMemberWithHttpStatusHeaderBuilder {
	s := &DeleteMemberWithHttpStatusHeader{}
	b := &DeleteMemberWithHttpStatusHeaderBuilder{s: s}
	return b
}

func (b *DeleteMemberWithHttpStatusHeaderBuilder) Region(v string) *DeleteMemberWithHttpStatusHeaderBuilder {
	b.s.Region = &v
	return b
}

func (b *DeleteMemberWithHttpStatusHeaderBuilder) Build() *DeleteMemberWithHttpStatusHeader {
	return b.s
}


