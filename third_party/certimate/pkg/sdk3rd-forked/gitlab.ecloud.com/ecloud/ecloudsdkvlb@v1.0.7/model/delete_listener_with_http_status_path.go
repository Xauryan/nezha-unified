// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteListenerWithHttpStatusPath struct {
    position.Path
    // 负载均衡监听器ID
	ListenerId *string `json:"listenerId,omitempty"`
}

func (s DeleteListenerWithHttpStatusPath) String() string {
	return utils.Beautify(s)
}

func (s DeleteListenerWithHttpStatusPath) GoString() string {
	return s.String()
}

func (s DeleteListenerWithHttpStatusPath) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteListenerWithHttpStatusPath) SetListenerId(v string) *DeleteListenerWithHttpStatusPath {
	s.ListenerId = &v
	return s
}


type DeleteListenerWithHttpStatusPathBuilder struct {
	s *DeleteListenerWithHttpStatusPath
}

func NewDeleteListenerWithHttpStatusPathBuilder() *DeleteListenerWithHttpStatusPathBuilder {
	s := &DeleteListenerWithHttpStatusPath{}
	b := &DeleteListenerWithHttpStatusPathBuilder{s: s}
	return b
}

func (b *DeleteListenerWithHttpStatusPathBuilder) ListenerId(v string) *DeleteListenerWithHttpStatusPathBuilder {
	b.s.ListenerId = &v
	return b
}

func (b *DeleteListenerWithHttpStatusPathBuilder) Build() *DeleteListenerWithHttpStatusPath {
	return b.s
}


