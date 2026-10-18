// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadbalanceByTagRequest struct {


	ListLoadbalanceByTagQuery *ListLoadbalanceByTagQuery `json:"listLoadbalanceByTagQuery,omitempty"`
}

func (s ListLoadbalanceByTagRequest) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceByTagRequest) GoString() string {
	return s.String()
}

func (s ListLoadbalanceByTagRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceByTagRequest) SetListLoadbalanceByTagQuery(v *ListLoadbalanceByTagQuery) *ListLoadbalanceByTagRequest {
	s.ListLoadbalanceByTagQuery = v
	return s
}


type ListLoadbalanceByTagRequestBuilder struct {
	s *ListLoadbalanceByTagRequest
}

func NewListLoadbalanceByTagRequestBuilder() *ListLoadbalanceByTagRequestBuilder {
	s := &ListLoadbalanceByTagRequest{}
	b := &ListLoadbalanceByTagRequestBuilder{s: s}
	return b
}

func (b *ListLoadbalanceByTagRequestBuilder) ListLoadbalanceByTagQuery(v *ListLoadbalanceByTagQuery) *ListLoadbalanceByTagRequestBuilder {
	b.s.ListLoadbalanceByTagQuery = v
	return b
}

func (b *ListLoadbalanceByTagRequestBuilder) Build() *ListLoadbalanceByTagRequest {
	return b.s
}


