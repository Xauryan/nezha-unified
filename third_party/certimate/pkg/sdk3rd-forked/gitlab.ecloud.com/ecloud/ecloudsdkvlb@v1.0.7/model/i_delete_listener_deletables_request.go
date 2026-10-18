// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type IDeleteListenerDeletablesRequest struct {


	IDeleteListenerDeletablesHeader *IDeleteListenerDeletablesHeader `json:"iDeleteListenerDeletablesHeader,omitempty"`

	IDeleteListenerDeletablesBody *IDeleteListenerDeletablesBody `json:"iDeleteListenerDeletablesBody,omitempty"`
}

func (s IDeleteListenerDeletablesRequest) String() string {
	return utils.Beautify(s)
}

func (s IDeleteListenerDeletablesRequest) GoString() string {
	return s.String()
}

func (s IDeleteListenerDeletablesRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *IDeleteListenerDeletablesRequest) SetIDeleteListenerDeletablesHeader(v *IDeleteListenerDeletablesHeader) *IDeleteListenerDeletablesRequest {
	s.IDeleteListenerDeletablesHeader = v
	return s
}

func (s *IDeleteListenerDeletablesRequest) SetIDeleteListenerDeletablesBody(v *IDeleteListenerDeletablesBody) *IDeleteListenerDeletablesRequest {
	s.IDeleteListenerDeletablesBody = v
	return s
}


type IDeleteListenerDeletablesRequestBuilder struct {
	s *IDeleteListenerDeletablesRequest
}

func NewIDeleteListenerDeletablesRequestBuilder() *IDeleteListenerDeletablesRequestBuilder {
	s := &IDeleteListenerDeletablesRequest{}
	b := &IDeleteListenerDeletablesRequestBuilder{s: s}
	return b
}

func (b *IDeleteListenerDeletablesRequestBuilder) IDeleteListenerDeletablesHeader(v *IDeleteListenerDeletablesHeader) *IDeleteListenerDeletablesRequestBuilder {
	b.s.IDeleteListenerDeletablesHeader = v
	return b
}

func (b *IDeleteListenerDeletablesRequestBuilder) IDeleteListenerDeletablesBody(v *IDeleteListenerDeletablesBody) *IDeleteListenerDeletablesRequestBuilder {
	b.s.IDeleteListenerDeletablesBody = v
	return b
}

func (b *IDeleteListenerDeletablesRequestBuilder) Build() *IDeleteListenerDeletablesRequest {
	return b.s
}


