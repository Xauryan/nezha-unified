// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteControlGroupRequest struct {


	BatchDeleteControlGroupBody *BatchDeleteControlGroupBody `json:"batchDeleteControlGroupBody,omitempty"`
}

func (s BatchDeleteControlGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteControlGroupRequest) GoString() string {
	return s.String()
}

func (s BatchDeleteControlGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteControlGroupRequest) SetBatchDeleteControlGroupBody(v *BatchDeleteControlGroupBody) *BatchDeleteControlGroupRequest {
	s.BatchDeleteControlGroupBody = v
	return s
}


type BatchDeleteControlGroupRequestBuilder struct {
	s *BatchDeleteControlGroupRequest
}

func NewBatchDeleteControlGroupRequestBuilder() *BatchDeleteControlGroupRequestBuilder {
	s := &BatchDeleteControlGroupRequest{}
	b := &BatchDeleteControlGroupRequestBuilder{s: s}
	return b
}

func (b *BatchDeleteControlGroupRequestBuilder) BatchDeleteControlGroupBody(v *BatchDeleteControlGroupBody) *BatchDeleteControlGroupRequestBuilder {
	b.s.BatchDeleteControlGroupBody = v
	return b
}

func (b *BatchDeleteControlGroupRequestBuilder) Build() *BatchDeleteControlGroupRequest {
	return b.s
}


