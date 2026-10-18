// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteListenerWithHttpStatusRequest struct {


	DeleteListenerWithHttpStatusHeader *DeleteListenerWithHttpStatusHeader `json:"deleteListenerWithHttpStatusHeader,omitempty"`

	DeleteListenerWithHttpStatusPath *DeleteListenerWithHttpStatusPath `json:"deleteListenerWithHttpStatusPath,omitempty"`
}

func (s DeleteListenerWithHttpStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerWithHttpStatusRequest) GoString() string {
	return s.String()
}

func (s DeleteListenerWithHttpStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerWithHttpStatusRequest) SetDeleteListenerWithHttpStatusHeader(v *DeleteListenerWithHttpStatusHeader) *DeleteListenerWithHttpStatusRequest {
	s.DeleteListenerWithHttpStatusHeader = v
	return s
}

func (s *DeleteListenerWithHttpStatusRequest) SetDeleteListenerWithHttpStatusPath(v *DeleteListenerWithHttpStatusPath) *DeleteListenerWithHttpStatusRequest {
	s.DeleteListenerWithHttpStatusPath = v
	return s
}


type DeleteListenerWithHttpStatusRequestBuilder struct {
	s *DeleteListenerWithHttpStatusRequest
}

func NewDeleteListenerWithHttpStatusRequestBuilder() *DeleteListenerWithHttpStatusRequestBuilder {
	s := &DeleteListenerWithHttpStatusRequest{}
	b := &DeleteListenerWithHttpStatusRequestBuilder{s: s}
	return b
}

func (b *DeleteListenerWithHttpStatusRequestBuilder) DeleteListenerWithHttpStatusHeader(v *DeleteListenerWithHttpStatusHeader) *DeleteListenerWithHttpStatusRequestBuilder {
	b.s.DeleteListenerWithHttpStatusHeader = v
	return b
}

func (b *DeleteListenerWithHttpStatusRequestBuilder) DeleteListenerWithHttpStatusPath(v *DeleteListenerWithHttpStatusPath) *DeleteListenerWithHttpStatusRequestBuilder {
	b.s.DeleteListenerWithHttpStatusPath = v
	return b
}

func (b *DeleteListenerWithHttpStatusRequestBuilder) Build() *DeleteListenerWithHttpStatusRequest {
	return b.s
}


