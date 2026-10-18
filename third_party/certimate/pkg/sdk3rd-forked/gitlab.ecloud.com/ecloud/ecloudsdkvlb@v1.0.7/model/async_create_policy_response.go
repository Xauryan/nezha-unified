// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreatePolicyResponseStateEnum string

// List of State
const (
    AsyncCreatePolicyResponseStateEnumOk AsyncCreatePolicyResponseStateEnum = "OK"
    AsyncCreatePolicyResponseStateEnumError AsyncCreatePolicyResponseStateEnum = "ERROR"
    AsyncCreatePolicyResponseStateEnumException AsyncCreatePolicyResponseStateEnum = "EXCEPTION"
    AsyncCreatePolicyResponseStateEnumAlarm AsyncCreatePolicyResponseStateEnum = "ALARM"
    AsyncCreatePolicyResponseStateEnumForbidden AsyncCreatePolicyResponseStateEnum = "FORBIDDEN"
)

type AsyncCreatePolicyResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码 OK（成功）,ERROR（错误）,EXCEPTION（异常）,ALARM（警告）,FORBIDDEN（禁止访问）
	State *AsyncCreatePolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s AsyncCreatePolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePolicyResponse) GoString() string {
	return s.String()
}

func (s AsyncCreatePolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePolicyResponse) SetAsyncKey(v string) *AsyncCreatePolicyResponse {
	s.AsyncKey = &v
	return s
}

func (s *AsyncCreatePolicyResponse) SetRequestId(v string) *AsyncCreatePolicyResponse {
	s.RequestId = &v
	return s
}

func (s *AsyncCreatePolicyResponse) SetErrorMessage(v string) *AsyncCreatePolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AsyncCreatePolicyResponse) SetErrorCode(v string) *AsyncCreatePolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *AsyncCreatePolicyResponse) SetState(v AsyncCreatePolicyResponseStateEnum) *AsyncCreatePolicyResponse {
	s.State = &v
	return s
}

func (s *AsyncCreatePolicyResponse) SetBody(v string) *AsyncCreatePolicyResponse {
	s.Body = &v
	return s
}

func (s *AsyncCreatePolicyResponse) SetErrorParams(v []string) *AsyncCreatePolicyResponse {
	s.ErrorParams = v
	return s
}


type AsyncCreatePolicyResponseBuilder struct {
	s *AsyncCreatePolicyResponse
}

func NewAsyncCreatePolicyResponseBuilder() *AsyncCreatePolicyResponseBuilder {
	s := &AsyncCreatePolicyResponse{}
	b := &AsyncCreatePolicyResponseBuilder{s: s}
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) AsyncKey(v string) *AsyncCreatePolicyResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) RequestId(v string) *AsyncCreatePolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) ErrorMessage(v string) *AsyncCreatePolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) ErrorCode(v string) *AsyncCreatePolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) State(v AsyncCreatePolicyResponseStateEnum) *AsyncCreatePolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) Body(v string) *AsyncCreatePolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) ErrorParams(v []string) *AsyncCreatePolicyResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *AsyncCreatePolicyResponseBuilder) Build() *AsyncCreatePolicyResponse {
	return b.s
}


