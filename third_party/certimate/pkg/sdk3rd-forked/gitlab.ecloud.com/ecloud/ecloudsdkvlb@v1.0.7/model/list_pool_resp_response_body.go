// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListPoolRespResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListPoolRespResponseContent `json:"content,omitempty"`
}

func (s ListPoolRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListPoolRespResponseBody) GoString() string {
	return s.String()
}

func (s ListPoolRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListPoolRespResponseBody) SetTotal(v int32) *ListPoolRespResponseBody {
	s.Total = &v
	return s
}

func (s *ListPoolRespResponseBody) SetContent(v []ListPoolRespResponseContent) *ListPoolRespResponseBody {
	s.Content = &v
	return s
}


type ListPoolRespResponseBodyBuilder struct {
	s *ListPoolRespResponseBody
}

func NewListPoolRespResponseBodyBuilder() *ListPoolRespResponseBodyBuilder {
	s := &ListPoolRespResponseBody{}
	b := &ListPoolRespResponseBodyBuilder{s: s}
	return b
}

func (b *ListPoolRespResponseBodyBuilder) Total(v int32) *ListPoolRespResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListPoolRespResponseBodyBuilder) Content(v []ListPoolRespResponseContent) *ListPoolRespResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListPoolRespResponseBodyBuilder) Build() *ListPoolRespResponseBody {
	return b.s
}


