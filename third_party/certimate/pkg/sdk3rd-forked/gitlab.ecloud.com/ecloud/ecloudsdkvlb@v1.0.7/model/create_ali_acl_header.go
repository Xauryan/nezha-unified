// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateAliAclHeader struct {
    position.Header
    // 请求头部,region值
	Region *string `json:"region,omitempty"`
}

func (s CreateAliAclHeader) String() string {
	return utils.Beautify(s)
}

func (s CreateAliAclHeader) GoString() string {
	return s.String()
}

func (s CreateAliAclHeader) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateAliAclHeader) SetRegion(v string) *CreateAliAclHeader {
	s.Region = &v
	return s
}


type CreateAliAclHeaderBuilder struct {
	s *CreateAliAclHeader
}

func NewCreateAliAclHeaderBuilder() *CreateAliAclHeaderBuilder {
	s := &CreateAliAclHeader{}
	b := &CreateAliAclHeaderBuilder{s: s}
	return b
}

func (b *CreateAliAclHeaderBuilder) Region(v string) *CreateAliAclHeaderBuilder {
	b.s.Region = &v
	return b
}

func (b *CreateAliAclHeaderBuilder) Build() *CreateAliAclHeader {
	return b.s
}


