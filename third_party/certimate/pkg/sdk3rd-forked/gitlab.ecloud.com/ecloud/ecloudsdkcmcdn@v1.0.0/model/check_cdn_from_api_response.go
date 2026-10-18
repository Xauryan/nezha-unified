// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CheckCdnFromApiResponse struct {

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

func (s CheckCdnFromApiResponse) String() string {
	return utils.Beautify(s)
}

func (s CheckCdnFromApiResponse) GoString() string {
	return s.String()
}

func (s CheckCdnFromApiResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CheckCdnFromApiResponse) SetRequestId(v string) *CheckCdnFromApiResponse {
	s.RequestId = &v
	return s
}

func (s *CheckCdnFromApiResponse) SetErrorMessage(v string) *CheckCdnFromApiResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CheckCdnFromApiResponse) SetErrorCode(v string) *CheckCdnFromApiResponse {
	s.ErrorCode = &v
	return s
}

func (s *CheckCdnFromApiResponse) SetState(v string) *CheckCdnFromApiResponse {
	s.State = &v
	return s
}

func (s *CheckCdnFromApiResponse) SetBody(v bool) *CheckCdnFromApiResponse {
	s.Body = &v
	return s
}


type CheckCdnFromApiResponseBuilder struct {
	s *CheckCdnFromApiResponse
}

func NewCheckCdnFromApiResponseBuilder() *CheckCdnFromApiResponseBuilder {
	s := &CheckCdnFromApiResponse{}
	b := &CheckCdnFromApiResponseBuilder{s: s}
	return b
}

func (b *CheckCdnFromApiResponseBuilder) RequestId(v string) *CheckCdnFromApiResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CheckCdnFromApiResponseBuilder) ErrorMessage(v string) *CheckCdnFromApiResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CheckCdnFromApiResponseBuilder) ErrorCode(v string) *CheckCdnFromApiResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CheckCdnFromApiResponseBuilder) State(v string) *CheckCdnFromApiResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CheckCdnFromApiResponseBuilder) Body(v bool) *CheckCdnFromApiResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CheckCdnFromApiResponseBuilder) Build() *CheckCdnFromApiResponse {
	return b.s
}


