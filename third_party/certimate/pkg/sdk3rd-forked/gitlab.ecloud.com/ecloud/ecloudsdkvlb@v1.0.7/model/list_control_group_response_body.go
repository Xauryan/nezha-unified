// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListControlGroupResponseContent `json:"content,omitempty"`
}

func (s ListControlGroupResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupResponseBody) GoString() string {
	return s.String()
}

func (s ListControlGroupResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupResponseBody) SetTotal(v int32) *ListControlGroupResponseBody {
	s.Total = &v
	return s
}

func (s *ListControlGroupResponseBody) SetContent(v []ListControlGroupResponseContent) *ListControlGroupResponseBody {
	s.Content = &v
	return s
}


type ListControlGroupResponseBodyBuilder struct {
	s *ListControlGroupResponseBody
}

func NewListControlGroupResponseBodyBuilder() *ListControlGroupResponseBodyBuilder {
	s := &ListControlGroupResponseBody{}
	b := &ListControlGroupResponseBodyBuilder{s: s}
	return b
}

func (b *ListControlGroupResponseBodyBuilder) Total(v int32) *ListControlGroupResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListControlGroupResponseBodyBuilder) Content(v []ListControlGroupResponseContent) *ListControlGroupResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListControlGroupResponseBodyBuilder) Build() *ListControlGroupResponseBody {
	return b.s
}


