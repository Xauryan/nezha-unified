// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeCdnDomainDetailResponseIpBlackList struct {

    // IP黑名单设置，上限为100条
	IpList *string `json:"ip_list,omitempty"`
}

func (s DescribeCdnDomainDetailResponseIpBlackList) String() string {
	return utils.Beautify(s)
}

func (s DescribeCdnDomainDetailResponseIpBlackList) GoString() string {
	return s.String()
}

func (s DescribeCdnDomainDetailResponseIpBlackList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeCdnDomainDetailResponseIpBlackList) SetIpList(v string) *DescribeCdnDomainDetailResponseIpBlackList {
	s.IpList = &v
	return s
}


type DescribeCdnDomainDetailResponseIpBlackListBuilder struct {
	s *DescribeCdnDomainDetailResponseIpBlackList
}

func NewDescribeCdnDomainDetailResponseIpBlackListBuilder() *DescribeCdnDomainDetailResponseIpBlackListBuilder {
	s := &DescribeCdnDomainDetailResponseIpBlackList{}
	b := &DescribeCdnDomainDetailResponseIpBlackListBuilder{s: s}
	return b
}

func (b *DescribeCdnDomainDetailResponseIpBlackListBuilder) IpList(v string) *DescribeCdnDomainDetailResponseIpBlackListBuilder {
	b.s.IpList = &v
	return b
}

func (b *DescribeCdnDomainDetailResponseIpBlackListBuilder) Build() *DescribeCdnDomainDetailResponseIpBlackList {
	return b.s
}


