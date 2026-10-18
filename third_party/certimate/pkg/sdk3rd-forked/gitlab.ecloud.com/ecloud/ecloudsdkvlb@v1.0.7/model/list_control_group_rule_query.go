// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupRuleQuery struct {
    position.Query
    // 描述
	Description *string `json:"description,omitempty"`
    // 列表分页每页大小，小于0将默认为10
	PageSize *int32 `json:"pageSize,omitempty"`
    // 列表分页页数，小于1会默认第1页
	Page *int32 `json:"page,omitempty"`
    // 根据IP或描述模糊查询
	QueryWord *string `json:"queryWord,omitempty"`
    // 访问控制组ID
	ControlGroupId *string `json:"controlGroupId,omitempty"`
    // 查询类型（description或ipAddress，该参数与queryword同时传参）
	QueryType *string `json:"queryType,omitempty"`
}

func (s ListControlGroupRuleQuery) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupRuleQuery) GoString() string {
	return s.String()
}

func (s ListControlGroupRuleQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupRuleQuery) SetDescription(v string) *ListControlGroupRuleQuery {
	s.Description = &v
	return s
}

func (s *ListControlGroupRuleQuery) SetPageSize(v int32) *ListControlGroupRuleQuery {
	s.PageSize = &v
	return s
}

func (s *ListControlGroupRuleQuery) SetPage(v int32) *ListControlGroupRuleQuery {
	s.Page = &v
	return s
}

func (s *ListControlGroupRuleQuery) SetQueryWord(v string) *ListControlGroupRuleQuery {
	s.QueryWord = &v
	return s
}

func (s *ListControlGroupRuleQuery) SetControlGroupId(v string) *ListControlGroupRuleQuery {
	s.ControlGroupId = &v
	return s
}

func (s *ListControlGroupRuleQuery) SetQueryType(v string) *ListControlGroupRuleQuery {
	s.QueryType = &v
	return s
}


type ListControlGroupRuleQueryBuilder struct {
	s *ListControlGroupRuleQuery
}

func NewListControlGroupRuleQueryBuilder() *ListControlGroupRuleQueryBuilder {
	s := &ListControlGroupRuleQuery{}
	b := &ListControlGroupRuleQueryBuilder{s: s}
	return b
}

func (b *ListControlGroupRuleQueryBuilder) Description(v string) *ListControlGroupRuleQueryBuilder {
	b.s.Description = &v
	return b
}

func (b *ListControlGroupRuleQueryBuilder) PageSize(v int32) *ListControlGroupRuleQueryBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ListControlGroupRuleQueryBuilder) Page(v int32) *ListControlGroupRuleQueryBuilder {
	b.s.Page = &v
	return b
}

func (b *ListControlGroupRuleQueryBuilder) QueryWord(v string) *ListControlGroupRuleQueryBuilder {
	b.s.QueryWord = &v
	return b
}

func (b *ListControlGroupRuleQueryBuilder) ControlGroupId(v string) *ListControlGroupRuleQueryBuilder {
	b.s.ControlGroupId = &v
	return b
}

func (b *ListControlGroupRuleQueryBuilder) QueryType(v string) *ListControlGroupRuleQueryBuilder {
	b.s.QueryType = &v
	return b
}

func (b *ListControlGroupRuleQueryBuilder) Build() *ListControlGroupRuleQuery {
	return b.s
}


