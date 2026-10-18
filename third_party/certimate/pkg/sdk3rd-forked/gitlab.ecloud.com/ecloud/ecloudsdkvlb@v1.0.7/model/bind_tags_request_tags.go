// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindTagsRequestTags struct {

    // 标签值
	Value *string `json:"value,omitempty"`
    // 标签键
	Key *string `json:"key,omitempty"`
}

func (s BindTagsRequestTags) String() string {
	return utils.Beautify(s)
}

func (s BindTagsRequestTags) GoString() string {
	return s.String()
}

func (s BindTagsRequestTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindTagsRequestTags) SetValue(v string) *BindTagsRequestTags {
	s.Value = &v
	return s
}

func (s *BindTagsRequestTags) SetKey(v string) *BindTagsRequestTags {
	s.Key = &v
	return s
}


type BindTagsRequestTagsBuilder struct {
	s *BindTagsRequestTags
}

func NewBindTagsRequestTagsBuilder() *BindTagsRequestTagsBuilder {
	s := &BindTagsRequestTags{}
	b := &BindTagsRequestTagsBuilder{s: s}
	return b
}

func (b *BindTagsRequestTagsBuilder) Value(v string) *BindTagsRequestTagsBuilder {
	b.s.Value = &v
	return b
}

func (b *BindTagsRequestTagsBuilder) Key(v string) *BindTagsRequestTagsBuilder {
	b.s.Key = &v
	return b
}

func (b *BindTagsRequestTagsBuilder) Build() *BindTagsRequestTags {
	return b.s
}


