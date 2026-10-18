// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestReqHeaderReq struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
    // header内容
	Detail *[]SetCdnDomainConfigRequestDetail `json:"detail,omitempty"`
}

func (s SetCdnDomainConfigRequestReqHeaderReq) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestReqHeaderReq) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestReqHeaderReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestReqHeaderReq) SetEnable(v bool) *SetCdnDomainConfigRequestReqHeaderReq {
	s.Enable = &v
	return s
}

func (s *SetCdnDomainConfigRequestReqHeaderReq) SetDetail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestReqHeaderReq {
	s.Detail = &v
	return s
}


type SetCdnDomainConfigRequestReqHeaderReqBuilder struct {
	s *SetCdnDomainConfigRequestReqHeaderReq
}

func NewSetCdnDomainConfigRequestReqHeaderReqBuilder() *SetCdnDomainConfigRequestReqHeaderReqBuilder {
	s := &SetCdnDomainConfigRequestReqHeaderReq{}
	b := &SetCdnDomainConfigRequestReqHeaderReqBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestReqHeaderReqBuilder) Enable(v bool) *SetCdnDomainConfigRequestReqHeaderReqBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestReqHeaderReqBuilder) Detail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestReqHeaderReqBuilder {
	b.s.Detail = &v
	return b
}

func (b *SetCdnDomainConfigRequestReqHeaderReqBuilder) Build() *SetCdnDomainConfigRequestReqHeaderReq {
	return b.s
}


