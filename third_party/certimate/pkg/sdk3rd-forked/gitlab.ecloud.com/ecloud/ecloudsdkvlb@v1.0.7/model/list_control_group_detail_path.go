// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListControlGroupDetailPath struct {
    position.Path
    // 访问控制组ID（未填写时将重定向到查询访问控制组列表接口）
	AccessControlGroupId *string `json:"accessControlGroupId,omitempty"`
}

func (s ListControlGroupDetailPath) String() string {
	return utils.Beautify(s)
}

func (s ListControlGroupDetailPath) GoString() string {
	return s.String()
}

func (s ListControlGroupDetailPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListControlGroupDetailPath) SetAccessControlGroupId(v string) *ListControlGroupDetailPath {
	s.AccessControlGroupId = &v
	return s
}


type ListControlGroupDetailPathBuilder struct {
	s *ListControlGroupDetailPath
}

func NewListControlGroupDetailPathBuilder() *ListControlGroupDetailPathBuilder {
	s := &ListControlGroupDetailPath{}
	b := &ListControlGroupDetailPathBuilder{s: s}
	return b
}

func (b *ListControlGroupDetailPathBuilder) AccessControlGroupId(v string) *ListControlGroupDetailPathBuilder {
	b.s.AccessControlGroupId = &v
	return b
}

func (b *ListControlGroupDetailPathBuilder) Build() *ListControlGroupDetailPath {
	return b.s
}


