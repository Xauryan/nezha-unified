// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteL7PolicyRequest struct {


	DeleteL7PolicyPath *DeleteL7PolicyPath `json:"deleteL7PolicyPath,omitempty"`
}

func (s DeleteL7PolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteL7PolicyRequest) GoString() string {
	return s.String()
}

func (s DeleteL7PolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteL7PolicyRequest) SetDeleteL7PolicyPath(v *DeleteL7PolicyPath) *DeleteL7PolicyRequest {
	s.DeleteL7PolicyPath = v
	return s
}


type DeleteL7PolicyRequestBuilder struct {
	s *DeleteL7PolicyRequest
}

func NewDeleteL7PolicyRequestBuilder() *DeleteL7PolicyRequestBuilder {
	s := &DeleteL7PolicyRequest{}
	b := &DeleteL7PolicyRequestBuilder{s: s}
	return b
}

func (b *DeleteL7PolicyRequestBuilder) DeleteL7PolicyPath(v *DeleteL7PolicyPath) *DeleteL7PolicyRequestBuilder {
	b.s.DeleteL7PolicyPath = v
	return b
}

func (b *DeleteL7PolicyRequestBuilder) Build() *DeleteL7PolicyRequest {
	return b.s
}


