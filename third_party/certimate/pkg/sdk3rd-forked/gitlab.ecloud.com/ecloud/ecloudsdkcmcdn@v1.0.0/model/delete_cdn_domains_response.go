// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteCdnDomainsResponse struct {

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

func (s DeleteCdnDomainsResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteCdnDomainsResponse) GoString() string {
	return s.String()
}

func (s DeleteCdnDomainsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteCdnDomainsResponse) SetRequestId(v string) *DeleteCdnDomainsResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteCdnDomainsResponse) SetErrorMessage(v string) *DeleteCdnDomainsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteCdnDomainsResponse) SetErrorCode(v string) *DeleteCdnDomainsResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCdnDomainsResponse) SetState(v string) *DeleteCdnDomainsResponse {
	s.State = &v
	return s
}

func (s *DeleteCdnDomainsResponse) SetBody(v bool) *DeleteCdnDomainsResponse {
	s.Body = &v
	return s
}


type DeleteCdnDomainsResponseBuilder struct {
	s *DeleteCdnDomainsResponse
}

func NewDeleteCdnDomainsResponseBuilder() *DeleteCdnDomainsResponseBuilder {
	s := &DeleteCdnDomainsResponse{}
	b := &DeleteCdnDomainsResponseBuilder{s: s}
	return b
}

func (b *DeleteCdnDomainsResponseBuilder) RequestId(v string) *DeleteCdnDomainsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteCdnDomainsResponseBuilder) ErrorMessage(v string) *DeleteCdnDomainsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteCdnDomainsResponseBuilder) ErrorCode(v string) *DeleteCdnDomainsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteCdnDomainsResponseBuilder) State(v string) *DeleteCdnDomainsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteCdnDomainsResponseBuilder) Body(v bool) *DeleteCdnDomainsResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteCdnDomainsResponseBuilder) Build() *DeleteCdnDomainsResponse {
	return b.s
}


