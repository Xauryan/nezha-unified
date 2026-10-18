// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateTCAclHeader struct {
    position.Header
    // 请求头部,region值
	Region *string `json:"region,omitempty"`
}

func (s CreateTCAclHeader) String() string {
	return utils.Beautify(s)
}

func (s CreateTCAclHeader) GoString() string {
	return s.String()
}

func (s CreateTCAclHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateTCAclHeader) SetRegion(v string) *CreateTCAclHeader {
	s.Region = &v
	return s
}


type CreateTCAclHeaderBuilder struct {
	s *CreateTCAclHeader
}

func NewCreateTCAclHeaderBuilder() *CreateTCAclHeaderBuilder {
	s := &CreateTCAclHeader{}
	b := &CreateTCAclHeaderBuilder{s: s}
	return b
}

func (b *CreateTCAclHeaderBuilder) Region(v string) *CreateTCAclHeaderBuilder {
	b.s.Region = &v
	return b
}

func (b *CreateTCAclHeaderBuilder) Build() *CreateTCAclHeader {
	return b.s
}


