// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderCreateAsyncWithHttpStatusRequestTags struct {

    // 标签值
	Value *string `json:"value,omitempty"`
    // 标签键
	Key *string `json:"key,omitempty"`
}

func (s ElbOrderCreateAsyncWithHttpStatusRequestTags) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderCreateAsyncWithHttpStatusRequestTags) GoString() string {
	return s.String()
}

func (s ElbOrderCreateAsyncWithHttpStatusRequestTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderCreateAsyncWithHttpStatusRequestTags) SetValue(v string) *ElbOrderCreateAsyncWithHttpStatusRequestTags {
	s.Value = &v
	return s
}

func (s *ElbOrderCreateAsyncWithHttpStatusRequestTags) SetKey(v string) *ElbOrderCreateAsyncWithHttpStatusRequestTags {
	s.Key = &v
	return s
}


type ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder struct {
	s *ElbOrderCreateAsyncWithHttpStatusRequestTags
}

func NewElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder() *ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder {
	s := &ElbOrderCreateAsyncWithHttpStatusRequestTags{}
	b := &ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder{s: s}
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder) Value(v string) *ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder {
	b.s.Value = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder) Key(v string) *ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder {
	b.s.Key = &v
	return b
}

func (b *ElbOrderCreateAsyncWithHttpStatusRequestTagsBuilder) Build() *ElbOrderCreateAsyncWithHttpStatusRequestTags {
	return b.s
}


