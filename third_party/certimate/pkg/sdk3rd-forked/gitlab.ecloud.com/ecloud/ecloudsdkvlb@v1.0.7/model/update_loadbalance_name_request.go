// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateLoadbalanceNameRequest struct {


	UpdateLoadbalanceNameBody *UpdateLoadbalanceNameBody `json:"updateLoadbalanceNameBody,omitempty"`
}

func (s UpdateLoadbalanceNameRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateLoadbalanceNameRequest) GoString() string {
	return s.String()
}

func (s UpdateLoadbalanceNameRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateLoadbalanceNameRequest) SetUpdateLoadbalanceNameBody(v *UpdateLoadbalanceNameBody) *UpdateLoadbalanceNameRequest {
	s.UpdateLoadbalanceNameBody = v
	return s
}


type UpdateLoadbalanceNameRequestBuilder struct {
	s *UpdateLoadbalanceNameRequest
}

func NewUpdateLoadbalanceNameRequestBuilder() *UpdateLoadbalanceNameRequestBuilder {
	s := &UpdateLoadbalanceNameRequest{}
	b := &UpdateLoadbalanceNameRequestBuilder{s: s}
	return b
}

func (b *UpdateLoadbalanceNameRequestBuilder) UpdateLoadbalanceNameBody(v *UpdateLoadbalanceNameBody) *UpdateLoadbalanceNameRequestBuilder {
	b.s.UpdateLoadbalanceNameBody = v
	return b
}

func (b *UpdateLoadbalanceNameRequestBuilder) Build() *UpdateLoadbalanceNameRequest {
	return b.s
}


