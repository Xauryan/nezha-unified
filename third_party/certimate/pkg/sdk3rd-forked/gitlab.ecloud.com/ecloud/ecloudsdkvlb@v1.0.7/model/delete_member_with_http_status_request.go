// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteMemberWithHttpStatusRequest struct {


	DeleteMemberWithHttpStatusHeader *DeleteMemberWithHttpStatusHeader `json:"deleteMemberWithHttpStatusHeader,omitempty"`

	DeleteMemberWithHttpStatusPath *DeleteMemberWithHttpStatusPath `json:"deleteMemberWithHttpStatusPath,omitempty"`
}

func (s DeleteMemberWithHttpStatusRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteMemberWithHttpStatusRequest) GoString() string {
	return s.String()
}

func (s DeleteMemberWithHttpStatusRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteMemberWithHttpStatusRequest) SetDeleteMemberWithHttpStatusHeader(v *DeleteMemberWithHttpStatusHeader) *DeleteMemberWithHttpStatusRequest {
	s.DeleteMemberWithHttpStatusHeader = v
	return s
}

func (s *DeleteMemberWithHttpStatusRequest) SetDeleteMemberWithHttpStatusPath(v *DeleteMemberWithHttpStatusPath) *DeleteMemberWithHttpStatusRequest {
	s.DeleteMemberWithHttpStatusPath = v
	return s
}


type DeleteMemberWithHttpStatusRequestBuilder struct {
	s *DeleteMemberWithHttpStatusRequest
}

func NewDeleteMemberWithHttpStatusRequestBuilder() *DeleteMemberWithHttpStatusRequestBuilder {
	s := &DeleteMemberWithHttpStatusRequest{}
	b := &DeleteMemberWithHttpStatusRequestBuilder{s: s}
	return b
}

func (b *DeleteMemberWithHttpStatusRequestBuilder) DeleteMemberWithHttpStatusHeader(v *DeleteMemberWithHttpStatusHeader) *DeleteMemberWithHttpStatusRequestBuilder {
	b.s.DeleteMemberWithHttpStatusHeader = v
	return b
}

func (b *DeleteMemberWithHttpStatusRequestBuilder) DeleteMemberWithHttpStatusPath(v *DeleteMemberWithHttpStatusPath) *DeleteMemberWithHttpStatusRequestBuilder {
	b.s.DeleteMemberWithHttpStatusPath = v
	return b
}

func (b *DeleteMemberWithHttpStatusRequestBuilder) Build() *DeleteMemberWithHttpStatusRequest {
	return b.s
}


