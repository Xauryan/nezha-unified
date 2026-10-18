// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestErrorPageReq struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
    // error page内容
	Detail *[]SetCdnDomainConfigRequestDetail `json:"detail,omitempty"`
}

func (s SetCdnDomainConfigRequestErrorPageReq) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestErrorPageReq) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestErrorPageReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestErrorPageReq) SetEnable(v bool) *SetCdnDomainConfigRequestErrorPageReq {
	s.Enable = &v
	return s
}

func (s *SetCdnDomainConfigRequestErrorPageReq) SetDetail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestErrorPageReq {
	s.Detail = &v
	return s
}


type SetCdnDomainConfigRequestErrorPageReqBuilder struct {
	s *SetCdnDomainConfigRequestErrorPageReq
}

func NewSetCdnDomainConfigRequestErrorPageReqBuilder() *SetCdnDomainConfigRequestErrorPageReqBuilder {
	s := &SetCdnDomainConfigRequestErrorPageReq{}
	b := &SetCdnDomainConfigRequestErrorPageReqBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestErrorPageReqBuilder) Enable(v bool) *SetCdnDomainConfigRequestErrorPageReqBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestErrorPageReqBuilder) Detail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestErrorPageReqBuilder {
	b.s.Detail = &v
	return b
}

func (b *SetCdnDomainConfigRequestErrorPageReqBuilder) Build() *SetCdnDomainConfigRequestErrorPageReq {
	return b.s
}


