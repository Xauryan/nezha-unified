// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupRequest struct {


	ListControlGroupQuery *ListControlGroupQuery `json:"listControlGroupQuery,omitempty"`
}

func (s ListControlGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupRequest) GoString() string {
	return s.String()
}

func (s ListControlGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupRequest) SetListControlGroupQuery(v *ListControlGroupQuery) *ListControlGroupRequest {
	s.ListControlGroupQuery = v
	return s
}


type ListControlGroupRequestBuilder struct {
	s *ListControlGroupRequest
}

func NewListControlGroupRequestBuilder() *ListControlGroupRequestBuilder {
	s := &ListControlGroupRequest{}
	b := &ListControlGroupRequestBuilder{s: s}
	return b
}

func (b *ListControlGroupRequestBuilder) ListControlGroupQuery(v *ListControlGroupQuery) *ListControlGroupRequestBuilder {
	b.s.ListControlGroupQuery = v
	return b
}

func (b *ListControlGroupRequestBuilder) Build() *ListControlGroupRequest {
	return b.s
}


