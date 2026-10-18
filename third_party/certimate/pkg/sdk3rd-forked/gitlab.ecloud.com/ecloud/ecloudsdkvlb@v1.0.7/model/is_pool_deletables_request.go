// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type IsPoolDeletablesRequest struct {


	IsPoolDeletablesBody *IsPoolDeletablesBody `json:"isPoolDeletablesBody,omitempty"`
}

func (s IsPoolDeletablesRequest) String() string {
	return utils.Beautify(s)
}

func (s IsPoolDeletablesRequest) GoString() string {
	return s.String()
}

func (s IsPoolDeletablesRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IsPoolDeletablesRequest) SetIsPoolDeletablesBody(v *IsPoolDeletablesBody) *IsPoolDeletablesRequest {
	s.IsPoolDeletablesBody = v
	return s
}


type IsPoolDeletablesRequestBuilder struct {
	s *IsPoolDeletablesRequest
}

func NewIsPoolDeletablesRequestBuilder() *IsPoolDeletablesRequestBuilder {
	s := &IsPoolDeletablesRequest{}
	b := &IsPoolDeletablesRequestBuilder{s: s}
	return b
}

func (b *IsPoolDeletablesRequestBuilder) IsPoolDeletablesBody(v *IsPoolDeletablesBody) *IsPoolDeletablesRequestBuilder {
	b.s.IsPoolDeletablesBody = v
	return b
}

func (b *IsPoolDeletablesRequestBuilder) Build() *IsPoolDeletablesRequest {
	return b.s
}


