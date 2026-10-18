// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadBalanceDetailRespResponseTags struct {

    // 标签ID
	TagId *string `json:"tagId,omitempty"`
    // 标签值
	TagValue *string `json:"tagValue,omitempty"`
    // 标签键
	TagKey *string `json:"tagKey,omitempty"`
}

func (s GetLoadBalanceDetailRespResponseTags) String() string {
	return utils.Beautify(s)
}

func (s GetLoadBalanceDetailRespResponseTags) GoString() string {
	return s.String()
}

func (s GetLoadBalanceDetailRespResponseTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadBalanceDetailRespResponseTags) SetTagId(v string) *GetLoadBalanceDetailRespResponseTags {
	s.TagId = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseTags) SetTagValue(v string) *GetLoadBalanceDetailRespResponseTags {
	s.TagValue = &v
	return s
}

func (s *GetLoadBalanceDetailRespResponseTags) SetTagKey(v string) *GetLoadBalanceDetailRespResponseTags {
	s.TagKey = &v
	return s
}


type GetLoadBalanceDetailRespResponseTagsBuilder struct {
	s *GetLoadBalanceDetailRespResponseTags
}

func NewGetLoadBalanceDetailRespResponseTagsBuilder() *GetLoadBalanceDetailRespResponseTagsBuilder {
	s := &GetLoadBalanceDetailRespResponseTags{}
	b := &GetLoadBalanceDetailRespResponseTagsBuilder{s: s}
	return b
}

func (b *GetLoadBalanceDetailRespResponseTagsBuilder) TagId(v string) *GetLoadBalanceDetailRespResponseTagsBuilder {
	b.s.TagId = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseTagsBuilder) TagValue(v string) *GetLoadBalanceDetailRespResponseTagsBuilder {
	b.s.TagValue = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseTagsBuilder) TagKey(v string) *GetLoadBalanceDetailRespResponseTagsBuilder {
	b.s.TagKey = &v
	return b
}

func (b *GetLoadBalanceDetailRespResponseTagsBuilder) Build() *GetLoadBalanceDetailRespResponseTags {
	return b.s
}


