// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteControlGroupBody struct {
    position.Body
    // 访问控制组ID
	AccessControlGroupIds []string 
}

func (s BatchDeleteControlGroupBody) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteControlGroupBody) GoString() string {
	return s.String()
}

func (s BatchDeleteControlGroupBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteControlGroupBody) SetAccessControlGroupIds(v []string) *BatchDeleteControlGroupBody {
	s.AccessControlGroupIds = v
	return s
}


type BatchDeleteControlGroupBodyBuilder struct {
	s *BatchDeleteControlGroupBody
}

func NewBatchDeleteControlGroupBodyBuilder() *BatchDeleteControlGroupBodyBuilder {
	s := &BatchDeleteControlGroupBody{}
	b := &BatchDeleteControlGroupBodyBuilder{s: s}
	return b
}

func (b *BatchDeleteControlGroupBodyBuilder) AccessControlGroupIds(v []string) *BatchDeleteControlGroupBodyBuilder {
	b.s.AccessControlGroupIds = v
	return b
}

func (b *BatchDeleteControlGroupBodyBuilder) Build() *BatchDeleteControlGroupBody {
	return b.s
}


