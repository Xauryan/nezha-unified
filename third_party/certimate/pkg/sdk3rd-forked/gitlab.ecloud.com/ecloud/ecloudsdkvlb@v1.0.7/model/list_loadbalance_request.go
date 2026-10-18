// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListLoadbalanceRequest struct {


	ListLoadbalanceQuery *ListLoadbalanceQuery `json:"listLoadbalanceQuery,omitempty"`
}

func (s ListLoadbalanceRequest) String() string {
	return utils.Beautify(s)
}

func (s ListLoadbalanceRequest) GoString() string {
	return s.String()
}

func (s ListLoadbalanceRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListLoadbalanceRequest) SetListLoadbalanceQuery(v *ListLoadbalanceQuery) *ListLoadbalanceRequest {
	s.ListLoadbalanceQuery = v
	return s
}


type ListLoadbalanceRequestBuilder struct {
	s *ListLoadbalanceRequest
}

func NewListLoadbalanceRequestBuilder() *ListLoadbalanceRequestBuilder {
	s := &ListLoadbalanceRequest{}
	b := &ListLoadbalanceRequestBuilder{s: s}
	return b
}

func (b *ListLoadbalanceRequestBuilder) ListLoadbalanceQuery(v *ListLoadbalanceQuery) *ListLoadbalanceRequestBuilder {
	b.s.ListLoadbalanceQuery = v
	return b
}

func (b *ListLoadbalanceRequestBuilder) Build() *ListLoadbalanceRequest {
	return b.s
}


