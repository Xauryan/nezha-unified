// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteControlGroupRulePath struct {
    position.Path
    // 访问控制组ID
	AccessControlGroupId *string `json:"accessControlGroupId,omitempty"`
}

func (s DeleteControlGroupRulePath) String() string {
	return utils.Beautify(s)
}

func (s DeleteControlGroupRulePath) GoString() string {
	return s.String()
}

func (s DeleteControlGroupRulePath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteControlGroupRulePath) SetAccessControlGroupId(v string) *DeleteControlGroupRulePath {
	s.AccessControlGroupId = &v
	return s
}


type DeleteControlGroupRulePathBuilder struct {
	s *DeleteControlGroupRulePath
}

func NewDeleteControlGroupRulePathBuilder() *DeleteControlGroupRulePathBuilder {
	s := &DeleteControlGroupRulePath{}
	b := &DeleteControlGroupRulePathBuilder{s: s}
	return b
}

func (b *DeleteControlGroupRulePathBuilder) AccessControlGroupId(v string) *DeleteControlGroupRulePathBuilder {
	b.s.AccessControlGroupId = &v
	return b
}

func (b *DeleteControlGroupRulePathBuilder) Build() *DeleteControlGroupRulePath {
	return b.s
}


