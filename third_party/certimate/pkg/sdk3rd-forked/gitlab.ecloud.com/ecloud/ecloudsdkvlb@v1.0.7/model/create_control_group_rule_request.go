// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateControlGroupRuleRequest struct {


	CreateControlGroupRuleBody *CreateControlGroupRuleBody `json:"createControlGroupRuleBody,omitempty"`
}

func (s CreateControlGroupRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateControlGroupRuleRequest) GoString() string {
	return s.String()
}

func (s CreateControlGroupRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateControlGroupRuleRequest) SetCreateControlGroupRuleBody(v *CreateControlGroupRuleBody) *CreateControlGroupRuleRequest {
	s.CreateControlGroupRuleBody = v
	return s
}


type CreateControlGroupRuleRequestBuilder struct {
	s *CreateControlGroupRuleRequest
}

func NewCreateControlGroupRuleRequestBuilder() *CreateControlGroupRuleRequestBuilder {
	s := &CreateControlGroupRuleRequest{}
	b := &CreateControlGroupRuleRequestBuilder{s: s}
	return b
}

func (b *CreateControlGroupRuleRequestBuilder) CreateControlGroupRuleBody(v *CreateControlGroupRuleBody) *CreateControlGroupRuleRequestBuilder {
	b.s.CreateControlGroupRuleBody = v
	return b
}

func (b *CreateControlGroupRuleRequestBuilder) Build() *CreateControlGroupRuleRequest {
	return b.s
}


