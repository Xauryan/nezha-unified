// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadbalanceResponseTags struct {

    // 标签ID
	TagId *string `json:"tagId,omitempty"`
    // 标签值
	TagValue *string `json:"tagValue,omitempty"`
    // 标签键
	TagKey *string `json:"tagKey,omitempty"`
}

func (s ListLoadbalanceResponseTags) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceResponseTags) GoString() string {
	return s.String()
}

func (s ListLoadbalanceResponseTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceResponseTags) SetTagId(v string) *ListLoadbalanceResponseTags {
	s.TagId = &v
	return s
}

func (s *ListLoadbalanceResponseTags) SetTagValue(v string) *ListLoadbalanceResponseTags {
	s.TagValue = &v
	return s
}

func (s *ListLoadbalanceResponseTags) SetTagKey(v string) *ListLoadbalanceResponseTags {
	s.TagKey = &v
	return s
}


type ListLoadbalanceResponseTagsBuilder struct {
	s *ListLoadbalanceResponseTags
}

func NewListLoadbalanceResponseTagsBuilder() *ListLoadbalanceResponseTagsBuilder {
	s := &ListLoadbalanceResponseTags{}
	b := &ListLoadbalanceResponseTagsBuilder{s: s}
	return b
}

func (b *ListLoadbalanceResponseTagsBuilder) TagId(v string) *ListLoadbalanceResponseTagsBuilder {
	b.s.TagId = &v
	return b
}

func (b *ListLoadbalanceResponseTagsBuilder) TagValue(v string) *ListLoadbalanceResponseTagsBuilder {
	b.s.TagValue = &v
	return b
}

func (b *ListLoadbalanceResponseTagsBuilder) TagKey(v string) *ListLoadbalanceResponseTagsBuilder {
	b.s.TagKey = &v
	return b
}

func (b *ListLoadbalanceResponseTagsBuilder) Build() *ListLoadbalanceResponseTags {
	return b.s
}


