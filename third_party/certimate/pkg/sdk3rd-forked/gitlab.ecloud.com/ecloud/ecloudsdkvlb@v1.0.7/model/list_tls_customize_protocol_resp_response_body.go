// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListTLSCustomizeProtocolRespResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListTLSCustomizeProtocolRespResponseContent `json:"content,omitempty"`
}

func (s ListTLSCustomizeProtocolRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListTLSCustomizeProtocolRespResponseBody) GoString() string {
	return s.String()
}

func (s ListTLSCustomizeProtocolRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListTLSCustomizeProtocolRespResponseBody) SetTotal(v int32) *ListTLSCustomizeProtocolRespResponseBody {
	s.Total = &v
	return s
}

func (s *ListTLSCustomizeProtocolRespResponseBody) SetContent(v []ListTLSCustomizeProtocolRespResponseContent) *ListTLSCustomizeProtocolRespResponseBody {
	s.Content = &v
	return s
}


type ListTLSCustomizeProtocolRespResponseBodyBuilder struct {
	s *ListTLSCustomizeProtocolRespResponseBody
}

func NewListTLSCustomizeProtocolRespResponseBodyBuilder() *ListTLSCustomizeProtocolRespResponseBodyBuilder {
	s := &ListTLSCustomizeProtocolRespResponseBody{}
	b := &ListTLSCustomizeProtocolRespResponseBodyBuilder{s: s}
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBodyBuilder) Total(v int32) *ListTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBodyBuilder) Content(v []ListTLSCustomizeProtocolRespResponseContent) *ListTLSCustomizeProtocolRespResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListTLSCustomizeProtocolRespResponseBodyBuilder) Build() *ListTLSCustomizeProtocolRespResponseBody {
	return b.s
}


