// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartCdnDomainQuery struct {
    position.Query
    // 域名Id列表
	DomainIds *string `json:"domainIds,omitempty"`
}

func (s StartCdnDomainQuery) String() string {
	return utils.Beautify(s)
}

func (s StartCdnDomainQuery) GoString() string {
	return s.String()
}

func (s StartCdnDomainQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartCdnDomainQuery) SetDomainIds(v string) *StartCdnDomainQuery {
	s.DomainIds = &v
	return s
}


type StartCdnDomainQueryBuilder struct {
	s *StartCdnDomainQuery
}

func NewStartCdnDomainQueryBuilder() *StartCdnDomainQueryBuilder {
	s := &StartCdnDomainQuery{}
	b := &StartCdnDomainQueryBuilder{s: s}
	return b
}

func (b *StartCdnDomainQueryBuilder) DomainIds(v string) *StartCdnDomainQueryBuilder {
	b.s.DomainIds = &v
	return b
}

func (b *StartCdnDomainQueryBuilder) Build() *StartCdnDomainQuery {
	return b.s
}


