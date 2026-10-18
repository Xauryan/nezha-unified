// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeOperationLogsResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *DescribeOperationLogsResponseBody `json:"body,omitempty"`
}

func (s DescribeOperationLogsResponse) String() string {
	return utils.Beautify(s)
}

func (s DescribeOperationLogsResponse) GoString() string {
	return s.String()
}

func (s DescribeOperationLogsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeOperationLogsResponse) SetRequestId(v string) *DescribeOperationLogsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOperationLogsResponse) SetErrorMessage(v string) *DescribeOperationLogsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeOperationLogsResponse) SetErrorCode(v string) *DescribeOperationLogsResponse {
	s.ErrorCode = &v
	return s
}

func (s *DescribeOperationLogsResponse) SetState(v string) *DescribeOperationLogsResponse {
	s.State = &v
	return s
}

func (s *DescribeOperationLogsResponse) SetBody(v *DescribeOperationLogsResponseBody) *DescribeOperationLogsResponse {
	s.Body = v
	return s
}


type DescribeOperationLogsResponseBuilder struct {
	s *DescribeOperationLogsResponse
}

func NewDescribeOperationLogsResponseBuilder() *DescribeOperationLogsResponseBuilder {
	s := &DescribeOperationLogsResponse{}
	b := &DescribeOperationLogsResponseBuilder{s: s}
	return b
}

func (b *DescribeOperationLogsResponseBuilder) RequestId(v string) *DescribeOperationLogsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DescribeOperationLogsResponseBuilder) ErrorMessage(v string) *DescribeOperationLogsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DescribeOperationLogsResponseBuilder) ErrorCode(v string) *DescribeOperationLogsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DescribeOperationLogsResponseBuilder) State(v string) *DescribeOperationLogsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DescribeOperationLogsResponseBuilder) Body(v *DescribeOperationLogsResponseBody) *DescribeOperationLogsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DescribeOperationLogsResponseBuilder) Build() *DescribeOperationLogsResponse {
	return b.s
}


