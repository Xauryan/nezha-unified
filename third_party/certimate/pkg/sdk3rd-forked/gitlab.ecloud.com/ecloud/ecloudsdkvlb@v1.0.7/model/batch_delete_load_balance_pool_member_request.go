// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteLoadBalancePoolMemberRequest struct {


	BatchDeleteLoadBalancePoolMemberPath *BatchDeleteLoadBalancePoolMemberPath `json:"batchDeleteLoadBalancePoolMemberPath,omitempty"`

	BatchDeleteLoadBalancePoolMemberBody *BatchDeleteLoadBalancePoolMemberBody `json:"batchDeleteLoadBalancePoolMemberBody,omitempty"`
}

func (s BatchDeleteLoadBalancePoolMemberRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteLoadBalancePoolMemberRequest) GoString() string {
	return s.String()
}

func (s BatchDeleteLoadBalancePoolMemberRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteLoadBalancePoolMemberRequest) SetBatchDeleteLoadBalancePoolMemberPath(v *BatchDeleteLoadBalancePoolMemberPath) *BatchDeleteLoadBalancePoolMemberRequest {
	s.BatchDeleteLoadBalancePoolMemberPath = v
	return s
}

func (s *BatchDeleteLoadBalancePoolMemberRequest) SetBatchDeleteLoadBalancePoolMemberBody(v *BatchDeleteLoadBalancePoolMemberBody) *BatchDeleteLoadBalancePoolMemberRequest {
	s.BatchDeleteLoadBalancePoolMemberBody = v
	return s
}


type BatchDeleteLoadBalancePoolMemberRequestBuilder struct {
	s *BatchDeleteLoadBalancePoolMemberRequest
}

func NewBatchDeleteLoadBalancePoolMemberRequestBuilder() *BatchDeleteLoadBalancePoolMemberRequestBuilder {
	s := &BatchDeleteLoadBalancePoolMemberRequest{}
	b := &BatchDeleteLoadBalancePoolMemberRequestBuilder{s: s}
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberRequestBuilder) BatchDeleteLoadBalancePoolMemberPath(v *BatchDeleteLoadBalancePoolMemberPath) *BatchDeleteLoadBalancePoolMemberRequestBuilder {
	b.s.BatchDeleteLoadBalancePoolMemberPath = v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberRequestBuilder) BatchDeleteLoadBalancePoolMemberBody(v *BatchDeleteLoadBalancePoolMemberBody) *BatchDeleteLoadBalancePoolMemberRequestBuilder {
	b.s.BatchDeleteLoadBalancePoolMemberBody = v
	return b
}

func (b *BatchDeleteLoadBalancePoolMemberRequestBuilder) Build() *BatchDeleteLoadBalancePoolMemberRequest {
	return b.s
}


