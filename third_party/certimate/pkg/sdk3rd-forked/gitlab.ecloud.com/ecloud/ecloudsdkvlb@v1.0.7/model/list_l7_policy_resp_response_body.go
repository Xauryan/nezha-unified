// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListL7PolicyRespResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListL7PolicyRespResponseContent `json:"content,omitempty"`
}

func (s ListL7PolicyRespResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListL7PolicyRespResponseBody) GoString() string {
	return s.String()
}

func (s ListL7PolicyRespResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListL7PolicyRespResponseBody) SetTotal(v int32) *ListL7PolicyRespResponseBody {
	s.Total = &v
	return s
}

func (s *ListL7PolicyRespResponseBody) SetContent(v []ListL7PolicyRespResponseContent) *ListL7PolicyRespResponseBody {
	s.Content = &v
	return s
}


type ListL7PolicyRespResponseBodyBuilder struct {
	s *ListL7PolicyRespResponseBody
}

func NewListL7PolicyRespResponseBodyBuilder() *ListL7PolicyRespResponseBodyBuilder {
	s := &ListL7PolicyRespResponseBody{}
	b := &ListL7PolicyRespResponseBodyBuilder{s: s}
	return b
}

func (b *ListL7PolicyRespResponseBodyBuilder) Total(v int32) *ListL7PolicyRespResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListL7PolicyRespResponseBodyBuilder) Content(v []ListL7PolicyRespResponseContent) *ListL7PolicyRespResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListL7PolicyRespResponseBodyBuilder) Build() *ListL7PolicyRespResponseBody {
	return b.s
}


