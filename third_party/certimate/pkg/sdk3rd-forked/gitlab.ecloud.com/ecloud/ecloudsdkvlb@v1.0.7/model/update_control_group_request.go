// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateControlGroupRequest struct {


	UpdateControlGroupBody *UpdateControlGroupBody `json:"updateControlGroupBody,omitempty"`
}

func (s UpdateControlGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateControlGroupRequest) GoString() string {
	return s.String()
}

func (s UpdateControlGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateControlGroupRequest) SetUpdateControlGroupBody(v *UpdateControlGroupBody) *UpdateControlGroupRequest {
	s.UpdateControlGroupBody = v
	return s
}


type UpdateControlGroupRequestBuilder struct {
	s *UpdateControlGroupRequest
}

func NewUpdateControlGroupRequestBuilder() *UpdateControlGroupRequestBuilder {
	s := &UpdateControlGroupRequest{}
	b := &UpdateControlGroupRequestBuilder{s: s}
	return b
}

func (b *UpdateControlGroupRequestBuilder) UpdateControlGroupBody(v *UpdateControlGroupBody) *UpdateControlGroupRequestBuilder {
	b.s.UpdateControlGroupBody = v
	return b
}

func (b *UpdateControlGroupRequestBuilder) Build() *UpdateControlGroupRequest {
	return b.s
}


