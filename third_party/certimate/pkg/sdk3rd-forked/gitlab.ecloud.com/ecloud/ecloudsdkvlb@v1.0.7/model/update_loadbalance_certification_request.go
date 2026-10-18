// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadbalanceCertificationRequest struct {


	UpdateLoadbalanceCertificationBody *UpdateLoadbalanceCertificationBody `json:"updateLoadbalanceCertificationBody,omitempty"`
}

func (s UpdateLoadbalanceCertificationRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadbalanceCertificationRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadbalanceCertificationRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadbalanceCertificationRequest) SetUpdateLoadbalanceCertificationBody(v *UpdateLoadbalanceCertificationBody) *UpdateLoadbalanceCertificationRequest {
	s.UpdateLoadbalanceCertificationBody = v
	return s
}


type UpdateLoadbalanceCertificationRequestBuilder struct {
	s *UpdateLoadbalanceCertificationRequest
}

func NewUpdateLoadbalanceCertificationRequestBuilder() *UpdateLoadbalanceCertificationRequestBuilder {
	s := &UpdateLoadbalanceCertificationRequest{}
	b := &UpdateLoadbalanceCertificationRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadbalanceCertificationRequestBuilder) UpdateLoadbalanceCertificationBody(v *UpdateLoadbalanceCertificationBody) *UpdateLoadbalanceCertificationRequestBuilder {
	b.s.UpdateLoadbalanceCertificationBody = v
	return b
}

func (b *UpdateLoadbalanceCertificationRequestBuilder) Build() *UpdateLoadbalanceCertificationRequest {
	return b.s
}


