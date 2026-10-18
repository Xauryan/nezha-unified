// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeletePoolAsyncRequest struct {


	DeletePoolAsyncPath *DeletePoolAsyncPath `json:"deletePoolAsyncPath,omitempty"`
}

func (s DeletePoolAsyncRequest) String() string {
	return utils.Beautify(s)
}

func (s DeletePoolAsyncRequest) GoString() string {
	return s.String()
}

func (s DeletePoolAsyncRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeletePoolAsyncRequest) SetDeletePoolAsyncPath(v *DeletePoolAsyncPath) *DeletePoolAsyncRequest {
	s.DeletePoolAsyncPath = v
	return s
}


type DeletePoolAsyncRequestBuilder struct {
	s *DeletePoolAsyncRequest
}

func NewDeletePoolAsyncRequestBuilder() *DeletePoolAsyncRequestBuilder {
	s := &DeletePoolAsyncRequest{}
	b := &DeletePoolAsyncRequestBuilder{s: s}
	return b
}

func (b *DeletePoolAsyncRequestBuilder) DeletePoolAsyncPath(v *DeletePoolAsyncPath) *DeletePoolAsyncRequestBuilder {
	b.s.DeletePoolAsyncPath = v
	return b
}

func (b *DeletePoolAsyncRequestBuilder) Build() *DeletePoolAsyncRequest {
	return b.s
}


