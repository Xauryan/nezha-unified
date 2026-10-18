// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ValidateControlGroupRequest struct {


	ValidateControlGroupPath *ValidateControlGroupPath `json:"validateControlGroupPath,omitempty"`
}

func (s ValidateControlGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s ValidateControlGroupRequest) GoString() string {
	return s.String()
}

func (s ValidateControlGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ValidateControlGroupRequest) SetValidateControlGroupPath(v *ValidateControlGroupPath) *ValidateControlGroupRequest {
	s.ValidateControlGroupPath = v
	return s
}


type ValidateControlGroupRequestBuilder struct {
	s *ValidateControlGroupRequest
}

func NewValidateControlGroupRequestBuilder() *ValidateControlGroupRequestBuilder {
	s := &ValidateControlGroupRequest{}
	b := &ValidateControlGroupRequestBuilder{s: s}
	return b
}

func (b *ValidateControlGroupRequestBuilder) ValidateControlGroupPath(v *ValidateControlGroupPath) *ValidateControlGroupRequestBuilder {
	b.s.ValidateControlGroupPath = v
	return b
}

func (b *ValidateControlGroupRequestBuilder) Build() *ValidateControlGroupRequest {
	return b.s
}


