// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type OpenCdnServiceFromApiBodyCdnProductTypeEnum string

// List of CdnProductType
const (
    OpenCdnServiceFromApiBodyCdnProductTypeEnumEcdn OpenCdnServiceFromApiBodyCdnProductTypeEnum = "ECDN"
)
type OpenCdnServiceFromApiBodyChargeTypeEnum string

// List of ChargeType
const (
    OpenCdnServiceFromApiBodyChargeTypeEnumFlow OpenCdnServiceFromApiBodyChargeTypeEnum = "FLOW"
    OpenCdnServiceFromApiBodyChargeTypeEnumDayPeak OpenCdnServiceFromApiBodyChargeTypeEnum = "DAY_PEAK"
)
type OpenCdnServiceFromApiBodyServiceProviderEnum string

// List of ServiceProvider
const (
    OpenCdnServiceFromApiBodyServiceProviderEnumMobile OpenCdnServiceFromApiBodyServiceProviderEnum = "MOBILE"
    OpenCdnServiceFromApiBodyServiceProviderEnumOther OpenCdnServiceFromApiBodyServiceProviderEnum = "OTHER"
)

type OpenCdnServiceFromApiBody struct {
    position.Body
    // 用户信息
	UserInfoReq *OpenCdnServiceFromApiRequestUserInfoReq `json:"userInfoReq,omitempty"`
    // 产品类型
	CdnProductType *OpenCdnServiceFromApiBodyCdnProductTypeEnum `json:"cdnProductType,omitempty"`
    // 计费类型，按流量计费：FLOW，按日峰值计费：DAY_PEAK
	ChargeType *OpenCdnServiceFromApiBodyChargeTypeEnum `json:"chargeType,omitempty"`
    // 服务提供厂商,移动网：MOBILE，移动电信联通三网：OTHER
	ServiceProvider *OpenCdnServiceFromApiBodyServiceProviderEnum `json:"serviceProvider,omitempty"`
}

func (s OpenCdnServiceFromApiBody) String() string {
	return utils.Beautify(s)
}

func (s OpenCdnServiceFromApiBody) GoString() string {
	return s.String()
}

func (s OpenCdnServiceFromApiBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OpenCdnServiceFromApiBody) SetUserInfoReq(v *OpenCdnServiceFromApiRequestUserInfoReq) *OpenCdnServiceFromApiBody {
	s.UserInfoReq = v
	return s
}

func (s *OpenCdnServiceFromApiBody) SetCdnProductType(v OpenCdnServiceFromApiBodyCdnProductTypeEnum) *OpenCdnServiceFromApiBody {
	s.CdnProductType = &v
	return s
}

func (s *OpenCdnServiceFromApiBody) SetChargeType(v OpenCdnServiceFromApiBodyChargeTypeEnum) *OpenCdnServiceFromApiBody {
	s.ChargeType = &v
	return s
}

func (s *OpenCdnServiceFromApiBody) SetServiceProvider(v OpenCdnServiceFromApiBodyServiceProviderEnum) *OpenCdnServiceFromApiBody {
	s.ServiceProvider = &v
	return s
}


type OpenCdnServiceFromApiBodyBuilder struct {
	s *OpenCdnServiceFromApiBody
}

func NewOpenCdnServiceFromApiBodyBuilder() *OpenCdnServiceFromApiBodyBuilder {
	s := &OpenCdnServiceFromApiBody{}
	b := &OpenCdnServiceFromApiBodyBuilder{s: s}
	return b
}

func (b *OpenCdnServiceFromApiBodyBuilder) UserInfoReq(v *OpenCdnServiceFromApiRequestUserInfoReq) *OpenCdnServiceFromApiBodyBuilder {
	b.s.UserInfoReq = v
	return b
}

func (b *OpenCdnServiceFromApiBodyBuilder) CdnProductType(v OpenCdnServiceFromApiBodyCdnProductTypeEnum) *OpenCdnServiceFromApiBodyBuilder {
	b.s.CdnProductType = &v
	return b
}

func (b *OpenCdnServiceFromApiBodyBuilder) ChargeType(v OpenCdnServiceFromApiBodyChargeTypeEnum) *OpenCdnServiceFromApiBodyBuilder {
	b.s.ChargeType = &v
	return b
}

func (b *OpenCdnServiceFromApiBodyBuilder) ServiceProvider(v OpenCdnServiceFromApiBodyServiceProviderEnum) *OpenCdnServiceFromApiBodyBuilder {
	b.s.ServiceProvider = &v
	return b
}

func (b *OpenCdnServiceFromApiBodyBuilder) Build() *OpenCdnServiceFromApiBody {
	return b.s
}


