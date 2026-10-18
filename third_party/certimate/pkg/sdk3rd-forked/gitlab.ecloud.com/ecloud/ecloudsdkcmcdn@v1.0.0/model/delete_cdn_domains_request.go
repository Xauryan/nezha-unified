// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteCdnDomainsRequest struct {


	DeleteCdnDomainsQuery *DeleteCdnDomainsQuery `json:"deleteCdnDomainsQuery,omitempty"`
}

func (s DeleteCdnDomainsRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteCdnDomainsRequest) GoString() string {
	return s.String()
}

func (s DeleteCdnDomainsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteCdnDomainsRequest) SetDeleteCdnDomainsQuery(v *DeleteCdnDomainsQuery) *DeleteCdnDomainsRequest {
	s.DeleteCdnDomainsQuery = v
	return s
}


type DeleteCdnDomainsRequestBuilder struct {
	s *DeleteCdnDomainsRequest
}

func NewDeleteCdnDomainsRequestBuilder() *DeleteCdnDomainsRequestBuilder {
	s := &DeleteCdnDomainsRequest{}
	b := &DeleteCdnDomainsRequestBuilder{s: s}
	return b
}

func (b *DeleteCdnDomainsRequestBuilder) DeleteCdnDomainsQuery(v *DeleteCdnDomainsQuery) *DeleteCdnDomainsRequestBuilder {
	b.s.DeleteCdnDomainsQuery = v
	return b
}

func (b *DeleteCdnDomainsRequestBuilder) Build() *DeleteCdnDomainsRequest {
	return b.s
}


