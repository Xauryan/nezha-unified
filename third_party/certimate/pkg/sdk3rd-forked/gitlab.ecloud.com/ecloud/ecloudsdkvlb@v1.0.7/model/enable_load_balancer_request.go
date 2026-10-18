// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EnableLoadBalancerRequest struct {


	EnableLoadBalancerQuery *EnableLoadBalancerQuery `json:"enableLoadBalancerQuery,omitempty"`

	EnableLoadBalancerPath *EnableLoadBalancerPath `json:"enableLoadBalancerPath,omitempty"`
}

func (s EnableLoadBalancerRequest) String() string {
	return utils.Beautify(s)
}

func (s EnableLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s EnableLoadBalancerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EnableLoadBalancerRequest) SetEnableLoadBalancerQuery(v *EnableLoadBalancerQuery) *EnableLoadBalancerRequest {
	s.EnableLoadBalancerQuery = v
	return s
}

func (s *EnableLoadBalancerRequest) SetEnableLoadBalancerPath(v *EnableLoadBalancerPath) *EnableLoadBalancerRequest {
	s.EnableLoadBalancerPath = v
	return s
}


type EnableLoadBalancerRequestBuilder struct {
	s *EnableLoadBalancerRequest
}

func NewEnableLoadBalancerRequestBuilder() *EnableLoadBalancerRequestBuilder {
	s := &EnableLoadBalancerRequest{}
	b := &EnableLoadBalancerRequestBuilder{s: s}
	return b
}

func (b *EnableLoadBalancerRequestBuilder) EnableLoadBalancerQuery(v *EnableLoadBalancerQuery) *EnableLoadBalancerRequestBuilder {
	b.s.EnableLoadBalancerQuery = v
	return b
}

func (b *EnableLoadBalancerRequestBuilder) EnableLoadBalancerPath(v *EnableLoadBalancerPath) *EnableLoadBalancerRequestBuilder {
	b.s.EnableLoadBalancerPath = v
	return b
}

func (b *EnableLoadBalancerRequestBuilder) Build() *EnableLoadBalancerRequest {
	return b.s
}


