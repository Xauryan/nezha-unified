// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyCdnDomainResponse struct {

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

func (s ModifyCdnDomainResponse) String() string {
	return utils.Beautify(s)
}

func (s ModifyCdnDomainResponse) GoString() string {
	return s.String()
}

func (s ModifyCdnDomainResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyCdnDomainResponse) SetRequestId(v string) *ModifyCdnDomainResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyCdnDomainResponse) SetErrorMessage(v string) *ModifyCdnDomainResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyCdnDomainResponse) SetErrorCode(v string) *ModifyCdnDomainResponse {
	s.ErrorCode = &v
	return s
}

func (s *ModifyCdnDomainResponse) SetState(v string) *ModifyCdnDomainResponse {
	s.State = &v
	return s
}

func (s *ModifyCdnDomainResponse) SetBody(v bool) *ModifyCdnDomainResponse {
	s.Body = &v
	return s
}


type ModifyCdnDomainResponseBuilder struct {
	s *ModifyCdnDomainResponse
}

func NewModifyCdnDomainResponseBuilder() *ModifyCdnDomainResponseBuilder {
	s := &ModifyCdnDomainResponse{}
	b := &ModifyCdnDomainResponseBuilder{s: s}
	return b
}

func (b *ModifyCdnDomainResponseBuilder) RequestId(v string) *ModifyCdnDomainResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ModifyCdnDomainResponseBuilder) ErrorMessage(v string) *ModifyCdnDomainResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ModifyCdnDomainResponseBuilder) ErrorCode(v string) *ModifyCdnDomainResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ModifyCdnDomainResponseBuilder) State(v string) *ModifyCdnDomainResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ModifyCdnDomainResponseBuilder) Body(v bool) *ModifyCdnDomainResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ModifyCdnDomainResponseBuilder) Build() *ModifyCdnDomainResponse {
	return b.s
}


