// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupRuleRequest struct {


	ListControlGroupRuleQuery *ListControlGroupRuleQuery `json:"listControlGroupRuleQuery,omitempty"`
}

func (s ListControlGroupRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupRuleRequest) GoString() string {
	return s.String()
}

func (s ListControlGroupRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupRuleRequest) SetListControlGroupRuleQuery(v *ListControlGroupRuleQuery) *ListControlGroupRuleRequest {
	s.ListControlGroupRuleQuery = v
	return s
}


type ListControlGroupRuleRequestBuilder struct {
	s *ListControlGroupRuleRequest
}

func NewListControlGroupRuleRequestBuilder() *ListControlGroupRuleRequestBuilder {
	s := &ListControlGroupRuleRequest{}
	b := &ListControlGroupRuleRequestBuilder{s: s}
	return b
}

func (b *ListControlGroupRuleRequestBuilder) ListControlGroupRuleQuery(v *ListControlGroupRuleQuery) *ListControlGroupRuleRequestBuilder {
	b.s.ListControlGroupRuleQuery = v
	return b
}

func (b *ListControlGroupRuleRequestBuilder) Build() *ListControlGroupRuleRequest {
	return b.s
}


