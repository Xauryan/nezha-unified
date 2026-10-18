// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteListenerAsyncRequest struct {


	DeleteListenerAsyncPath *DeleteListenerAsyncPath `json:"deleteListenerAsyncPath,omitempty"`
}

func (s DeleteListenerAsyncRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerAsyncRequest) GoString() string {
	return s.String()
}

func (s DeleteListenerAsyncRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerAsyncRequest) SetDeleteListenerAsyncPath(v *DeleteListenerAsyncPath) *DeleteListenerAsyncRequest {
	s.DeleteListenerAsyncPath = v
	return s
}


type DeleteListenerAsyncRequestBuilder struct {
	s *DeleteListenerAsyncRequest
}

func NewDeleteListenerAsyncRequestBuilder() *DeleteListenerAsyncRequestBuilder {
	s := &DeleteListenerAsyncRequest{}
	b := &DeleteListenerAsyncRequestBuilder{s: s}
	return b
}

func (b *DeleteListenerAsyncRequestBuilder) DeleteListenerAsyncPath(v *DeleteListenerAsyncPath) *DeleteListenerAsyncRequestBuilder {
	b.s.DeleteListenerAsyncPath = v
	return b
}

func (b *DeleteListenerAsyncRequestBuilder) Build() *DeleteListenerAsyncRequest {
	return b.s
}


