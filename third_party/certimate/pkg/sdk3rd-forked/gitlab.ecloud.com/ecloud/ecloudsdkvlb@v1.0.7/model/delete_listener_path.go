// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteListenerPath struct {
    position.Path
    // 负载均衡监听器的ID
	ListenerId *string `json:"listenerId,omitempty"`
}

func (s DeleteListenerPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerPath) GoString() string {
	return s.String()
}

func (s DeleteListenerPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerPath) SetListenerId(v string) *DeleteListenerPath {
	s.ListenerId = &v
	return s
}


type DeleteListenerPathBuilder struct {
	s *DeleteListenerPath
}

func NewDeleteListenerPathBuilder() *DeleteListenerPathBuilder {
	s := &DeleteListenerPath{}
	b := &DeleteListenerPathBuilder{s: s}
	return b
}

func (b *DeleteListenerPathBuilder) ListenerId(v string) *DeleteListenerPathBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *DeleteListenerPathBuilder) Build() *DeleteListenerPath {
	return b.s
}


