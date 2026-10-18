// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAsyncResultQuery struct {
    position.Query
    // 获取异步执行结果的Key
	RedisKey *string `json:"redisKey,omitempty"`
}

func (s GetAsyncResultQuery) String() string {
	return utils.Beautify(s)
}

func (s GetAsyncResultQuery) GoString() string {
	return s.String()
}

func (s GetAsyncResultQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAsyncResultQuery) SetRedisKey(v string) *GetAsyncResultQuery {
	s.RedisKey = &v
	return s
}


type GetAsyncResultQueryBuilder struct {
	s *GetAsyncResultQuery
}

func NewGetAsyncResultQueryBuilder() *GetAsyncResultQueryBuilder {
	s := &GetAsyncResultQuery{}
	b := &GetAsyncResultQueryBuilder{s: s}
	return b
}

func (b *GetAsyncResultQueryBuilder) RedisKey(v string) *GetAsyncResultQueryBuilder {
	b.s.RedisKey = &v
	return b
}

func (b *GetAsyncResultQueryBuilder) Build() *GetAsyncResultQuery {
	return b.s
}


