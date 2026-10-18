// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderDeleteWithHttpStatusResponseBody struct {

    // 是否删除
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

func (s ElbOrderDeleteWithHttpStatusResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderDeleteWithHttpStatusResponseBody) GoString() string {
	return s.String()
}

func (s ElbOrderDeleteWithHttpStatusResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderDeleteWithHttpStatusResponseBody) SetIsDeleted(v bool) *ElbOrderDeleteWithHttpStatusResponseBody {
	s.IsDeleted = &v
	return s
}


type ElbOrderDeleteWithHttpStatusResponseBodyBuilder struct {
	s *ElbOrderDeleteWithHttpStatusResponseBody
}

func NewElbOrderDeleteWithHttpStatusResponseBodyBuilder() *ElbOrderDeleteWithHttpStatusResponseBodyBuilder {
	s := &ElbOrderDeleteWithHttpStatusResponseBody{}
	b := &ElbOrderDeleteWithHttpStatusResponseBodyBuilder{s: s}
	return b
}

func (b *ElbOrderDeleteWithHttpStatusResponseBodyBuilder) IsDeleted(v bool) *ElbOrderDeleteWithHttpStatusResponseBodyBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *ElbOrderDeleteWithHttpStatusResponseBodyBuilder) Build() *ElbOrderDeleteWithHttpStatusResponseBody {
	return b.s
}


