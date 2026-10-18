package config

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/consts"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type Config struct {
	AccessKey               *string           `json:"accessKey"`
	SecretKey               *string           `json:"secretKey"`
	PoolId                  *string           `json:"poolId"`
	ReadTimeout             *int32            `json:"readTimeout"`
	ConnectTimeout          *int32            `json:"connectTimeout"`
	Protocol                *string           `json:"protocol"`
	AutoRetry               *bool             `json:"autoRetry"`
	IgnoreSSL               *bool             `json:"ignoreSSL"`
	CertFile                *string           `json:"certFile"`
	ClientCertFile          *string           `json:"clientCertFile"`
	ClientKeyFile           *string           `json:"clientKeyFile"`
	MaxRetryTimes           *int32            `json:"maxRetryTimes"`
	RetryPeriod             *int64            `json:"retryPeriod"`
	MaxDuringTime           *int64            `json:"maxDuringTime"`
	HttpProxy               *string           `json:"httpProxy"`
	HttpsProxy              *string           `json:"httpsProxy"`
	AuthType                *consts.AuthType  `json:"authType"`
	IgnoreGateway           *bool             `json:"ignoreGateway"`
	CentralTransportEnabled *bool             `json:"centralTransportEnabled"`
	GlobalQueryParams       map[string]string `json:"globalQueryParams"`
	GlobalHeaderParams      map[string]string `json:"globalHeaderParams"`
}

func NewConfig() *Config {
	return &Config{
		Protocol:                utils.String("http"),
		IgnoreSSL:               utils.Bool(true),
		AuthType:                utils.AuthType(consts.AK),
		IgnoreGateway:           utils.Bool(false),
		CentralTransportEnabled: utils.Bool(true),
		GlobalHeaderParams:      map[string]string{},
		GlobalQueryParams:       map[string]string{},
	}
}

func (c *Config) String() string {
	return utils.Beautify(c)
}

func (c *Config) GoString() string {
	return c.String()
}

func (c *Config) ToJsonString() string {
	return utils.ToJsonString(c)
}

type ConfigBuilder struct {
	config *Config
}

func NewConfigBuilder() *ConfigBuilder {
	config := NewConfig()
	b := &ConfigBuilder{config: config}
	return b
}

func (c *ConfigBuilder) AccessKey(accessKey string) *ConfigBuilder {
	c.config.AccessKey = utils.String(accessKey)
	return c
}

func (c *ConfigBuilder) SecretKey(secretKey string) *ConfigBuilder {
	c.config.SecretKey = utils.String(secretKey)
	return c
}

func (c *ConfigBuilder) PoolId(poolId string) *ConfigBuilder {
	c.config.PoolId = utils.String(poolId)
	return c
}

func (c *ConfigBuilder) ReadTimeOut(readTimeOut int32) *ConfigBuilder {
	c.config.ReadTimeout = utils.Int32(readTimeOut)
	return c
}

func (c *ConfigBuilder) ConnectTimeout(connectTimeout int32) *ConfigBuilder {
	c.config.ConnectTimeout = utils.Int32(connectTimeout)
	return c
}

func (c *ConfigBuilder) Protocol(protocol string) *ConfigBuilder {
	c.config.Protocol = utils.String(protocol)
	return c
}

func (c *ConfigBuilder) AutoRetry(autoRetry bool) *ConfigBuilder {
	c.config.AutoRetry = utils.Bool(autoRetry)
	return c
}

func (c *ConfigBuilder) IgnoreSSL(ignoreSSL bool) *ConfigBuilder {
	c.config.IgnoreSSL = utils.Bool(ignoreSSL)
	return c
}

func (c *ConfigBuilder) CertFile(certFile string) *ConfigBuilder {
	c.config.CertFile = utils.String(certFile)
	return c
}

func (c *ConfigBuilder) ClientCertFile(clientCertFile string) *ConfigBuilder {
	c.config.ClientCertFile = utils.String(clientCertFile)
	return c
}

func (c *ConfigBuilder) ClientKeyFile(clientKeyFile string) *ConfigBuilder {
	c.config.ClientKeyFile = utils.String(clientKeyFile)
	return c
}

func (c *ConfigBuilder) MaxRetryTimes(maxRetryTimes int32) *ConfigBuilder {
	c.config.MaxRetryTimes = utils.Int32(maxRetryTimes)
	return c
}

func (c *ConfigBuilder) RetryPeriod(retryPeriod int64) *ConfigBuilder {
	c.config.RetryPeriod = utils.Int64(retryPeriod)
	return c
}

func (c *ConfigBuilder) MaxDuringTime(maxDuringTime int64) *ConfigBuilder {
	c.config.MaxDuringTime = utils.Int64(maxDuringTime)
	return c
}

func (c *ConfigBuilder) HttpProxy(httpProxy string) *ConfigBuilder {
	c.config.HttpProxy = utils.String(httpProxy)
	return c
}

func (c *ConfigBuilder) HttpsProxy(httpsProxy string) *ConfigBuilder {
	c.config.HttpsProxy = utils.String(httpsProxy)
	return c
}

func (c *ConfigBuilder) AuthType(authType consts.AuthType) *ConfigBuilder {
	c.config.AuthType = utils.AuthType(authType)
	return c
}

func (c *ConfigBuilder) IgnoreGateway(ignoreGateway bool) *ConfigBuilder {
	c.config.IgnoreGateway = utils.Bool(ignoreGateway)
	return c
}

func (c *ConfigBuilder) CentralTransportEnabled(centralTransportEnabled bool) *ConfigBuilder {
	c.config.CentralTransportEnabled = utils.Bool(centralTransportEnabled)
	return c
}

func (c *ConfigBuilder) GlobalQueryParams(globalQueryParams map[string]string) *ConfigBuilder {
	c.config.GlobalQueryParams = globalQueryParams
	return c
}

func (c *ConfigBuilder) GlobalHeaderParams(globalHeaderParams map[string]string) *ConfigBuilder {
	c.config.GlobalHeaderParams = globalHeaderParams
	return c
}

func (c *ConfigBuilder) Build() *Config {
	return c.config
}
