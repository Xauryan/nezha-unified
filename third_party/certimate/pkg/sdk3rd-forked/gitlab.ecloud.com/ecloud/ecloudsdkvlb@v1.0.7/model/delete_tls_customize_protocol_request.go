// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteTLSCustomizeProtocolRequest struct {


	DeleteTLSCustomizeProtocolPath *DeleteTLSCustomizeProtocolPath `json:"deleteTLSCustomizeProtocolPath,omitempty"`
}

func (s DeleteTLSCustomizeProtocolRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteTLSCustomizeProtocolRequest) GoString() string {
	return s.String()
}

func (s DeleteTLSCustomizeProtocolRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteTLSCustomizeProtocolRequest) SetDeleteTLSCustomizeProtocolPath(v *DeleteTLSCustomizeProtocolPath) *DeleteTLSCustomizeProtocolRequest {
	s.DeleteTLSCustomizeProtocolPath = v
	return s
}


type DeleteTLSCustomizeProtocolRequestBuilder struct {
	s *DeleteTLSCustomizeProtocolRequest
}

func NewDeleteTLSCustomizeProtocolRequestBuilder() *DeleteTLSCustomizeProtocolRequestBuilder {
	s := &DeleteTLSCustomizeProtocolRequest{}
	b := &DeleteTLSCustomizeProtocolRequestBuilder{s: s}
	return b
}

func (b *DeleteTLSCustomizeProtocolRequestBuilder) DeleteTLSCustomizeProtocolPath(v *DeleteTLSCustomizeProtocolPath) *DeleteTLSCustomizeProtocolRequestBuilder {
	b.s.DeleteTLSCustomizeProtocolPath = v
	return b
}

func (b *DeleteTLSCustomizeProtocolRequestBuilder) Build() *DeleteTLSCustomizeProtocolRequest {
	return b.s
}


