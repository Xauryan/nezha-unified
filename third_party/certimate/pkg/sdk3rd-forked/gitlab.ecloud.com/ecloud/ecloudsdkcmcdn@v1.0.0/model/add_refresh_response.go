// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddRefreshResponse struct {

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

func (s AddRefreshResponse) String() string {
	return utils.Beautify(s)
}

func (s AddRefreshResponse) GoString() string {
	return s.String()
}

func (s AddRefreshResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddRefreshResponse) SetRequestId(v string) *AddRefreshResponse {
	s.RequestId = &v
	return s
}

func (s *AddRefreshResponse) SetErrorMessage(v string) *AddRefreshResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddRefreshResponse) SetErrorCode(v string) *AddRefreshResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddRefreshResponse) SetState(v string) *AddRefreshResponse {
	s.State = &v
	return s
}

func (s *AddRefreshResponse) SetBody(v bool) *AddRefreshResponse {
	s.Body = &v
	return s
}


type AddRefreshResponseBuilder struct {
	s *AddRefreshResponse
}

func NewAddRefreshResponseBuilder() *AddRefreshResponseBuilder {
	s := &AddRefreshResponse{}
	b := &AddRefreshResponseBuilder{s: s}
	return b
}

func (b *AddRefreshResponseBuilder) RequestId(v string) *AddRefreshResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddRefreshResponseBuilder) ErrorMessage(v string) *AddRefreshResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddRefreshResponseBuilder) ErrorCode(v string) *AddRefreshResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddRefreshResponseBuilder) State(v string) *AddRefreshResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddRefreshResponseBuilder) Body(v bool) *AddRefreshResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddRefreshResponseBuilder) Build() *AddRefreshResponse {
	return b.s
}


