// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ElbOrderRenewRequest struct {


	ElbOrderRenewBody *ElbOrderRenewBody `json:"elbOrderRenewBody,omitempty"`
}

func (s ElbOrderRenewRequest) String() string {
	return utils.Beautify(s)
}

func (s ElbOrderRenewRequest) GoString() string {
	return s.String()
}

func (s ElbOrderRenewRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ElbOrderRenewRequest) SetElbOrderRenewBody(v *ElbOrderRenewBody) *ElbOrderRenewRequest {
	s.ElbOrderRenewBody = v
	return s
}


type ElbOrderRenewRequestBuilder struct {
	s *ElbOrderRenewRequest
}

func NewElbOrderRenewRequestBuilder() *ElbOrderRenewRequestBuilder {
	s := &ElbOrderRenewRequest{}
	b := &ElbOrderRenewRequestBuilder{s: s}
	return b
}

func (b *ElbOrderRenewRequestBuilder) ElbOrderRenewBody(v *ElbOrderRenewBody) *ElbOrderRenewRequestBuilder {
	b.s.ElbOrderRenewBody = v
	return b
}

func (b *ElbOrderRenewRequestBuilder) Build() *ElbOrderRenewRequest {
	return b.s
}


