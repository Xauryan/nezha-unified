// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ValidateControlGroupPath struct {
    position.Path
    // 访问控制组ID
	AccessControlGroupId *string `json:"accessControlGroupId,omitempty"`
}

func (s ValidateControlGroupPath) String() string {
	return utils.Beautify(s)
}

func (s ValidateControlGroupPath) GoString() string {
	return s.String()
}

func (s ValidateControlGroupPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ValidateControlGroupPath) SetAccessControlGroupId(v string) *ValidateControlGroupPath {
	s.AccessControlGroupId = &v
	return s
}


type ValidateControlGroupPathBuilder struct {
	s *ValidateControlGroupPath
}

func NewValidateControlGroupPathBuilder() *ValidateControlGroupPathBuilder {
	s := &ValidateControlGroupPath{}
	b := &ValidateControlGroupPathBuilder{s: s}
	return b
}

func (b *ValidateControlGroupPathBuilder) AccessControlGroupId(v string) *ValidateControlGroupPathBuilder {
	b.s.AccessControlGroupId = &v
	return b
}

func (b *ValidateControlGroupPathBuilder) Build() *ValidateControlGroupPath {
	return b.s
}


