// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupRuleResponseBody struct {

    // 满足查询条件的总个数
	Total *int32 `json:"total,omitempty"`
    // 分页查询返回的具体结构体
	Content *[]ListControlGroupRuleResponseContent `json:"content,omitempty"`
}

func (s ListControlGroupRuleResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s ListControlGroupRuleResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupRuleResponseBody) SetTotal(v int32) *ListControlGroupRuleResponseBody {
	s.Total = &v
	return s
}

func (s *ListControlGroupRuleResponseBody) SetContent(v []ListControlGroupRuleResponseContent) *ListControlGroupRuleResponseBody {
	s.Content = &v
	return s
}


type ListControlGroupRuleResponseBodyBuilder struct {
	s *ListControlGroupRuleResponseBody
}

func NewListControlGroupRuleResponseBodyBuilder() *ListControlGroupRuleResponseBodyBuilder {
	s := &ListControlGroupRuleResponseBody{}
	b := &ListControlGroupRuleResponseBodyBuilder{s: s}
	return b
}

func (b *ListControlGroupRuleResponseBodyBuilder) Total(v int32) *ListControlGroupRuleResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ListControlGroupRuleResponseBodyBuilder) Content(v []ListControlGroupRuleResponseContent) *ListControlGroupRuleResponseBodyBuilder {
	b.s.Content = &v
	return b
}

func (b *ListControlGroupRuleResponseBodyBuilder) Build() *ListControlGroupRuleResponseBody {
	return b.s
}


