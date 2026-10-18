// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderDeleteResponseBody struct {

    // 是否删除
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

func (s ElbOrderDeleteResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderDeleteResponseBody) GoString() string {
	return s.String()
}

func (s ElbOrderDeleteResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderDeleteResponseBody) SetIsDeleted(v bool) *ElbOrderDeleteResponseBody {
	s.IsDeleted = &v
	return s
}


type ElbOrderDeleteResponseBodyBuilder struct {
	s *ElbOrderDeleteResponseBody
}

func NewElbOrderDeleteResponseBodyBuilder() *ElbOrderDeleteResponseBodyBuilder {
	s := &ElbOrderDeleteResponseBody{}
	b := &ElbOrderDeleteResponseBodyBuilder{s: s}
	return b
}

func (b *ElbOrderDeleteResponseBodyBuilder) IsDeleted(v bool) *ElbOrderDeleteResponseBodyBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *ElbOrderDeleteResponseBodyBuilder) Build() *ElbOrderDeleteResponseBody {
	return b.s
}


