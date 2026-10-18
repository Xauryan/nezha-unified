// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteListenerWithHttpStatusHeader struct {
    position.Header
    // 请求头部,region值
	Region *string `json:"region,omitempty"`
}

func (s DeleteListenerWithHttpStatusHeader) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerWithHttpStatusHeader) GoString() string {
	return s.String()
}

func (s DeleteListenerWithHttpStatusHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerWithHttpStatusHeader) SetRegion(v string) *DeleteListenerWithHttpStatusHeader {
	s.Region = &v
	return s
}


type DeleteListenerWithHttpStatusHeaderBuilder struct {
	s *DeleteListenerWithHttpStatusHeader
}

func NewDeleteListenerWithHttpStatusHeaderBuilder() *DeleteListenerWithHttpStatusHeaderBuilder {
	s := &DeleteListenerWithHttpStatusHeader{}
	b := &DeleteListenerWithHttpStatusHeaderBuilder{s: s}
	return b
}

func (b *DeleteListenerWithHttpStatusHeaderBuilder) Region(v string) *DeleteListenerWithHttpStatusHeaderBuilder {
	b.s.Region = &v
	return b
}

func (b *DeleteListenerWithHttpStatusHeaderBuilder) Build() *DeleteListenerWithHttpStatusHeader {
	return b.s
}


