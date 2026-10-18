// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestUrlParameterReq struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
    // 过滤参数内容
	Detail *SetCdnDomainConfigRequestDetail `json:"detail,omitempty"`
}

func (s SetCdnDomainConfigRequestUrlParameterReq) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestUrlParameterReq) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestUrlParameterReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestUrlParameterReq) SetEnable(v bool) *SetCdnDomainConfigRequestUrlParameterReq {
	s.Enable = &v
	return s
}

func (s *SetCdnDomainConfigRequestUrlParameterReq) SetDetail(v *SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestUrlParameterReq {
	s.Detail = v
	return s
}


type SetCdnDomainConfigRequestUrlParameterReqBuilder struct {
	s *SetCdnDomainConfigRequestUrlParameterReq
}

func NewSetCdnDomainConfigRequestUrlParameterReqBuilder() *SetCdnDomainConfigRequestUrlParameterReqBuilder {
	s := &SetCdnDomainConfigRequestUrlParameterReq{}
	b := &SetCdnDomainConfigRequestUrlParameterReqBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestUrlParameterReqBuilder) Enable(v bool) *SetCdnDomainConfigRequestUrlParameterReqBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestUrlParameterReqBuilder) Detail(v *SetCdnDomainConfigRequestDetail) *SetCdnDomainConfigRequestUrlParameterReqBuilder {
	b.s.Detail = v
	return b
}

func (b *SetCdnDomainConfigRequestUrlParameterReqBuilder) Build() *SetCdnDomainConfigRequestUrlParameterReq {
	return b.s
}


