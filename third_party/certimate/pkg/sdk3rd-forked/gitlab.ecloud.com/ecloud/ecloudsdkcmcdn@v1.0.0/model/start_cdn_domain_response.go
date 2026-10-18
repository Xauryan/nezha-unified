// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartCdnDomainResponse struct {

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

func (s StartCdnDomainResponse) String() string {
	return utils.Beautify(s)
}

func (s StartCdnDomainResponse) GoString() string {
	return s.String()
}

func (s StartCdnDomainResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartCdnDomainResponse) SetRequestId(v string) *StartCdnDomainResponse {
	s.RequestId = &v
	return s
}

func (s *StartCdnDomainResponse) SetErrorMessage(v string) *StartCdnDomainResponse {
	s.ErrorMessage = &v
	return s
}

func (s *StartCdnDomainResponse) SetErrorCode(v string) *StartCdnDomainResponse {
	s.ErrorCode = &v
	return s
}

func (s *StartCdnDomainResponse) SetState(v string) *StartCdnDomainResponse {
	s.State = &v
	return s
}

func (s *StartCdnDomainResponse) SetBody(v bool) *StartCdnDomainResponse {
	s.Body = &v
	return s
}


type StartCdnDomainResponseBuilder struct {
	s *StartCdnDomainResponse
}

func NewStartCdnDomainResponseBuilder() *StartCdnDomainResponseBuilder {
	s := &StartCdnDomainResponse{}
	b := &StartCdnDomainResponseBuilder{s: s}
	return b
}

func (b *StartCdnDomainResponseBuilder) RequestId(v string) *StartCdnDomainResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *StartCdnDomainResponseBuilder) ErrorMessage(v string) *StartCdnDomainResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *StartCdnDomainResponseBuilder) ErrorCode(v string) *StartCdnDomainResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *StartCdnDomainResponseBuilder) State(v string) *StartCdnDomainResponseBuilder {
	b.s.State = &v
	return b
}

func (b *StartCdnDomainResponseBuilder) Body(v bool) *StartCdnDomainResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *StartCdnDomainResponseBuilder) Build() *StartCdnDomainResponse {
	return b.s
}


