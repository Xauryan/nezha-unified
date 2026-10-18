// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadbalanceResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListLoadbalanceResponseContent `json:"content,omitempty"`
}

func (s ListLoadbalanceResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceResponseBody) GoString() string {
	return s.String()
}

func (s ListLoadbalanceResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceResponseBody) SetTotal(v int32) *ListLoadbalanceResponseBody {
	s.Total = &v
	return s
}

func (s *ListLoadbalanceResponseBody) SetContent(v []ListLoadbalanceResponseContent) *ListLoadbalanceResponseBody {
	s.Content = &v
	return s
}


type ListLoadbalanceResponseBodyBuilder struct {
	s *ListLoadbalanceResponseBody
}

func NewListLoadbalanceResponseBodyBuilder() *ListLoadbalanceResponseBodyBuilder {
	s := &ListLoadbalanceResponseBody{}
	b := &ListLoadbalanceResponseBodyBuilder{s: s}
	return b
}

func (b *ListLoadbalanceResponseBodyBuilder) Total(v int32) *ListLoadbalanceResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListLoadbalanceResponseBodyBuilder) Content(v []ListLoadbalanceResponseContent) *ListLoadbalanceResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListLoadbalanceResponseBodyBuilder) Build() *ListLoadbalanceResponseBody {
	return b.s
}


