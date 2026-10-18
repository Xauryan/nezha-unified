// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteListenerAsyncPath struct {
    position.Path
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
}

func (s DeleteListenerAsyncPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerAsyncPath) GoString() string {
	return s.String()
}

func (s DeleteListenerAsyncPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerAsyncPath) SetListenerId(v string) *DeleteListenerAsyncPath {
	s.ListenerId = &v
	return s
}


type DeleteListenerAsyncPathBuilder struct {
	s *DeleteListenerAsyncPath
}

func NewDeleteListenerAsyncPathBuilder() *DeleteListenerAsyncPathBuilder {
	s := &DeleteListenerAsyncPath{}
	b := &DeleteListenerAsyncPathBuilder{s: s}
	return b
}

func (b *DeleteListenerAsyncPathBuilder) ListenerId(v string) *DeleteListenerAsyncPathBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *DeleteListenerAsyncPathBuilder) Build() *DeleteListenerAsyncPath {
	return b.s
}


