// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AsyncCreatePolicyRequest struct {


	AsyncCreatePolicyBody *AsyncCreatePolicyBody `json:"asyncCreatePolicyBody,omitempty"`
}

func (s AsyncCreatePolicyRequest) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePolicyRequest) GoString() string {
	return s.String()
}

func (s AsyncCreatePolicyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePolicyRequest) SetAsyncCreatePolicyBody(v *AsyncCreatePolicyBody) *AsyncCreatePolicyRequest {
	s.AsyncCreatePolicyBody = v
	return s
}


type AsyncCreatePolicyRequestBuilder struct {
	s *AsyncCreatePolicyRequest
}

func NewAsyncCreatePolicyRequestBuilder() *AsyncCreatePolicyRequestBuilder {
	s := &AsyncCreatePolicyRequest{}
	b := &AsyncCreatePolicyRequestBuilder{s: s}
	return b
}

func (b *AsyncCreatePolicyRequestBuilder) AsyncCreatePolicyBody(v *AsyncCreatePolicyBody) *AsyncCreatePolicyRequestBuilder {
	b.s.AsyncCreatePolicyBody = v
	return b
}

func (b *AsyncCreatePolicyRequestBuilder) Build() *AsyncCreatePolicyRequest {
	return b.s
}


