// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OpenCdnServiceFromApiResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *string `json:"body,omitempty"`
}

func (s OpenCdnServiceFromApiResponse) String() string {
	return utils.Beautify(s)
}

func (s OpenCdnServiceFromApiResponse) GoString() string {
	return s.String()
}

func (s OpenCdnServiceFromApiResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OpenCdnServiceFromApiResponse) SetRequestId(v string) *OpenCdnServiceFromApiResponse {
	s.RequestId = &v
	return s
}

func (s *OpenCdnServiceFromApiResponse) SetErrorMessage(v string) *OpenCdnServiceFromApiResponse {
	s.ErrorMessage = &v
	return s
}

func (s *OpenCdnServiceFromApiResponse) SetErrorCode(v string) *OpenCdnServiceFromApiResponse {
	s.ErrorCode = &v
	return s
}

func (s *OpenCdnServiceFromApiResponse) SetState(v string) *OpenCdnServiceFromApiResponse {
	s.State = &v
	return s
}

func (s *OpenCdnServiceFromApiResponse) SetBody(v string) *OpenCdnServiceFromApiResponse {
	s.Body = &v
	return s
}


type OpenCdnServiceFromApiResponseBuilder struct {
	s *OpenCdnServiceFromApiResponse
}

func NewOpenCdnServiceFromApiResponseBuilder() *OpenCdnServiceFromApiResponseBuilder {
	s := &OpenCdnServiceFromApiResponse{}
	b := &OpenCdnServiceFromApiResponseBuilder{s: s}
	return b
}

func (b *OpenCdnServiceFromApiResponseBuilder) RequestId(v string) *OpenCdnServiceFromApiResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *OpenCdnServiceFromApiResponseBuilder) ErrorMessage(v string) *OpenCdnServiceFromApiResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *OpenCdnServiceFromApiResponseBuilder) ErrorCode(v string) *OpenCdnServiceFromApiResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *OpenCdnServiceFromApiResponseBuilder) State(v string) *OpenCdnServiceFromApiResponseBuilder {
	b.s.State = &v
	return b
}

func (b *OpenCdnServiceFromApiResponseBuilder) Body(v string) *OpenCdnServiceFromApiResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *OpenCdnServiceFromApiResponseBuilder) Build() *OpenCdnServiceFromApiResponse {
	return b.s
}


