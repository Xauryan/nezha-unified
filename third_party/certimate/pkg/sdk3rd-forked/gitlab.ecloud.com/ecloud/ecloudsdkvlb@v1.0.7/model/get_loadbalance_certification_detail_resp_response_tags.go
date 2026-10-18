// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetLoadbalanceCertificationDetailRespResponseTags struct {

    // 标签ID
	TagId *string `json:"tagId,omitempty"`
    // 标签值
	TagValue *string `json:"tagValue,omitempty"`
    // 标签健
	TagKey *string `json:"tagKey,omitempty"`
}

func (s GetLoadbalanceCertificationDetailRespResponseTags) String() string {
	return utils.Beautify(s)
}

func (s GetLoadbalanceCertificationDetailRespResponseTags) GoString() string {
	return s.String()
}

func (s GetLoadbalanceCertificationDetailRespResponseTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetLoadbalanceCertificationDetailRespResponseTags) SetTagId(v string) *GetLoadbalanceCertificationDetailRespResponseTags {
	s.TagId = &v
	return s
}

func (s *GetLoadbalanceCertificationDetailRespResponseTags) SetTagValue(v string) *GetLoadbalanceCertificationDetailRespResponseTags {
	s.TagValue = &v
	return s
}

func (s *GetLoadbalanceCertificationDetailRespResponseTags) SetTagKey(v string) *GetLoadbalanceCertificationDetailRespResponseTags {
	s.TagKey = &v
	return s
}


type GetLoadbalanceCertificationDetailRespResponseTagsBuilder struct {
	s *GetLoadbalanceCertificationDetailRespResponseTags
}

func NewGetLoadbalanceCertificationDetailRespResponseTagsBuilder() *GetLoadbalanceCertificationDetailRespResponseTagsBuilder {
	s := &GetLoadbalanceCertificationDetailRespResponseTags{}
	b := &GetLoadbalanceCertificationDetailRespResponseTagsBuilder{s: s}
	return b
}

func (b *GetLoadbalanceCertificationDetailRespResponseTagsBuilder) TagId(v string) *GetLoadbalanceCertificationDetailRespResponseTagsBuilder {
	b.s.TagId = &v
	return b
}

func (b *GetLoadbalanceCertificationDetailRespResponseTagsBuilder) TagValue(v string) *GetLoadbalanceCertificationDetailRespResponseTagsBuilder {
	b.s.TagValue = &v
	return b
}

func (b *GetLoadbalanceCertificationDetailRespResponseTagsBuilder) TagKey(v string) *GetLoadbalanceCertificationDetailRespResponseTagsBuilder {
	b.s.TagKey = &v
	return b
}

func (b *GetLoadbalanceCertificationDetailRespResponseTagsBuilder) Build() *GetLoadbalanceCertificationDetailRespResponseTags {
	return b.s
}


