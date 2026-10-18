// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeletePoolResponseStateEnum string

// List of State
const (
    DeletePoolResponseStateEnumOk DeletePoolResponseStateEnum = "OK"
    DeletePoolResponseStateEnumError DeletePoolResponseStateEnum = "ERROR"
    DeletePoolResponseStateEnumException DeletePoolResponseStateEnum = "EXCEPTION"
    DeletePoolResponseStateEnumAlarm DeletePoolResponseStateEnum = "ALARM"
    DeletePoolResponseStateEnumForbidden DeletePoolResponseStateEnum = "FORBIDDEN"
)

type DeletePoolResponse struct {

    // 异步流程跟踪ID
	AsyncKey *string `json:"asyncKey,omitempty"`
    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码：OK（成功），ERROR（错误），EXCEPTION（异常），ALARM（警告），FORBIDDEN（禁止访问）
	State *DeletePoolResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
    // 统一错误码的自定义参数
	ErrorParams []string `json:"errorParams,omitempty"`
}

func (s DeletePoolResponse) String() string {
	return utils.Beautify(s)
}

func (s DeletePoolResponse) GoString() string {
	return s.String()
}

func (s DeletePoolResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePoolResponse) SetAsyncKey(v string) *DeletePoolResponse {
	s.AsyncKey = &v
	return s
}

func (s *DeletePoolResponse) SetRequestId(v string) *DeletePoolResponse {
	s.RequestId = &v
	return s
}

func (s *DeletePoolResponse) SetErrorMessage(v string) *DeletePoolResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePoolResponse) SetErrorCode(v string) *DeletePoolResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeletePoolResponse) SetState(v DeletePoolResponseStateEnum) *DeletePoolResponse {
	s.State = &v
	return s
}

func (s *DeletePoolResponse) SetBody(v string) *DeletePoolResponse {
	s.Body = &v
	return s
}

func (s *DeletePoolResponse) SetErrorParams(v []string) *DeletePoolResponse {
	s.ErrorParams = v
	return s
}


type DeletePoolResponseBuilder struct {
	s *DeletePoolResponse
}

func NewDeletePoolResponseBuilder() *DeletePoolResponseBuilder {
	s := &DeletePoolResponse{}
	b := &DeletePoolResponseBuilder{s: s}
	return b
}

func (b *DeletePoolResponseBuilder) AsyncKey(v string) *DeletePoolResponseBuilder {
	b.s.AsyncKey = &v
	return b
}

func (b *DeletePoolResponseBuilder) RequestId(v string) *DeletePoolResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeletePoolResponseBuilder) ErrorMessage(v string) *DeletePoolResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeletePoolResponseBuilder) ErrorCode(v string) *DeletePoolResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeletePoolResponseBuilder) State(v DeletePoolResponseStateEnum) *DeletePoolResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeletePoolResponseBuilder) Body(v string) *DeletePoolResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeletePoolResponseBuilder) ErrorParams(v []string) *DeletePoolResponseBuilder {
	b.s.ErrorParams = v
	return b
}

func (b *DeletePoolResponseBuilder) Build() *DeletePoolResponse {
	return b.s
}


