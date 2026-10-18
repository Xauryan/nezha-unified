// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestIpBlackList struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
    // IP黑名单设置，上限为100条
	IpList *string `json:"ip_list,omitempty"`
}

func (s SetCdnDomainConfigRequestIpBlackList) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestIpBlackList) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestIpBlackList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestIpBlackList) SetEnable(v bool) *SetCdnDomainConfigRequestIpBlackList {
	s.Enable = &v
	return s
}

func (s *SetCdnDomainConfigRequestIpBlackList) SetIpList(v string) *SetCdnDomainConfigRequestIpBlackList {
	s.IpList = &v
	return s
}


type SetCdnDomainConfigRequestIpBlackListBuilder struct {
	s *SetCdnDomainConfigRequestIpBlackList
}

func NewSetCdnDomainConfigRequestIpBlackListBuilder() *SetCdnDomainConfigRequestIpBlackListBuilder {
	s := &SetCdnDomainConfigRequestIpBlackList{}
	b := &SetCdnDomainConfigRequestIpBlackListBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestIpBlackListBuilder) Enable(v bool) *SetCdnDomainConfigRequestIpBlackListBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestIpBlackListBuilder) IpList(v string) *SetCdnDomainConfigRequestIpBlackListBuilder {
	b.s.IpList = &v
	return b
}

func (b *SetCdnDomainConfigRequestIpBlackListBuilder) Build() *SetCdnDomainConfigRequestIpBlackList {
	return b.s
}


