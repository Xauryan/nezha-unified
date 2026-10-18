// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StopCdnDomainQuery struct {
    position.Query
    // 域名Id列表
	DomainIds *string `json:"domainIds,omitempty"`
}

func (s StopCdnDomainQuery) String() string {
	return utils.Beautify(s)
}

func (s StopCdnDomainQuery) GoString() string {
	return s.String()
}

func (s StopCdnDomainQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StopCdnDomainQuery) SetDomainIds(v string) *StopCdnDomainQuery {
	s.DomainIds = &v
	return s
}


type StopCdnDomainQueryBuilder struct {
	s *StopCdnDomainQuery
}

func NewStopCdnDomainQueryBuilder() *StopCdnDomainQueryBuilder {
	s := &StopCdnDomainQuery{}
	b := &StopCdnDomainQueryBuilder{s: s}
	return b
}

func (b *StopCdnDomainQueryBuilder) DomainIds(v string) *StopCdnDomainQueryBuilder {
	b.s.DomainIds = &v
	return b
}

func (b *StopCdnDomainQueryBuilder) Build() *StopCdnDomainQuery {
	return b.s
}


