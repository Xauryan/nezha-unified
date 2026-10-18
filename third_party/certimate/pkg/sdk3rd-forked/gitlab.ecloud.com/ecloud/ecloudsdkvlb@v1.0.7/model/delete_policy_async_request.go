// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePolicyAsyncRequest struct {


	DeletePolicyAsyncPath *DeletePolicyAsyncPath `json:"deletePolicyAsyncPath,omitempty"`
}

func (s DeletePolicyAsyncRequest) String() string {
	return utils.Beautify(s)
}

func (s DeletePolicyAsyncRequest) GoString() string {
	return s.String()
}

func (s DeletePolicyAsyncRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePolicyAsyncRequest) SetDeletePolicyAsyncPath(v *DeletePolicyAsyncPath) *DeletePolicyAsyncRequest {
	s.DeletePolicyAsyncPath = v
	return s
}


type DeletePolicyAsyncRequestBuilder struct {
	s *DeletePolicyAsyncRequest
}

func NewDeletePolicyAsyncRequestBuilder() *DeletePolicyAsyncRequestBuilder {
	s := &DeletePolicyAsyncRequest{}
	b := &DeletePolicyAsyncRequestBuilder{s: s}
	return b
}

func (b *DeletePolicyAsyncRequestBuilder) DeletePolicyAsyncPath(v *DeletePolicyAsyncPath) *DeletePolicyAsyncRequestBuilder {
	b.s.DeletePolicyAsyncPath = v
	return b
}

func (b *DeletePolicyAsyncRequestBuilder) Build() *DeletePolicyAsyncRequest {
	return b.s
}


