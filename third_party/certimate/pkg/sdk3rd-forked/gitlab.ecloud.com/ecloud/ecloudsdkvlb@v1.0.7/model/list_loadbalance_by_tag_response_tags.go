// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadbalanceByTagResponseTags struct {

    // 标签ID
	TagId *string `json:"tagId,omitempty"`
    // 标签值
	TagValue *string `json:"tagValue,omitempty"`
    // 标签键
	TagKey *string `json:"tagKey,omitempty"`
}

func (s ListLoadbalanceByTagResponseTags) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceByTagResponseTags) GoString() string {
	return s.String()
}

func (s ListLoadbalanceByTagResponseTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceByTagResponseTags) SetTagId(v string) *ListLoadbalanceByTagResponseTags {
	s.TagId = &v
	return s
}

func (s *ListLoadbalanceByTagResponseTags) SetTagValue(v string) *ListLoadbalanceByTagResponseTags {
	s.TagValue = &v
	return s
}

func (s *ListLoadbalanceByTagResponseTags) SetTagKey(v string) *ListLoadbalanceByTagResponseTags {
	s.TagKey = &v
	return s
}


type ListLoadbalanceByTagResponseTagsBuilder struct {
	s *ListLoadbalanceByTagResponseTags
}

func NewListLoadbalanceByTagResponseTagsBuilder() *ListLoadbalanceByTagResponseTagsBuilder {
	s := &ListLoadbalanceByTagResponseTags{}
	b := &ListLoadbalanceByTagResponseTagsBuilder{s: s}
	return b
}

func (b *ListLoadbalanceByTagResponseTagsBuilder) TagId(v string) *ListLoadbalanceByTagResponseTagsBuilder {
	b.s.TagId = &v
	return b
}

func (b *ListLoadbalanceByTagResponseTagsBuilder) TagValue(v string) *ListLoadbalanceByTagResponseTagsBuilder {
	b.s.TagValue = &v
	return b
}

func (b *ListLoadbalanceByTagResponseTagsBuilder) TagKey(v string) *ListLoadbalanceByTagResponseTagsBuilder {
	b.s.TagKey = &v
	return b
}

func (b *ListLoadbalanceByTagResponseTagsBuilder) Build() *ListLoadbalanceByTagResponseTags {
	return b.s
}


