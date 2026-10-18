// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupDetailRequest struct {


	ListControlGroupDetailPath *ListControlGroupDetailPath `json:"listControlGroupDetailPath,omitempty"`
}

func (s ListControlGroupDetailRequest) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupDetailRequest) GoString() string {
	return s.String()
}

func (s ListControlGroupDetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupDetailRequest) SetListControlGroupDetailPath(v *ListControlGroupDetailPath) *ListControlGroupDetailRequest {
	s.ListControlGroupDetailPath = v
	return s
}


type ListControlGroupDetailRequestBuilder struct {
	s *ListControlGroupDetailRequest
}

func NewListControlGroupDetailRequestBuilder() *ListControlGroupDetailRequestBuilder {
	s := &ListControlGroupDetailRequest{}
	b := &ListControlGroupDetailRequestBuilder{s: s}
	return b
}

func (b *ListControlGroupDetailRequestBuilder) ListControlGroupDetailPath(v *ListControlGroupDetailPath) *ListControlGroupDetailRequestBuilder {
	b.s.ListControlGroupDetailPath = v
	return b
}

func (b *ListControlGroupDetailRequestBuilder) Build() *ListControlGroupDetailRequest {
	return b.s
}


