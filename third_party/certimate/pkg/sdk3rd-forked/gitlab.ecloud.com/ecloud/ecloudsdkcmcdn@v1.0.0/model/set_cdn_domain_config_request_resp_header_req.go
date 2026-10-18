// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestRespHeaderReq struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
    // header内容
	Detail *[]SetCdnDomainConfigRequestDetail `json:"detail,omitempty"`
}

func (s SetCdnDomainConfigRequestRespHeaderReq) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestRespHeaderReq) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestRespHeaderReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestRespHeaderReq) SetEnable(v bool) *SetCdnDomainConfigRequestRespHeaderReq {
	s.Enable = &v
	return s
}

func (s *SetCdnDomainConfigRequestRespHeaderReq) SetDetail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestRespHeaderReq {
	s.Detail = &v
	return s
}


type SetCdnDomainConfigRequestRespHeaderReqBuilder struct {
	s *SetCdnDomainConfigRequestRespHeaderReq
}

func NewSetCdnDomainConfigRequestRespHeaderReqBuilder() *SetCdnDomainConfigRequestRespHeaderReqBuilder {
	s := &SetCdnDomainConfigRequestRespHeaderReq{}
	b := &SetCdnDomainConfigRequestRespHeaderReqBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestRespHeaderReqBuilder) Enable(v bool) *SetCdnDomainConfigRequestRespHeaderReqBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestRespHeaderReqBuilder) Detail(v []SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestRespHeaderReqBuilder {
	b.s.Detail = &v
	return b
}

func (b *SetCdnDomainConfigRequestRespHeaderReqBuilder) Build() *SetCdnDomainConfigRequestRespHeaderReq {
	return b.s
}


