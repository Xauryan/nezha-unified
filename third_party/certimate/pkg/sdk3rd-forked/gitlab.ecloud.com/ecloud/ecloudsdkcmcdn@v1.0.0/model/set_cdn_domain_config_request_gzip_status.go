// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigRequestGzipStatus struct {

    // 是否启用
	Enable *bool `json:"enable,omitempty"`
}

func (s SetCdnDomainConfigRequestGzipStatus) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigRequestGzipStatus) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigRequestGzipStatus) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigRequestGzipStatus) SetEnable(v bool) *SetCdnDomainConfigRequestGzipStatus {
	s.Enable = &v
	return s
}


type SetCdnDomainConfigRequestGzipStatusBuilder struct {
	s *SetCdnDomainConfigRequestGzipStatus
}

func NewSetCdnDomainConfigRequestGzipStatusBuilder() *SetCdnDomainConfigRequestGzipStatusBuilder {
	s := &SetCdnDomainConfigRequestGzipStatus{}
	b := &SetCdnDomainConfigRequestGzipStatusBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigRequestGzipStatusBuilder) Enable(v bool) *SetCdnDomainConfigRequestGzipStatusBuilder {
	b.s.Enable = &v
	return b
}

func (b *SetCdnDomainConfigRequestGzipStatusBuilder) Build() *SetCdnDomainConfigRequestGzipStatus {
	return b.s
}


