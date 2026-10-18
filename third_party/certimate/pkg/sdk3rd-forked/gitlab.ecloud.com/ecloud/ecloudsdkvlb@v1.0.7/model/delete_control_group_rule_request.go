// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteControlGroupRuleRequest struct {


	DeleteControlGroupRuleBody *DeleteControlGroupRuleBody `json:"deleteControlGroupRuleBody,omitempty"`

	DeleteControlGroupRulePath *DeleteControlGroupRulePath `json:"deleteControlGroupRulePath,omitempty"`
}

func (s DeleteControlGroupRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteControlGroupRuleRequest) GoString() string {
	return s.String()
}

func (s DeleteControlGroupRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteControlGroupRuleRequest) SetDeleteControlGroupRuleBody(v *DeleteControlGroupRuleBody) *DeleteControlGroupRuleRequest {
	s.DeleteControlGroupRuleBody = v
	return s
}

func (s *DeleteControlGroupRuleRequest) SetDeleteControlGroupRulePath(v *DeleteControlGroupRulePath) *DeleteControlGroupRuleRequest {
	s.DeleteControlGroupRulePath = v
	return s
}


type DeleteControlGroupRuleRequestBuilder struct {
	s *DeleteControlGroupRuleRequest
}

func NewDeleteControlGroupRuleRequestBuilder() *DeleteControlGroupRuleRequestBuilder {
	s := &DeleteControlGroupRuleRequest{}
	b := &DeleteControlGroupRuleRequestBuilder{s: s}
	return b
}

func (b *DeleteControlGroupRuleRequestBuilder) DeleteControlGroupRuleBody(v *DeleteControlGroupRuleBody) *DeleteControlGroupRuleRequestBuilder {
	b.s.DeleteControlGroupRuleBody = v
	return b
}

func (b *DeleteControlGroupRuleRequestBuilder) DeleteControlGroupRulePath(v *DeleteControlGroupRulePath) *DeleteControlGroupRuleRequestBuilder {
	b.s.DeleteControlGroupRulePath = v
	return b
}

func (b *DeleteControlGroupRuleRequestBuilder) Build() *DeleteControlGroupRuleRequest {
	return b.s
}


