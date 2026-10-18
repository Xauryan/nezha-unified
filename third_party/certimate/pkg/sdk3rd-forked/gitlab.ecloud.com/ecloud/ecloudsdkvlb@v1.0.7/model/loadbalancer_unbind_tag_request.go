// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoadbalancerUnbindTagRequest struct {


	LoadbalancerUnbindTagPath *LoadbalancerUnbindTagPath `json:"loadbalancerUnbindTagPath,omitempty"`
}

func (s LoadbalancerUnbindTagRequest) String() string {
	return utils.Beautify(s)
}

func (s LoadbalancerUnbindTagRequest) GoString() string {
	return s.String()
}

func (s LoadbalancerUnbindTagRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoadbalancerUnbindTagRequest) SetLoadbalancerUnbindTagPath(v *LoadbalancerUnbindTagPath) *LoadbalancerUnbindTagRequest {
	s.LoadbalancerUnbindTagPath = v
	return s
}


type LoadbalancerUnbindTagRequestBuilder struct {
	s *LoadbalancerUnbindTagRequest
}

func NewLoadbalancerUnbindTagRequestBuilder() *LoadbalancerUnbindTagRequestBuilder {
	s := &LoadbalancerUnbindTagRequest{}
	b := &LoadbalancerUnbindTagRequestBuilder{s: s}
	return b
}

func (b *LoadbalancerUnbindTagRequestBuilder) LoadbalancerUnbindTagPath(v *LoadbalancerUnbindTagPath) *LoadbalancerUnbindTagRequestBuilder {
	b.s.LoadbalancerUnbindTagPath = v
	return b
}

func (b *LoadbalancerUnbindTagRequestBuilder) Build() *LoadbalancerUnbindTagRequest {
	return b.s
}


