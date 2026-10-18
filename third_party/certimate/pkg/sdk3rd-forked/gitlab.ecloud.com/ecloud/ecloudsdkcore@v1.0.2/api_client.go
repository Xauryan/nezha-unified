package ecloudsdkcore

import (
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/client"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/consts"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/errs"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/param"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/response"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/retry"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/sign"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

// APIClient manages communication
type APIClient struct {
	Config      *config.Config
	HttpClient  *client.HttpClient
	HttpRequest *request.HttpRequest
}

// NewAPIClient creates a new API client.
func NewAPIClient() *APIClient {
	return &APIClient{HttpClient: client.NewHttpClient()}
}

// NewCustomizedAPIClient  creates a new customized API client.
func NewCustomizedAPIClient(config *config.Config, httpRequest *request.HttpRequest) *APIClient {
	return &APIClient{
		Config:      config,
		HttpClient:  client.NewHttpClient(),
		HttpRequest: httpRequest,
	}
}

func DefaultApiClient(config *config.Config, httpRequest *request.HttpRequest) *APIClient {
	return NewCustomizedAPIClient(config, httpRequest)
}

// InitConfig init default configuration
func InitConfig(c *config.Config) {
	if utils.IsUnSet(c.AutoRetry) {
		autoRetry := false
		c.AutoRetry = &autoRetry
	}
	if utils.IsUnSet(c.IgnoreSSL) {
		ignoreSSL := true
		c.IgnoreSSL = &ignoreSSL
	}
	if utils.IsUnSet(c.AuthType) {
		c.AuthType = utils.AuthType(consts.AK)
	}
	if utils.IsUnSet(c.IgnoreGateway) {
		ignoreGateway := false
		c.IgnoreGateway = &ignoreGateway
	}
	if utils.IsUnSet(c.CentralTransportEnabled) {
		centralTransportEnabled := true
		c.CentralTransportEnabled = &centralTransportEnabled
	}
}

// Excute entry for http call
func (c *APIClient) Excute(params *param.Params, rc *config.RuntimeConfig,
	returnType interface{}) (*response.HttpResponse, error) {
	httpReq := c.HttpRequest
	httpReq = new(request.HttpRequest)
	if err := utils.DeepCopy(httpReq, c.HttpRequest); err != nil {
		errs.NewServerRequestError("copy object error", nil)
	}
	if rc == nil {
		rc = &config.RuntimeConfig{}
	}
	c.resetHttpRequest(httpReq)
	if rc == nil {
		rc = &config.RuntimeConfig{}
	}
	if err := c.buildHttpRequest(params, rc, httpReq); err != nil {
		return nil, err
	}
	if err := c.buildHttpClient(rc); err != nil {
		return nil, err
	}
	autoRetry := utils.DefaultBool(rc.AutoRetry, c.Config.AutoRetry)
	if utils.BoolValue(autoRetry) {
		retryTemplate := c.buildRetryTemplate(rc)
		res, err := retryTemplate.Call(func() (interface{}, error) {
			return c.HttpClient.Execute(httpReq, returnType)
		})
		if err != nil {
			return nil, err
		}
		httpResponse, _ := res.(*response.HttpResponse)
		return httpResponse, nil
	}
	return c.HttpClient.Execute(httpReq, returnType)
}

func (c *APIClient) buildRetryTemplate(rc *config.RuntimeConfig) *retry.Template {
	templateBuilder := retry.NewBuilder()
	maxRetryTimes := utils.DefaultInt32(rc.MaxRetryTimes, c.Config.MaxRetryTimes)
	if maxRetryTimes != nil {
		templateBuilder.SetRetryTimes(utils.Int32Value(maxRetryTimes))
	}
	retryPeriod := utils.DefaultInt64(rc.RetryPeriod, c.Config.RetryPeriod)
	if retryPeriod != nil {
		templateBuilder.SetRetryPolicy(retry.WaitRetryPolicy(utils.Int64Value(retryPeriod)))
	} else {
		templateBuilder.SetRetryPolicy(retry.NoWait())
	}
	maxDuringTime := utils.DefaultInt64(rc.MaxDuringTime, c.Config.MaxDuringTime)
	if maxDuringTime != nil {
		templateBuilder.SetMaxDuringTime(utils.Int64Value(maxDuringTime))
	}
	return templateBuilder.Build()
}

func (c *APIClient) buildHttpRequest(params *param.Params, rc *config.RuntimeConfig, httpReq *request.HttpRequest) error {
	httpReq.ConvertRequest(params.Request)
	httpReq.BuildApiParams(params, c.Config, rc)
	authType := c.Config.AuthType
	if utils.IsSet(rc.AuthType) {
		authType = rc.AuthType
	}
	if *authType == consts.AK {
		err := sign.Sign(httpReq, utils.StringValue(c.Config.AccessKey), utils.StringValue(c.Config.SecretKey))
		if err != nil {
			return errs.NewSignatureError("signature error: "+err.Error(), err)
		}
	}
	httpReq.BuildFinalUrl()
	return nil
}

func (c *APIClient) buildHttpClient(rc *config.RuntimeConfig) error {
	readTimeout := utils.DefaultInt32(rc.ReadTimeout, c.Config.ReadTimeout)
	if readTimeout != nil {
		c.HttpClient.SetReadTimeout(utils.Int32Value(readTimeout))
	}
	connectTimeout := utils.DefaultInt32(rc.ConnectTimeout, c.Config.ConnectTimeout)
	if connectTimeout != nil {
		c.HttpClient.SetConnectTimeout(utils.Int32Value(connectTimeout))
	}
	ignoreSSL := utils.DefaultBool(rc.IgnoreSSL, c.Config.IgnoreSSL)
	ignore := true
	if ignoreSSL != nil {
		ignore = utils.BoolValue(ignoreSSL)
	}
	certFile := utils.StringValue(c.Config.CertFile)
	if rc.CertFile != nil {
		certFile = utils.StringValue(rc.CertFile)
	}
	clientCertFile := utils.StringValue(c.Config.ClientCertFile)
	if rc.ClientCertFile != nil {
		clientCertFile = utils.StringValue(rc.ClientCertFile)
	}
	clientKeyFile := utils.StringValue(c.Config.ClientKeyFile)
	if rc.ClientKeyFile != nil {
		clientKeyFile = utils.StringValue(rc.ClientKeyFile)
	}
	return c.HttpClient.ApplySSLSettings(ignore, certFile, clientCertFile, clientKeyFile)
}

func (c *APIClient) resetHttpRequest(req *request.HttpRequest) {
	req.Reset()
}
