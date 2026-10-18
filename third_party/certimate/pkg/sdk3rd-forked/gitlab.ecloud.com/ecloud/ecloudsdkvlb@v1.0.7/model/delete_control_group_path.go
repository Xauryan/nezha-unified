// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteControlGroupPath struct {
    position.Path
    // 访问控制组ID
	AccessControlGroupId *string `json:"accessControlGroupId,omitempty"`
}

func (s DeleteControlGroupPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteControlGroupPath) GoString() string {
	return s.String()
}

func (s DeleteControlGroupPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteControlGroupPath) SetAccessControlGroupId(v string) *DeleteControlGroupPath {
	s.AccessControlGroupId = &v
	return s
}


type DeleteControlGroupPathBuilder struct {
	s *DeleteControlGroupPath
}

func NewDeleteControlGroupPathBuilder() *DeleteControlGroupPathBuilder {
	s := &DeleteControlGroupPath{}
	b := &DeleteControlGroupPathBuilder{s: s}
	return b
}

func (b *DeleteControlGroupPathBuilder) AccessControlGroupId(v string) *DeleteControlGroupPathBuilder {
	b.s.AccessControlGroupId = &v
	return b
}

func (b *DeleteControlGroupPathBuilder) Build() *DeleteControlGroupPath {
	return b.s
}


