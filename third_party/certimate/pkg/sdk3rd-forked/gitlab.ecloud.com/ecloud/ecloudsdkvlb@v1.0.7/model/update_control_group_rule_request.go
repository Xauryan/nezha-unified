// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateControlGroupRuleRequest struct {


	UpdateControlGroupRuleBody *UpdateControlGroupRuleBody `json:"updateControlGroupRuleBody,omitempty"`
}

func (s UpdateControlGroupRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateControlGroupRuleRequest) GoString() string {
	return s.String()
}

func (s UpdateControlGroupRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateControlGroupRuleRequest) SetUpdateControlGroupRuleBody(v *UpdateControlGroupRuleBody) *UpdateControlGroupRuleRequest {
	s.UpdateControlGroupRuleBody = v
	return s
}


type UpdateControlGroupRuleRequestBuilder struct {
	s *UpdateControlGroupRuleRequest
}

func NewUpdateControlGroupRuleRequestBuilder() *UpdateControlGroupRuleRequestBuilder {
	s := &UpdateControlGroupRuleRequest{}
	b := &UpdateControlGroupRuleRequestBuilder{s: s}
	return b
}

func (b *UpdateControlGroupRuleRequestBuilder) UpdateControlGroupRuleBody(v *UpdateControlGroupRuleBody) *UpdateControlGroupRuleRequestBuilder {
	b.s.UpdateControlGroupRuleBody = v
	return b
}

func (b *UpdateControlGroupRuleRequestBuilder) Build() *UpdateControlGroupRuleRequest {
	return b.s
}


