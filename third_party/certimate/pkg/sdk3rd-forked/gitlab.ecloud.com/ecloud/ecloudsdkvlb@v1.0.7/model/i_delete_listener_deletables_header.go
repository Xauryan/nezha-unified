// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type IDeleteListenerDeletablesHeader struct {
    position.Header
    // 请求头部,region值
	Region *string `json:"region,omitempty"`
}

func (s IDeleteListenerDeletablesHeader) String() string {
	return utils.Beautify(s)
}

func (s IDeleteListenerDeletablesHeader) GoString() string {
	return s.String()
}

func (s IDeleteListenerDeletablesHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IDeleteListenerDeletablesHeader) SetRegion(v string) *IDeleteListenerDeletablesHeader {
	s.Region = &v
	return s
}


type IDeleteListenerDeletablesHeaderBuilder struct {
	s *IDeleteListenerDeletablesHeader
}

func NewIDeleteListenerDeletablesHeaderBuilder() *IDeleteListenerDeletablesHeaderBuilder {
	s := &IDeleteListenerDeletablesHeader{}
	b := &IDeleteListenerDeletablesHeaderBuilder{s: s}
	return b
}

func (b *IDeleteListenerDeletablesHeaderBuilder) Region(v string) *IDeleteListenerDeletablesHeaderBuilder {
	b.s.Region = &v
	return b
}

func (b *IDeleteListenerDeletablesHeaderBuilder) Build() *IDeleteListenerDeletablesHeader {
	return b.s
}


