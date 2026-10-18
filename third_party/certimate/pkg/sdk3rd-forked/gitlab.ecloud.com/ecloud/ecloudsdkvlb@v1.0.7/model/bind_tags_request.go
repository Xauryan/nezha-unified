// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindTagsRequest struct {


	BindTagsBody *BindTagsBody `json:"bindTagsBody,omitempty"`
}

func (s BindTagsRequest) String() string {
	return utils.Beautify(s)
}

func (s BindTagsRequest) GoString() string {
	return s.String()
}

func (s BindTagsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindTagsRequest) SetBindTagsBody(v *BindTagsBody) *BindTagsRequest {
	s.BindTagsBody = v
	return s
}


type BindTagsRequestBuilder struct {
	s *BindTagsRequest
}

func NewBindTagsRequestBuilder() *BindTagsRequestBuilder {
	s := &BindTagsRequest{}
	b := &BindTagsRequestBuilder{s: s}
	return b
}

func (b *BindTagsRequestBuilder) BindTagsBody(v *BindTagsBody) *BindTagsRequestBuilder {
	b.s.BindTagsBody = v
	return b
}

func (b *BindTagsRequestBuilder) Build() *BindTagsRequest {
	return b.s
}


