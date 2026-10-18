// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteControlGroupRequest struct {


	DeleteControlGroupPath *DeleteControlGroupPath `json:"deleteControlGroupPath,omitempty"`
}

func (s DeleteControlGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteControlGroupRequest) GoString() string {
	return s.String()
}

func (s DeleteControlGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteControlGroupRequest) SetDeleteControlGroupPath(v *DeleteControlGroupPath) *DeleteControlGroupRequest {
	s.DeleteControlGroupPath = v
	return s
}


type DeleteControlGroupRequestBuilder struct {
	s *DeleteControlGroupRequest
}

func NewDeleteControlGroupRequestBuilder() *DeleteControlGroupRequestBuilder {
	s := &DeleteControlGroupRequest{}
	b := &DeleteControlGroupRequestBuilder{s: s}
	return b
}

func (b *DeleteControlGroupRequestBuilder) DeleteControlGroupPath(v *DeleteControlGroupPath) *DeleteControlGroupRequestBuilder {
	b.s.DeleteControlGroupPath = v
	return b
}

func (b *DeleteControlGroupRequestBuilder) Build() *DeleteControlGroupRequest {
	return b.s
}


