// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CertificationBindTagsRequestTags struct {

    // 标签值
	Value *string `json:"value,omitempty"`
    // 标签键
	Key *string `json:"key,omitempty"`
}

func (s CertificationBindTagsRequestTags) String() string {
	return utils.Beautify(s)
}

func (s CertificationBindTagsRequestTags) GoString() string {
	return s.String()
}

func (s CertificationBindTagsRequestTags) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CertificationBindTagsRequestTags) SetValue(v string) *CertificationBindTagsRequestTags {
	s.Value = &v
	return s
}

func (s *CertificationBindTagsRequestTags) SetKey(v string) *CertificationBindTagsRequestTags {
	s.Key = &v
	return s
}


type CertificationBindTagsRequestTagsBuilder struct {
	s *CertificationBindTagsRequestTags
}

func NewCertificationBindTagsRequestTagsBuilder() *CertificationBindTagsRequestTagsBuilder {
	s := &CertificationBindTagsRequestTags{}
	b := &CertificationBindTagsRequestTagsBuilder{s: s}
	return b
}

func (b *CertificationBindTagsRequestTagsBuilder) Value(v string) *CertificationBindTagsRequestTagsBuilder {
	b.s.Value = &v
	return b
}

func (b *CertificationBindTagsRequestTagsBuilder) Key(v string) *CertificationBindTagsRequestTagsBuilder {
	b.s.Key = &v
	return b
}

func (b *CertificationBindTagsRequestTagsBuilder) Build() *CertificationBindTagsRequestTags {
	return b.s
}


