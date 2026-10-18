// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetAsyncResultResponseStateEnum string

// List of State
const (
    GetAsyncResultResponseStateEnumOk GetAsyncResultResponseStateEnum = "OK"
    GetAsyncResultResponseStateEnumError GetAsyncResultResponseStateEnum = "ERROR"
    GetAsyncResultResponseStateEnumException GetAsyncResultResponseStateEnum = "EXCEPTION"
    GetAsyncResultResponseStateEnumAlarm GetAsyncResultResponseStateEnum = "ALARM"
    GetAsyncResultResponseStateEnumForbidden GetAsyncResultResponseStateEnum = "FORBIDDEN"
)

type GetAsyncResultResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码
	State *GetAsyncResultResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s GetAsyncResultResponse) String() string {
	return utils.Beautify(s)
}

func (s GetAsyncResultResponse) GoString() string {
	return s.String()
}

func (s GetAsyncResultResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAsyncResultResponse) SetAsyncKey(v string) *GetAsyncResultResponse {
	s.AsyncKey = &v
	return s
}

func (s *GetAsyncResultResponse) SetRequestId(v string) *GetAsyncResultResponse {
	s.RequestId = &v
	return s
}

func (s *GetAsyncResultResponse) SetErrorMessage(v string) *GetAsyncResultResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncResultResponse) SetErrorCode(v string) *GetAsyncResultResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncResultResponse) SetState(v GetAsyncResultResponseStateEnum) *GetAsyncResultResponse {
	s.State = &v
	return s
}

func (s *GetAsyncResultResponse) SetBody(v string) *GetAsyncResultResponse {
	s.Body = &v
	return s
}

func (s *GetAsyncResultResponse) SetErrorParams(v []string) *GetAsyncResultResponse {
	s.ErrorParams = v
	return s
}


type GetAsyncResultResponseBuilder struct {
	s *GetAsyncResultResponse
}

func NewGetAsyncResultResponseBuilder() *GetAsyncResultResponseBuilder {
	s := &GetAsyncResultResponse{}
	b := &GetAsyncResultResponseBuilder{s: s}
	return b
}

func (b *GetAsyncResultResponseBuilder) AsyncKey(v string) *GetAsyncResultResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *GetAsyncResultResponseBuilder) RequestId(v string) *GetAsyncResultResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetAsyncResultResponseBuilder) ErrorMessage(v string) *GetAsyncResultResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetAsyncResultResponseBuilder) ErrorCode(v string) *GetAsyncResultResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetAsyncResultResponseBuilder) State(v GetAsyncResultResponseStateEnum) *GetAsyncResultResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetAsyncResultResponseBuilder) Body(v string) *GetAsyncResultResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetAsyncResultResponseBuilder) ErrorParams(v []string) *GetAsyncResultResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *GetAsyncResultResponseBuilder) Build() *GetAsyncResultResponse {
	return b.s
}


