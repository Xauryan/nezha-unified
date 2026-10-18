// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CheckCdnFromApiPathProductTypeEnum string

// List of ProductType
const (
    CheckCdnFromApiPathProductTypeEnumEcdn CheckCdnFromApiPathProductTypeEnum = "ECDN"
)

type CheckCdnFromApiPath struct {
    position.Path
    // 产品类型
	ProductType *CheckCdnFromApiPathProductTypeEnum `json:"productType,omitempty"`
}

func (s CheckCdnFromApiPath) String() string {
	return utils.Beautify(s)
}

func (s CheckCdnFromApiPath) GoString() string {
	return s.String()
}

func (s CheckCdnFromApiPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CheckCdnFromApiPath) SetProductType(v CheckCdnFromApiPathProductTypeEnum) *CheckCdnFromApiPath {
	s.ProductType = &v
	return s
}


type CheckCdnFromApiPathBuilder struct {
	s *CheckCdnFromApiPath
}

func NewCheckCdnFromApiPathBuilder() *CheckCdnFromApiPathBuilder {
	s := &CheckCdnFromApiPath{}
	b := &CheckCdnFromApiPathBuilder{s: s}
	return b
}

func (b *CheckCdnFromApiPathBuilder) ProductType(v CheckCdnFromApiPathProductTypeEnum) *CheckCdnFromApiPathBuilder {
	b.s.ProductType = &v
	return b
}

func (b *CheckCdnFromApiPathBuilder) Build() *CheckCdnFromApiPath {
	return b.s
}


