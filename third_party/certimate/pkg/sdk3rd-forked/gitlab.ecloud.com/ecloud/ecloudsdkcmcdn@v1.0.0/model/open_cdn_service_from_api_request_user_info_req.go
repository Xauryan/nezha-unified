// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum string

// List of IdType
const (
    OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnumIdCard OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum = "ID_CARD"
    OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnumPassport OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum = "PASSPORT"
    OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnumOfficerCard OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum = "OFFICER_CARD"
    OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnumMtps OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum = "MTPS"
)
type OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum string

// List of UnitNature
const (
    OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnumArmy OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum = "ARMY"
    OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnumGovernment OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum = "GOVERNMENT"
    OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnumInstitution OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum = "INSTITUTION"
    OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnumEnterprise OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum = "ENTERPRISE"
    OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnumPersonal OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum = "PERSONAL"
    OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnumComminity OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum = "COMMINITY"
    OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnumOther OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum = "OTHER"
)
type OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum string

// List of OfficerIdType
const (
    OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnumIdCard OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum = "ID_CARD"
    OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnumPassport OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum = "PASSPORT"
    OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnumOfficerCard OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum = "OFFICER_CARD"
    OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnumMtps OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum = "MTPS"
)

type OpenCdnServiceFromApiRequestUserInfoReq struct {

    // 网络信息安全责任人姓名
	OfficerEmployee *string `json:"officerEmployee,omitempty"`
    // 客户单位地址
	Address *string `json:"address,omitempty"`
    // 证件类型
	IdType *OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum `json:"idType,omitempty"`
    // 客户单位名称,企业名称，包括中文、英文字母、数字
	UnitName *string `json:"unitName,omitempty"`
    // 网络信息安全责任人证件号码
	OfficerIdNo *string `json:"officerIdNo,omitempty"`
    // 客户单位属性
	UnitNature *OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum `json:"unitNature,omitempty"`
    // 客户手机号
	Mobile *string `json:"mobile,omitempty"`
    // 证件号码
	IdNumber *string `json:"idNumber,omitempty"`
    // 客户邮箱
	Email *string `json:"email,omitempty"`
    // 网络信息安全责任人证件类型
	OfficerIdType *OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum `json:"officerIdType,omitempty"`
}

func (s OpenCdnServiceFromApiRequestUserInfoReq) String() string {
	return utils.Beautify(s)
}

func (s OpenCdnServiceFromApiRequestUserInfoReq) GoString() string {
	return s.String()
}

func (s OpenCdnServiceFromApiRequestUserInfoReq) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetOfficerEmployee(v string) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.OfficerEmployee = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetAddress(v string) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.Address = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetIdType(v OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.IdType = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetUnitName(v string) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.UnitName = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetOfficerIdNo(v string) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.OfficerIdNo = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetUnitNature(v OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.UnitNature = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetMobile(v string) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.Mobile = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetIdNumber(v string) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.IdNumber = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetEmail(v string) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.Email = &v
	return s
}

func (s *OpenCdnServiceFromApiRequestUserInfoReq) SetOfficerIdType(v OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum) *OpenCdnServiceFromApiRequestUserInfoReq {
	s.OfficerIdType = &v
	return s
}


type OpenCdnServiceFromApiRequestUserInfoReqBuilder struct {
	s *OpenCdnServiceFromApiRequestUserInfoReq
}

func NewOpenCdnServiceFromApiRequestUserInfoReqBuilder() *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	s := &OpenCdnServiceFromApiRequestUserInfoReq{}
	b := &OpenCdnServiceFromApiRequestUserInfoReqBuilder{s: s}
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) OfficerEmployee(v string) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.OfficerEmployee = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) Address(v string) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.Address = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) IdType(v OpenCdnServiceFromApiRequestUserInfoReqIdTypeEnum) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.IdType = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) UnitName(v string) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.UnitName = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) OfficerIdNo(v string) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.OfficerIdNo = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) UnitNature(v OpenCdnServiceFromApiRequestUserInfoReqUnitNatureEnum) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.UnitNature = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) Mobile(v string) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.Mobile = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) IdNumber(v string) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.IdNumber = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) Email(v string) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.Email = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) OfficerIdType(v OpenCdnServiceFromApiRequestUserInfoReqOfficerIdTypeEnum) *OpenCdnServiceFromApiRequestUserInfoReqBuilder {
	b.s.OfficerIdType = &v
	return b
}

func (b *OpenCdnServiceFromApiRequestUserInfoReqBuilder) Build() *OpenCdnServiceFromApiRequestUserInfoReq {
	return b.s
}


