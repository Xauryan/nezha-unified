// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetCdnDomainConfigResponse struct {

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

func (s SetCdnDomainConfigResponse) String() string {
	return utils.Beautify(s)
}

func (s SetCdnDomainConfigResponse) GoString() string {
	return s.String()
}

func (s SetCdnDomainConfigResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetCdnDomainConfigResponse) SetRequestId(v string) *SetCdnDomainConfigResponse {
	s.RequestId = &v
	return s
}

func (s *SetCdnDomainConfigResponse) SetErrorMessage(v string) *SetCdnDomainConfigResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SetCdnDomainConfigResponse) SetErrorCode(v string) *SetCdnDomainConfigResponse {
	s.ErrorCode = &v
	return s
}

func (s *SetCdnDomainConfigResponse) SetState(v string) *SetCdnDomainConfigResponse {
	s.State = &v
	return s
}

func (s *SetCdnDomainConfigResponse) SetBody(v bool) *SetCdnDomainConfigResponse {
	s.Body = &v
	return s
}


type SetCdnDomainConfigResponseBuilder struct {
	s *SetCdnDomainConfigResponse
}

func NewSetCdnDomainConfigResponseBuilder() *SetCdnDomainConfigResponseBuilder {
	s := &SetCdnDomainConfigResponse{}
	b := &SetCdnDomainConfigResponseBuilder{s: s}
	return b
}

func (b *SetCdnDomainConfigResponseBuilder) RequestId(v string) *SetCdnDomainConfigResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SetCdnDomainConfigResponseBuilder) ErrorMessage(v string) *SetCdnDomainConfigResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SetCdnDomainConfigResponseBuilder) ErrorCode(v string) *SetCdnDomainConfigResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SetCdnDomainConfigResponseBuilder) State(v string) *SetCdnDomainConfigResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SetCdnDomainConfigResponseBuilder) Body(v bool) *SetCdnDomainConfigResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *SetCdnDomainConfigResponseBuilder) Build() *SetCdnDomainConfigResponse {
	return b.s
}


