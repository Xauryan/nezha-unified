// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderCreateAsyncRequestTags struct {

    // 标签值
	Value *string `json:"value,omitempty"`
    // 标签键
	Key *string `json:"key,omitempty"`
}

func (s ElbOrderCreateAsyncRequestTags) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncRequestTags) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncRequestTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncRequestTags) SetValue(v string) *ElbOrderCreateAsyncRequestTags {
	s.Value = &v
	return s
}

func (s *ElbOrderCreateAsyncRequestTags) SetKey(v string) *ElbOrderCreateAsyncRequestTags {
	s.Key = &v
	return s
}


type ElbOrderCreateAsyncRequestTagsBuilder struct {
	s *ElbOrderCreateAsyncRequestTags
}

func NewElbOrderCreateAsyncRequestTagsBuilder() *ElbOrderCreateAsyncRequestTagsBuilder {
	s := &ElbOrderCreateAsyncRequestTags{}
	b := &ElbOrderCreateAsyncRequestTagsBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncRequestTagsBuilder) Value(v string) *ElbOrderCreateAsyncRequestTagsBuilder {
	b.s.Value = &v
	return b
}

func (b *ElbOrderCreateAsyncRequestTagsBuilder) Key(v string) *ElbOrderCreateAsyncRequestTagsBuilder {
	b.s.Key = &v
	return b
}

func (b *ElbOrderCreateAsyncRequestTagsBuilder) Build() *ElbOrderCreateAsyncRequestTags {
	return b.s
}


