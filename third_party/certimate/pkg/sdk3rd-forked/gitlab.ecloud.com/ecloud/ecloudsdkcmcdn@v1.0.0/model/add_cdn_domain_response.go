// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddCdnDomainResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *int32 `json:"body,omitempty"`
}

func (s AddCdnDomainResponse) String() string {
	return utils.Beautify(s)
}

func (s AddCdnDomainResponse) GoString() string {
	return s.String()
}

func (s AddCdnDomainResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddCdnDomainResponse) SetRequestId(v string) *AddCdnDomainResponse {
	s.RequestId = &v
	return s
}

func (s *AddCdnDomainResponse) SetErrorMessage(v string) *AddCdnDomainResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddCdnDomainResponse) SetErrorCode(v string) *AddCdnDomainResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddCdnDomainResponse) SetState(v string) *AddCdnDomainResponse {
	s.State = &v
	return s
}

func (s *AddCdnDomainResponse) SetBody(v int32) *AddCdnDomainResponse {
	s.Body = &v
	return s
}


type AddCdnDomainResponseBuilder struct {
	s *AddCdnDomainResponse
}

func NewAddCdnDomainResponseBuilder() *AddCdnDomainResponseBuilder {
	s := &AddCdnDomainResponse{}
	b := &AddCdnDomainResponseBuilder{s: s}
	return b
}

func (b *AddCdnDomainResponseBuilder) RequestId(v string) *AddCdnDomainResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddCdnDomainResponseBuilder) ErrorMessage(v string) *AddCdnDomainResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddCdnDomainResponseBuilder) ErrorCode(v string) *AddCdnDomainResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddCdnDomainResponseBuilder) State(v string) *AddCdnDomainResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddCdnDomainResponseBuilder) Body(v int32) *AddCdnDomainResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddCdnDomainResponseBuilder) Build() *AddCdnDomainResponse {
	return b.s
}


