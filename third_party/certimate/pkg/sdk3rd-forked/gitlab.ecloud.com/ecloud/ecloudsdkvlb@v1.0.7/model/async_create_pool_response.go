// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncCreatePoolResponseStateEnum string

// List of State
const (
    AsyncCreatePoolResponseStateEnumOk AsyncCreatePoolResponseStateEnum = "OK"
    AsyncCreatePoolResponseStateEnumError AsyncCreatePoolResponseStateEnum = "ERROR"
    AsyncCreatePoolResponseStateEnumException AsyncCreatePoolResponseStateEnum = "EXCEPTION"
    AsyncCreatePoolResponseStateEnumAlarm AsyncCreatePoolResponseStateEnum = "ALARM"
    AsyncCreatePoolResponseStateEnumForbidden AsyncCreatePoolResponseStateEnum = "FORBIDDEN"
)

type AsyncCreatePoolResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码： OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *AsyncCreatePoolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s AsyncCreatePoolResponse) String() string {
	return utils.Beautify(s)
}

func (s AsyncCreatePoolResponse) GoString() string {
	return s.String()
}

func (s AsyncCreatePoolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncCreatePoolResponse) SetAsyncKey(v string) *AsyncCreatePoolResponse {
	s.AsyncKey = &v
	return s
}

func (s *AsyncCreatePoolResponse) SetRequestId(v string) *AsyncCreatePoolResponse {
	s.RequestId = &v
	return s
}

func (s *AsyncCreatePoolResponse) SetErrorMessage(v string) *AsyncCreatePoolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AsyncCreatePoolResponse) SetErrorCode(v string) *AsyncCreatePoolResponse {
	s.ErrorCode = &v
	return s
}

func (s *AsyncCreatePoolResponse) SetState(v AsyncCreatePoolResponseStateEnum) *AsyncCreatePoolResponse {
	s.State = &v
	return s
}

func (s *AsyncCreatePoolResponse) SetBody(v string) *AsyncCreatePoolResponse {
	s.Body = &v
	return s
}

func (s *AsyncCreatePoolResponse) SetErrorParams(v []string) *AsyncCreatePoolResponse {
	s.ErrorParams = v
	return s
}


type AsyncCreatePoolResponseBuilder struct {
	s *AsyncCreatePoolResponse
}

func NewAsyncCreatePoolResponseBuilder() *AsyncCreatePoolResponseBuilder {
	s := &AsyncCreatePoolResponse{}
	b := &AsyncCreatePoolResponseBuilder{s: s}
	return b
}

func (b *AsyncCreatePoolResponseBuilder) AsyncKey(v string) *AsyncCreatePoolResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *AsyncCreatePoolResponseBuilder) RequestId(v string) *AsyncCreatePoolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AsyncCreatePoolResponseBuilder) ErrorMessage(v string) *AsyncCreatePoolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AsyncCreatePoolResponseBuilder) ErrorCode(v string) *AsyncCreatePoolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AsyncCreatePoolResponseBuilder) State(v AsyncCreatePoolResponseStateEnum) *AsyncCreatePoolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AsyncCreatePoolResponseBuilder) Body(v string) *AsyncCreatePoolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AsyncCreatePoolResponseBuilder) ErrorParams(v []string) *AsyncCreatePoolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *AsyncCreatePoolResponseBuilder) Build() *AsyncCreatePoolResponse {
	return b.s
}


