// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeCdnDomainDetailResponseRespHeaderDetails struct {

    // 参数类型
	Type *int32 `json:"type,omitempty"`
    // 取值
	Value *string `json:"value,omitempty"`
}

func (s DescribeCdnDomainDetailResponseRespHeaderDetails) String() string {
	return utils.Beautify(s)
}

func (s DescribeCdnDomainDetailResponseRespHeaderDetails) GoString() string {
	return s.String()
}

func (s DescribeCdnDomainDetailResponseRespHeaderDetails) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeCdnDomainDetailResponseRespHeaderDetails) SetType(v int32) *DescribeCdnDomainDetailResponseRespHeaderDetails {
	s.Type = &v
	return s
}

func (s *DescribeCdnDomainDetailResponseRespHeaderDetails) SetValue(v string) *DescribeCdnDomainDetailResponseRespHeaderDetails {
	s.Value = &v
	return s
}


type DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder struct {
	s *DescribeCdnDomainDetailResponseRespHeaderDetails
}

func NewDescribeCdnDomainDetailResponseRespHeaderDetailsBuilder() *DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder {
	s := &DescribeCdnDomainDetailResponseRespHeaderDetails{}
	b := &DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder{s: s}
	return b
}

func (b *DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder) Type(v int32) *DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder {
	b.s.Type = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder) Value(v string) *DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder {
	b.s.Value = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseRespHeaderDetailsBuilder) Build() *DescribeCdnDomainDetailResponseRespHeaderDetails {
	return b.s
}


