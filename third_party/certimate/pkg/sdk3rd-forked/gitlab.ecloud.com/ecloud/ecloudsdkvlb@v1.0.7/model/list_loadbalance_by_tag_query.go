// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadbalanceByTagQuery struct {
    position.Query
    // 标签值，不为空时，标签键不能为空
	TagValue *string `json:"tagValue,omitempty"`
    // 是否返回标签内容
	IsDescribeTags *bool `json:"isDescribeTags,omitempty"`
    // 标签键，不为空时，标签值不能为空
	TagKey *string `json:"tagKey,omitempty"`
}

func (s ListLoadbalanceByTagQuery) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceByTagQuery) GoString() string {
	return s.String()
}

func (s ListLoadbalanceByTagQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceByTagQuery) SetTagValue(v string) *ListLoadbalanceByTagQuery {
	s.TagValue = &v
	return s
}

func (s *ListLoadbalanceByTagQuery) SetIsDescribeTags(v bool) *ListLoadbalanceByTagQuery {
	s.IsDescribeTags = &v
	return s
}

func (s *ListLoadbalanceByTagQuery) SetTagKey(v string) *ListLoadbalanceByTagQuery {
	s.TagKey = &v
	return s
}


type ListLoadbalanceByTagQueryBuilder struct {
	s *ListLoadbalanceByTagQuery
}

func NewListLoadbalanceByTagQueryBuilder() *ListLoadbalanceByTagQueryBuilder {
	s := &ListLoadbalanceByTagQuery{}
	b := &ListLoadbalanceByTagQueryBuilder{s: s}
	return b
}

func (b *ListLoadbalanceByTagQueryBuilder) TagValue(v string) *ListLoadbalanceByTagQueryBuilder {
	b.s.TagValue = &v
	return b
}

func (b *ListLoadbalanceByTagQueryBuilder) IsDescribeTags(v bool) *ListLoadbalanceByTagQueryBuilder {
	b.s.IsDescribeTags = &v
	return b
}

func (b *ListLoadbalanceByTagQueryBuilder) TagKey(v string) *ListLoadbalanceByTagQueryBuilder {
	b.s.TagKey = &v
	return b
}

func (b *ListLoadbalanceByTagQueryBuilder) Build() *ListLoadbalanceByTagQuery {
	return b.s
}


