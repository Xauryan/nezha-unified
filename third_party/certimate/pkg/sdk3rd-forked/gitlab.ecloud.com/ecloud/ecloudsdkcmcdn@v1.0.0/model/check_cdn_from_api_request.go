// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CheckCdnFromApiRequest struct {


	CheckCdnFromApiPath *CheckCdnFromApiPath `json:"checkCdnFromApiPath,omitempty"`
}

func (s CheckCdnFromApiRequest) String() string {
	return utils.Beautify(s)
}

func (s CheckCdnFromApiRequest) GoString() string {
	return s.String()
}

func (s CheckCdnFromApiRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CheckCdnFromApiRequest) SetCheckCdnFromApiPath(v *CheckCdnFromApiPath) *CheckCdnFromApiRequest {
	s.CheckCdnFromApiPath = v
	return s
}


type CheckCdnFromApiRequestBuilder struct {
	s *CheckCdnFromApiRequest
}

func NewCheckCdnFromApiRequestBuilder() *CheckCdnFromApiRequestBuilder {
	s := &CheckCdnFromApiRequest{}
	b := &CheckCdnFromApiRequestBuilder{s: s}
	return b
}

func (b *CheckCdnFromApiRequestBuilder) CheckCdnFromApiPath(v *CheckCdnFromApiPath) *CheckCdnFromApiRequestBuilder {
	b.s.CheckCdnFromApiPath = v
	return b
}

func (b *CheckCdnFromApiRequestBuilder) Build() *CheckCdnFromApiRequest {
	return b.s
}


