// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteListenerRequest struct {


	DeleteListenerPath *DeleteListenerPath `json:"deleteListenerPath,omitempty"`
}

func (s DeleteListenerRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerRequest) GoString() string {
	return s.String()
}

func (s DeleteListenerRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerRequest) SetDeleteListenerPath(v *DeleteListenerPath) *DeleteListenerRequest {
	s.DeleteListenerPath = v
	return s
}


type DeleteListenerRequestBuilder struct {
	s *DeleteListenerRequest
}

func NewDeleteListenerRequestBuilder() *DeleteListenerRequestBuilder {
	s := &DeleteListenerRequest{}
	b := &DeleteListenerRequestBuilder{s: s}
	return b
}

func (b *DeleteListenerRequestBuilder) DeleteListenerPath(v *DeleteListenerPath) *DeleteListenerRequestBuilder {
	b.s.DeleteListenerPath = v
	return b
}

func (b *DeleteListenerRequestBuilder) Build() *DeleteListenerRequest {
	return b.s
}


