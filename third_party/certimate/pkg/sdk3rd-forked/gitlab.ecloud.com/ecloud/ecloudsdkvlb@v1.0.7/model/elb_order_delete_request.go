// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderDeleteRequest struct {


	ElbOrderDeleteQuery *ElbOrderDeleteQuery `json:"elbOrderDeleteQuery,omitempty"`
}

func (s ElbOrderDeleteRequest) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderDeleteRequest) GoString() string {
	return s.String()
}

func (s ElbOrderDeleteRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderDeleteRequest) SetElbOrderDeleteQuery(v *ElbOrderDeleteQuery) *ElbOrderDeleteRequest {
	s.ElbOrderDeleteQuery = v
	return s
}


type ElbOrderDeleteRequestBuilder struct {
	s *ElbOrderDeleteRequest
}

func NewElbOrderDeleteRequestBuilder() *ElbOrderDeleteRequestBuilder {
	s := &ElbOrderDeleteRequest{}
	b := &ElbOrderDeleteRequestBuilder{s: s}
	return b
}

func (b *ElbOrderDeleteRequestBuilder) ElbOrderDeleteQuery(v *ElbOrderDeleteQuery) *ElbOrderDeleteRequestBuilder {
	b.s.ElbOrderDeleteQuery = v
	return b
}

func (b *ElbOrderDeleteRequestBuilder) Build() *ElbOrderDeleteRequest {
	return b.s
}


