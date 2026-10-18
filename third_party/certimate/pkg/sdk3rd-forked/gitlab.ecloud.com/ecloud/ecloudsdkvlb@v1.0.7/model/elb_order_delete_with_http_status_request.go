// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderDeleteWithHttpStatusRequest struct {


	ElbOrderDeleteWithHttpStatusQuery *ElbOrderDeleteWithHttpStatusQuery `json:"elbOrderDeleteWithHttpStatusQuery,omitempty"`
}

func (s ElbOrderDeleteWithHttpStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderDeleteWithHttpStatusRequest) GoString() string {
	return s.String()
}

func (s ElbOrderDeleteWithHttpStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderDeleteWithHttpStatusRequest) SetElbOrderDeleteWithHttpStatusQuery(v *ElbOrderDeleteWithHttpStatusQuery) *ElbOrderDeleteWithHttpStatusRequest {
	s.ElbOrderDeleteWithHttpStatusQuery = v
	return s
}


type ElbOrderDeleteWithHttpStatusRequestBuilder struct {
	s *ElbOrderDeleteWithHttpStatusRequest
}

func NewElbOrderDeleteWithHttpStatusRequestBuilder() *ElbOrderDeleteWithHttpStatusRequestBuilder {
	s := &ElbOrderDeleteWithHttpStatusRequest{}
	b := &ElbOrderDeleteWithHttpStatusRequestBuilder{s: s}
	return b
}

func (b *ElbOrderDeleteWithHttpStatusRequestBuilder) ElbOrderDeleteWithHttpStatusQuery(v *ElbOrderDeleteWithHttpStatusQuery) *ElbOrderDeleteWithHttpStatusRequestBuilder {
	b.s.ElbOrderDeleteWithHttpStatusQuery = v
	return b
}

func (b *ElbOrderDeleteWithHttpStatusRequestBuilder) Build() *ElbOrderDeleteWithHttpStatusRequest {
	return b.s
}


