// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestReferrerAntiStealingLink struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
    // 防盗链配置内容
	Detail *SetCdnDomainConfigRequestDetail `json:"detail,omitempty"`
}

func (s SetCdnDomainConfigRequestReferrerAntiStealingLink) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestReferrerAntiStealingLink) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestReferrerAntiStealingLink) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestReferrerAntiStealingLink) SetEnable(v bool) *SetCdnDomainConfigRequestReferrerAntiStealingLink {
	s.Enable = &v
	return s
}

func (s *SetCdnDomainConfigRequestReferrerAntiStealingLink) SetDetail(v *SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestReferrerAntiStealingLink {
	s.Detail = v
	return s
}


type SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder struct {
	s *SetCdnDomainConfigRequestReferrerAntiStealingLink
}

func NewSetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder() *SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder {
	s := &SetCdnDomainConfigRequestReferrerAntiStealingLink{}
	b := &SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder) Enable(v bool) *SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder) Detail(v *SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder {
	b.s.Detail = v
	return b
}

func (b *SetCdnDomainConfigRequestReferrerAntiStealingLinkBuilder) Build() *SetCdnDomainConfigRequestReferrerAntiStealingLink {
	return b.s
}


