// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePoolRequest struct {


	DeletePoolPath *DeletePoolPath `json:"deletePoolPath,omitempty"`
}

func (s DeletePoolRequest) String() string {
	return utils.Beautify(s)
}

func (s DeletePoolRequest) GoString() string {
	return s.String()
}

func (s DeletePoolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePoolRequest) SetDeletePoolPath(v *DeletePoolPath) *DeletePoolRequest {
	s.DeletePoolPath = v
	return s
}


type DeletePoolRequestBuilder struct {
	s *DeletePoolRequest
}

func NewDeletePoolRequestBuilder() *DeletePoolRequestBuilder {
	s := &DeletePoolRequest{}
	b := &DeletePoolRequestBuilder{s: s}
	return b
}

func (b *DeletePoolRequestBuilder) DeletePoolPath(v *DeletePoolPath) *DeletePoolRequestBuilder {
	b.s.DeletePoolPath = v
	return b
}

func (b *DeletePoolRequestBuilder) Build() *DeletePoolRequest {
	return b.s
}


