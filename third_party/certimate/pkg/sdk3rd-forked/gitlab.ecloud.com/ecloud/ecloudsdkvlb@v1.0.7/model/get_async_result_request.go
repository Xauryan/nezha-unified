// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAsyncResultRequest struct {


	GetAsyncResultQuery *GetAsyncResultQuery `json:"getAsyncResultQuery,omitempty"`
}

func (s GetAsyncResultRequest) String() string {
	return utils.Beautify(s)
}

func (s GetAsyncResultRequest) GoString() string {
	return s.String()
}

func (s GetAsyncResultRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAsyncResultRequest) SetGetAsyncResultQuery(v *GetAsyncResultQuery) *GetAsyncResultRequest {
	s.GetAsyncResultQuery = v
	return s
}


type GetAsyncResultRequestBuilder struct {
	s *GetAsyncResultRequest
}

func NewGetAsyncResultRequestBuilder() *GetAsyncResultRequestBuilder {
	s := &GetAsyncResultRequest{}
	b := &GetAsyncResultRequestBuilder{s: s}
	return b
}

func (b *GetAsyncResultRequestBuilder) GetAsyncResultQuery(v *GetAsyncResultQuery) *GetAsyncResultRequestBuilder {
	b.s.GetAsyncResultQuery = v
	return b
}

func (b *GetAsyncResultRequestBuilder) Build() *GetAsyncResultRequest {
	return b.s
}


