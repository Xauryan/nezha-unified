// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DescribeUserCrtListResponse struct {

    // 请求Id
	RequestId *string `json:"requestId,omitempty"`
    // 页面国际化错误提示
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 状态是否正常
	State *string `json:"state,omitempty"`
    // 响应结果数据
	Body *DescribeUserCrtListResponseBody `json:"body,omitempty"`
}

func (s DescribeUserCrtListResponse) String() string {
	return utils.Beautify(s)
}

func (s DescribeUserCrtListResponse) GoString() string {
	return s.String()
}

func (s DescribeUserCrtListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DescribeUserCrtListResponse) SetRequestId(v string) *DescribeUserCrtListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUserCrtListResponse) SetErrorMessage(v string) *DescribeUserCrtListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeUserCrtListResponse) SetErrorCode(v string) *DescribeUserCrtListResponse {
	s.ErrorCode = &v
	return s
}

func (s *DescribeUserCrtListResponse) SetState(v string) *DescribeUserCrtListResponse {
	s.State = &v
	return s
}

func (s *DescribeUserCrtListResponse) SetBody(v *DescribeUserCrtListResponseBody) *DescribeUserCrtListResponse {
	s.Body = v
	return s
}


type DescribeUserCrtListResponseBuilder struct {
	s *DescribeUserCrtListResponse
}

func NewDescribeUserCrtListResponseBuilder() *DescribeUserCrtListResponseBuilder {
	s := &DescribeUserCrtListResponse{}
	b := &DescribeUserCrtListResponseBuilder{s: s}
	return b
}

func (b *DescribeUserCrtListResponseBuilder) RequestId(v string) *DescribeUserCrtListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DescribeUserCrtListResponseBuilder) ErrorMessage(v string) *DescribeUserCrtListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DescribeUserCrtListResponseBuilder) ErrorCode(v string) *DescribeUserCrtListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DescribeUserCrtListResponseBuilder) State(v string) *DescribeUserCrtListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DescribeUserCrtListResponseBuilder) Body(v *DescribeUserCrtListResponseBody) *DescribeUserCrtListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *DescribeUserCrtListResponseBuilder) Build() *DescribeUserCrtListResponse {
	return b.s
}


