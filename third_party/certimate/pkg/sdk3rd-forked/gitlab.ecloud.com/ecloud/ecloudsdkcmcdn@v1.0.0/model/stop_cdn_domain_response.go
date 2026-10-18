// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StopCdnDomainResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *bool `json:"body,omitempty"`
}

func (s StopCdnDomainResponse) String() string {
	return utils.Beautify(s)
}

func (s StopCdnDomainResponse) GoString() string {
	return s.String()
}

func (s StopCdnDomainResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StopCdnDomainResponse) SetRequestId(v string) *StopCdnDomainResponse {
	s.RequestId = &v
	return s
}

func (s *StopCdnDomainResponse) SetErrorMessage(v string) *StopCdnDomainResponse {
	s.ErrorMessage = &v
	return s
}

func (s *StopCdnDomainResponse) SetErrorCode(v string) *StopCdnDomainResponse {
	s.ErrorCode = &v
	return s
}

func (s *StopCdnDomainResponse) SetState(v string) *StopCdnDomainResponse {
	s.State = &v
	return s
}

func (s *StopCdnDomainResponse) SetBody(v bool) *StopCdnDomainResponse {
	s.Body = &v
	return s
}


type StopCdnDomainResponseBuilder struct {
	s *StopCdnDomainResponse
}

func NewStopCdnDomainResponseBuilder() *StopCdnDomainResponseBuilder {
	s := &StopCdnDomainResponse{}
	b := &StopCdnDomainResponseBuilder{s: s}
	return b
}

func (b *StopCdnDomainResponseBuilder) RequestId(v string) *StopCdnDomainResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *StopCdnDomainResponseBuilder) ErrorMessage(v string) *StopCdnDomainResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *StopCdnDomainResponseBuilder) ErrorCode(v string) *StopCdnDomainResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *StopCdnDomainResponseBuilder) State(v string) *StopCdnDomainResponseBuilder {
	b.s.State = &v
	return b
}

func (b *StopCdnDomainResponseBuilder) Body(v bool) *StopCdnDomainResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *StopCdnDomainResponseBuilder) Build() *StopCdnDomainResponse {
	return b.s
}


