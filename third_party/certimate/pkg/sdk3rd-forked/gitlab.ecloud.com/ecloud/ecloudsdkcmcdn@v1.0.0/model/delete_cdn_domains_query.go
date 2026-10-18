// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteCdnDomainsQuery struct {
    position.Query
    // 域名Id列表
	DomainIds *string `json:"domainIds,omitempty"`
}

func (s DeleteCdnDomainsQuery) String() string {
	return utils.Beautify(s)
}

func (s DeleteCdnDomainsQuery) GoString() string {
	return s.String()
}

func (s DeleteCdnDomainsQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteCdnDomainsQuery) SetDomainIds(v string) *DeleteCdnDomainsQuery {
	s.DomainIds = &v
	return s
}


type DeleteCdnDomainsQueryBuilder struct {
	s *DeleteCdnDomainsQuery
}

func NewDeleteCdnDomainsQueryBuilder() *DeleteCdnDomainsQueryBuilder {
	s := &DeleteCdnDomainsQuery{}
	b := &DeleteCdnDomainsQueryBuilder{s: s}
	return b
}

func (b *DeleteCdnDomainsQueryBuilder) DomainIds(v string) *DeleteCdnDomainsQueryBuilder {
	b.s.DomainIds = &v
	return b
}

func (b *DeleteCdnDomainsQueryBuilder) Build() *DeleteCdnDomainsQuery {
	return b.s
}


