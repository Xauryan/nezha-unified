// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListNewServerPortsForMemberRequest struct {


	ListNewServerPortsForMemberQuery *ListNewServerPortsForMemberQuery `json:"listNewServerPortsForMemberQuery,omitempty"`
}

func (s ListNewServerPortsForMemberRequest) String() string {
	return utils.Beautify(s)
}

func (s ListNewServerPortsForMemberRequest) GoString() string {
	return s.String()
}

func (s ListNewServerPortsForMemberRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListNewServerPortsForMemberRequest) SetListNewServerPortsForMemberQuery(v *ListNewServerPortsForMemberQuery) *ListNewServerPortsForMemberRequest {
	s.ListNewServerPortsForMemberQuery = v
	return s
}


type ListNewServerPortsForMemberRequestBuilder struct {
	s *ListNewServerPortsForMemberRequest
}

func NewListNewServerPortsForMemberRequestBuilder() *ListNewServerPortsForMemberRequestBuilder {
	s := &ListNewServerPortsForMemberRequest{}
	b := &ListNewServerPortsForMemberRequestBuilder{s: s}
	return b
}

func (b *ListNewServerPortsForMemberRequestBuilder) ListNewServerPortsForMemberQuery(v *ListNewServerPortsForMemberQuery) *ListNewServerPortsForMemberRequestBuilder {
	b.s.ListNewServerPortsForMemberQuery = v
	return b
}

func (b *ListNewServerPortsForMemberRequestBuilder) Build() *ListNewServerPortsForMemberRequest {
	return b.s
}


