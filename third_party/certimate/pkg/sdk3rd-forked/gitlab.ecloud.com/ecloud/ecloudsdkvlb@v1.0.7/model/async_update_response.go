// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AsyncUpdateResponseStateEnum string

// List of State
const (
    AsyncUpdateResponseStateEnumOk AsyncUpdateResponseStateEnum = "OK"
    AsyncUpdateResponseStateEnumError AsyncUpdateResponseStateEnum = "ERROR"
    AsyncUpdateResponseStateEnumException AsyncUpdateResponseStateEnum = "EXCEPTION"
    AsyncUpdateResponseStateEnumAlarm AsyncUpdateResponseStateEnum = "ALARM"
    AsyncUpdateResponseStateEnumForbidden AsyncUpdateResponseStateEnum = "FORBIDDEN"
)

type AsyncUpdateResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *AsyncUpdateResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s AsyncUpdateResponse) String() string {
	return utils.Beautify(s)
}

func (s AsyncUpdateResponse) GoString() string {
	return s.String()
}

func (s AsyncUpdateResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AsyncUpdateResponse) SetAsyncKey(v string) *AsyncUpdateResponse {
	s.AsyncKey = &v
	return s
}

func (s *AsyncUpdateResponse) SetRequestId(v string) *AsyncUpdateResponse {
	s.RequestId = &v
	return s
}

func (s *AsyncUpdateResponse) SetErrorMessage(v string) *AsyncUpdateResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AsyncUpdateResponse) SetErrorCode(v string) *AsyncUpdateResponse {
	s.ErrorCode = &v
	return s
}

func (s *AsyncUpdateResponse) SetState(v AsyncUpdateResponseStateEnum) *AsyncUpdateResponse {
	s.State = &v
	return s
}

func (s *AsyncUpdateResponse) SetBody(v string) *AsyncUpdateResponse {
	s.Body = &v
	return s
}

func (s *AsyncUpdateResponse) SetErrorParams(v []string) *AsyncUpdateResponse {
	s.ErrorParams = v
	return s
}


type AsyncUpdateResponseBuilder struct {
	s *AsyncUpdateResponse
}

func NewAsyncUpdateResponseBuilder() *AsyncUpdateResponseBuilder {
	s := &AsyncUpdateResponse{}
	b := &AsyncUpdateResponseBuilder{s: s}
	return b
}

func (b *AsyncUpdateResponseBuilder) AsyncKey(v string) *AsyncUpdateResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *AsyncUpdateResponseBuilder) RequestId(v string) *AsyncUpdateResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AsyncUpdateResponseBuilder) ErrorMessage(v string) *AsyncUpdateResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AsyncUpdateResponseBuilder) ErrorCode(v string) *AsyncUpdateResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AsyncUpdateResponseBuilder) State(v AsyncUpdateResponseStateEnum) *AsyncUpdateResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AsyncUpdateResponseBuilder) Body(v string) *AsyncUpdateResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AsyncUpdateResponseBuilder) ErrorParams(v []string) *AsyncUpdateResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *AsyncUpdateResponseBuilder) Build() *AsyncUpdateResponse {
	return b.s
}


