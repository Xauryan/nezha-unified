// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package ecloudsdkcmcdn

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcmcdn/model"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/param"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type Client struct {
	apiClient   *ecloudsdkcore.APIClient
	config      *config.Config
	httpRequest *request.HttpRequest
	allRegions  map[string]string
}

func NewClient(config *config.Config) *Client {
	httpRequest := request.DefaultHttpRequest()
	httpRequest.Product = product
	httpRequest.Version = version
	httpRequest.SdkVersion = sdkVersion
	ecloudsdkcore.InitConfig(config)
	apiClient := ecloudsdkcore.DefaultApiClient(config, httpRequest)
	client := &Client{
		apiClient:   apiClient,
		config:      config,
		httpRequest: httpRequest,
	}
	client.allRegions = client.initRegions()
	client.setEndpoint(config, httpRequest)
	return client
}

const (
    product    string = "CM-CDN"
    version           = "v1"
    sdkVersion        = "1.0.0"
)

func (c *Client) initRegions() map[string]string {
	m := map[string]string{
        "CIDC-RP-04" : "https://console-yunnan-1.cmecloud.cn:8443",
        "CIDC-RP-16" : "https://console-qinghai-1.cmecloud.cn:8443",
        "CIDC-RP-25" : "https://console-wuxi-1.cmecloud.cn:8443",
        "CIDC-RP-26" : "https://console-dongguan-1.cmecloud.cn:8443",
        "CIDC-RP-27" : "https://console-yaan-1.cmecloud.cn:8443",
        "CIDC-RP-28" : "https://console-zhengzhou-1.cmecloud.cn:8443",
        "CIDC-RP-29" : "https://console-beijing-2.cmecloud.cn:8443",
        "CIDC-RP-30" : "https://console-zhuzhou-1.cmecloud.cn:8443",
        "CIDC-RP-31" : "https://console-jinan-1.cmecloud.cn:8443",
        "CIDC-RP-32" : "https://console-xian-1.cmecloud.cn:8443",
        "CIDC-RP-33" : "https://console-shanghai-1.cmecloud.cn:8443",
        "CIDC-RP-34" : "https://console-chongqing-1.cmecloud.cn:8443",
        "CIDC-RP-35" : "https://console-ningbo-1.cmecloud.cn:8443",
        "CIDC-RP-36" : "https://console-tianjin-1.cmecloud.cn:8443",
        "CIDC-RP-37" : "https://console-jilin-1.cmecloud.cn:8443",
        "CIDC-RP-38" : "https://console-hubei-1.cmecloud.cn:8443",
        "CIDC-RP-39" : "https://console-jiangxi-1.cmecloud.cn:8443",
        "CIDC-RP-40" : "https://console-gansu-1.cmecloud.cn:8443",
        "CIDC-RP-41" : "https://console-shanxi-1.cmecloud.cn:8443",
        "CIDC-RP-42" : "https://console-liaoning-1.cmecloud.cn:8443",
        "CIDC-RP-43" : "https://console-yunnan-2.cmecloud.cn:8443",
        "CIDC-RP-44" : "https://console-hebei-1.cmecloud.cn:8443",
        "CIDC-RP-45" : "https://console-fujian-1.cmecloud.cn:8443",
        "CIDC-RP-46" : "https://console-guangxi-1.cmecloud.cn:8443",
        "CIDC-RP-47" : "https://console-anhui-1.cmecloud.cn:8443",
        "CIDC-RP-48" : "https://console-huhehaote-1.cmecloud.cn:8443",
        "CIDC-RP-49" : "https://console-guiyang-1.cmecloud.cn:8443",
        "CIDC-CORE-00" : "https://ecloud.10086.cn",
        "CIDC-RP-53" : "https://console-hainan-1.cmecloud.cn:8443",
        "CIDC-RP-54" : "https://console-xinjiang-1.cmecloud.cn:8443",
        "CIDC-RP-55" : "http://console-heilongjiang-1.cmecloud.cn:18080",
        "CIDC-BRP-25" : "",
        "CIDC-RP-60" : "",
        "CIDC-RP-61" : "",
        "CIDC-RP-62" : "",
	}
	return m
}

func (c *Client) setEndpoint(config *config.Config, httpRequest *request.HttpRequest) {
	if utils.IsUnSet(config.PoolId) {
		httpRequest.Endpoint = utils.DefaultEndpoint
		return
	}
	endpoint := c.allRegions[*config.PoolId]
	if utils.IsUnSet(endpoint) {
		httpRequest.Endpoint = utils.DefaultEndpoint
		return
	}
	httpRequest.Endpoint = endpoint
}

// DeleteDomainServerCertificate 删除证书信息
func (c *Client) DeleteDomainServerCertificate(request *model.DeleteDomainServerCertificateRequest) (*model.DeleteDomainServerCertificateResponse, error) {
    return c.DeleteDomainServerCertificateWithConfig(request, nil)
}

// DeleteDomainServerCertificateWithConfig 删除证书信息
func (c *Client) DeleteDomainServerCertificateWithConfig(request *model.DeleteDomainServerCertificateRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteDomainServerCertificateResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteDomainServerCertificate").
        Uri("/domainManager/openapi/certificate/deleteDomainServerCertificate/{uniqueId}").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/certificate/deleteDomainServerCertificate/{uniqueId}").
        Protocol("https").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteDomainServerCertificateResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DescribeOperationLogs 查询刷新|预热
func (c *Client) DescribeOperationLogs(request *model.DescribeOperationLogsRequest) (*model.DescribeOperationLogsResponse, error) {
    return c.DescribeOperationLogsWithConfig(request, nil)
}

// DescribeOperationLogsWithConfig 查询刷新|预热
func (c *Client) DescribeOperationLogsWithConfig(request *model.DescribeOperationLogsRequest, runtimeConfig *config.RuntimeConfig) (*model.DescribeOperationLogsResponse, error) {
    params := param.NewParamsBuilder().
        Action("describeOperationLogs").
        Uri("/domainManager/openapi/content/describeOperationLogs").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/content/describeOperationLogs").
        Protocol("https").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.DescribeOperationLogsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DescribeCdnCertificateDetail 查询证书详情
func (c *Client) DescribeCdnCertificateDetail(request *model.DescribeCdnCertificateDetailRequest) (*model.DescribeCdnCertificateDetailResponse, error) {
    return c.DescribeCdnCertificateDetailWithConfig(request, nil)
}

// DescribeCdnCertificateDetailWithConfig 查询证书详情
func (c *Client) DescribeCdnCertificateDetailWithConfig(request *model.DescribeCdnCertificateDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.DescribeCdnCertificateDetailResponse, error) {
    params := param.NewParamsBuilder().
        Action("describeCdnCertificateDetail").
        Uri("/domainManager/openapi/certificate/describeCdnCertificateDetail/{uniqueId}").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/certificate/describeCdnCertificateDetail/{uniqueId}").
        Protocol("https").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.DescribeCdnCertificateDetailResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddCdnDomain 新增加速域名
func (c *Client) AddCdnDomain(request *model.AddCdnDomainRequest) (*model.AddCdnDomainResponse, error) {
    return c.AddCdnDomainWithConfig(request, nil)
}

// AddCdnDomainWithConfig 新增加速域名
func (c *Client) AddCdnDomainWithConfig(request *model.AddCdnDomainRequest, runtimeConfig *config.RuntimeConfig) (*model.AddCdnDomainResponse, error) {
    params := param.NewParamsBuilder().
        Action("addCdnDomain").
        Uri("/domainManager/openapi/domain/addCdnDomain").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/addCdnDomain").
        Protocol("https").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddCdnDomainResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DescribeUserCrtList 查询证书列表
func (c *Client) DescribeUserCrtList(request *model.DescribeUserCrtListRequest) (*model.DescribeUserCrtListResponse, error) {
    return c.DescribeUserCrtListWithConfig(request, nil)
}

// DescribeUserCrtListWithConfig 查询证书列表
func (c *Client) DescribeUserCrtListWithConfig(request *model.DescribeUserCrtListRequest, runtimeConfig *config.RuntimeConfig) (*model.DescribeUserCrtListResponse, error) {
    params := param.NewParamsBuilder().
        Action("describeUserCrtList").
        Uri("/domainManager/openapi/certificate/describeUserCrtList").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/certificate/describeUserCrtList").
        Protocol("https").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.DescribeUserCrtListResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DeleteCdnDomains 删除加速域名
func (c *Client) DeleteCdnDomains(request *model.DeleteCdnDomainsRequest) (*model.DeleteCdnDomainsResponse, error) {
    return c.DeleteCdnDomainsWithConfig(request, nil)
}

// DeleteCdnDomainsWithConfig 删除加速域名
func (c *Client) DeleteCdnDomainsWithConfig(request *model.DeleteCdnDomainsRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteCdnDomainsResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteCdnDomains").
        Uri("/domainManager/openapi/domain/deleteCdnDomain").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/deleteCdnDomain").
        Protocol("https").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteCdnDomainsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// UpdateDomainServerCertificate 修改证书信息
func (c *Client) UpdateDomainServerCertificate(request *model.UpdateDomainServerCertificateRequest) (*model.UpdateDomainServerCertificateResponse, error) {
    return c.UpdateDomainServerCertificateWithConfig(request, nil)
}

// UpdateDomainServerCertificateWithConfig 修改证书信息
func (c *Client) UpdateDomainServerCertificateWithConfig(request *model.UpdateDomainServerCertificateRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateDomainServerCertificateResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateDomainServerCertificate").
        Uri("/openapi/certificate/updateDomainServerCertificate").
        GatewayUri("/api/openapi-ecdn/openapi/certificate/updateDomainServerCertificate").
        Protocol("https").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateDomainServerCertificateResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DescribeCdnDomainDetail 查询域名详情
func (c *Client) DescribeCdnDomainDetail(request *model.DescribeCdnDomainDetailRequest) (*model.DescribeCdnDomainDetailResponse, error) {
    return c.DescribeCdnDomainDetailWithConfig(request, nil)
}

// DescribeCdnDomainDetailWithConfig 查询域名详情
func (c *Client) DescribeCdnDomainDetailWithConfig(request *model.DescribeCdnDomainDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.DescribeCdnDomainDetailResponse, error) {
    params := param.NewParamsBuilder().
        Action("describeCdnDomainDetail").
        Uri("/domainManager/openapi/domain/describeCdnDomainDetail/{domainId}").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/describeCdnDomainDetail/{domainId}").
        Protocol("https").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.DescribeCdnDomainDetailResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// DescribeUserDomains 查询域名列表
func (c *Client) DescribeUserDomains(request *model.DescribeUserDomainsRequest) (*model.DescribeUserDomainsResponse, error) {
    return c.DescribeUserDomainsWithConfig(request, nil)
}

// DescribeUserDomainsWithConfig 查询域名列表
func (c *Client) DescribeUserDomainsWithConfig(request *model.DescribeUserDomainsRequest, runtimeConfig *config.RuntimeConfig) (*model.DescribeUserDomainsResponse, error) {
    params := param.NewParamsBuilder().
        Action("describeUserDomains").
        Uri("/domainManager/openapi/domain/describeUserDomains").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/describeUserDomains").
        Protocol("https").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.DescribeUserDomainsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// ModifyCdnDomain 修改域名信息
func (c *Client) ModifyCdnDomain(request *model.ModifyCdnDomainRequest) (*model.ModifyCdnDomainResponse, error) {
    return c.ModifyCdnDomainWithConfig(request, nil)
}

// ModifyCdnDomainWithConfig 修改域名信息
func (c *Client) ModifyCdnDomainWithConfig(request *model.ModifyCdnDomainRequest, runtimeConfig *config.RuntimeConfig) (*model.ModifyCdnDomainResponse, error) {
    params := param.NewParamsBuilder().
        Action("modifyCdnDomain").
        Uri("/domainManager/openapi/domain/modifyCdnDomain").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/modifyCdnDomain").
        Protocol("https").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.ModifyCdnDomainResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// SetCdnDomainConfig 修改域名配置
func (c *Client) SetCdnDomainConfig(request *model.SetCdnDomainConfigRequest) (*model.SetCdnDomainConfigResponse, error) {
    return c.SetCdnDomainConfigWithConfig(request, nil)
}

// SetCdnDomainConfigWithConfig 修改域名配置
func (c *Client) SetCdnDomainConfigWithConfig(request *model.SetCdnDomainConfigRequest, runtimeConfig *config.RuntimeConfig) (*model.SetCdnDomainConfigResponse, error) {
    params := param.NewParamsBuilder().
        Action("setCdnDomainConfig").
        Uri("/domainManager/openapi/domain/setCdnDomainConfig").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/setCdnDomainConfig").
        Protocol("https").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.SetCdnDomainConfigResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// MonitorBandwidthFlow 查询域名用量
func (c *Client) MonitorBandwidthFlow(request *model.MonitorBandwidthFlowRequest) (*model.MonitorBandwidthFlowResponse, error) {
    return c.MonitorBandwidthFlowWithConfig(request, nil)
}

// MonitorBandwidthFlowWithConfig 查询域名用量
func (c *Client) MonitorBandwidthFlowWithConfig(request *model.MonitorBandwidthFlowRequest, runtimeConfig *config.RuntimeConfig) (*model.MonitorBandwidthFlowResponse, error) {
    params := param.NewParamsBuilder().
        Action("monitorBandwidthFlow").
        Uri("/statistics/openapi/analysis/describeDomainBandwidthFlow").
        GatewayUri("/api/openapi-ecdn/statistics/openapi/analysis/describeDomainBandwidthFlow").
        Protocol("https").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.MonitorBandwidthFlowResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// StartCdnDomain 启用域名加速
func (c *Client) StartCdnDomain(request *model.StartCdnDomainRequest) (*model.StartCdnDomainResponse, error) {
    return c.StartCdnDomainWithConfig(request, nil)
}

// StartCdnDomainWithConfig 启用域名加速
func (c *Client) StartCdnDomainWithConfig(request *model.StartCdnDomainRequest, runtimeConfig *config.RuntimeConfig) (*model.StartCdnDomainResponse, error) {
    params := param.NewParamsBuilder().
        Action("startCdnDomain").
        Uri("/domainManager/openapi/domain/startCdnDomain").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/startCdnDomain").
        Protocol("https").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.StartCdnDomainResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// CheckCdnFromApi 查询开通状态
func (c *Client) CheckCdnFromApi(request *model.CheckCdnFromApiRequest) (*model.CheckCdnFromApiResponse, error) {
    return c.CheckCdnFromApiWithConfig(request, nil)
}

// CheckCdnFromApiWithConfig 查询开通状态
func (c *Client) CheckCdnFromApiWithConfig(request *model.CheckCdnFromApiRequest, runtimeConfig *config.RuntimeConfig) (*model.CheckCdnFromApiResponse, error) {
    params := param.NewParamsBuilder().
        Action("checkCdnFromApi").
        Uri("/orderManager/openapi/order/checkCdnFromApi/{productType}").
        GatewayUri("/api/openapi-ecdn/orderManager/openapi/order/checkCdnFromApi/{productType}").
        Protocol("https").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.CheckCdnFromApiResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// StopCdnDomain 暂停域名加速
func (c *Client) StopCdnDomain(request *model.StopCdnDomainRequest) (*model.StopCdnDomainResponse, error) {
    return c.StopCdnDomainWithConfig(request, nil)
}

// StopCdnDomainWithConfig 暂停域名加速
func (c *Client) StopCdnDomainWithConfig(request *model.StopCdnDomainRequest, runtimeConfig *config.RuntimeConfig) (*model.StopCdnDomainResponse, error) {
    params := param.NewParamsBuilder().
        Action("stopCdnDomain").
        Uri("/domainManager/openapi/domain/stopCdnDomain").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/domain/stopCdnDomain").
        Protocol("https").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.StopCdnDomainResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// OpenCdnServiceFromApi 用户开通服务
func (c *Client) OpenCdnServiceFromApi(request *model.OpenCdnServiceFromApiRequest) (*model.OpenCdnServiceFromApiResponse, error) {
    return c.OpenCdnServiceFromApiWithConfig(request, nil)
}

// OpenCdnServiceFromApiWithConfig 用户开通服务
func (c *Client) OpenCdnServiceFromApiWithConfig(request *model.OpenCdnServiceFromApiRequest, runtimeConfig *config.RuntimeConfig) (*model.OpenCdnServiceFromApiResponse, error) {
    params := param.NewParamsBuilder().
        Action("openCdnServiceFromApi").
        Uri("/orderManager/openapi/order/openCdnServiceFromApi").
        GatewayUri("/api/openapi-ecdn/orderManager/openapi/order/openCdnServiceFromApi").
        Protocol("https").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.OpenCdnServiceFromApiResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddDomainServerCertificate 新增证书
func (c *Client) AddDomainServerCertificate(request *model.AddDomainServerCertificateRequest) (*model.AddDomainServerCertificateResponse, error) {
    return c.AddDomainServerCertificateWithConfig(request, nil)
}

// AddDomainServerCertificateWithConfig 新增证书
func (c *Client) AddDomainServerCertificateWithConfig(request *model.AddDomainServerCertificateRequest, runtimeConfig *config.RuntimeConfig) (*model.AddDomainServerCertificateResponse, error) {
    params := param.NewParamsBuilder().
        Action("addDomainServerCertificate").
        Uri("/domainManager/openapi/certificate/addDomainServerCertificate").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/certificate/addDomainServerCertificate").
        Protocol("https").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddDomainServerCertificateResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

// AddRefresh 新增刷新|预热
func (c *Client) AddRefresh(request *model.AddRefreshRequest) (*model.AddRefreshResponse, error) {
    return c.AddRefreshWithConfig(request, nil)
}

// AddRefreshWithConfig 新增刷新|预热
func (c *Client) AddRefreshWithConfig(request *model.AddRefreshRequest, runtimeConfig *config.RuntimeConfig) (*model.AddRefreshResponse, error) {
    params := param.NewParamsBuilder().
        Action("addRefresh").
        Uri("/domainManager/openapi/content/addRefresh").
        GatewayUri("/api/openapi-ecdn/domainManager/openapi/content/addRefresh").
        Protocol("https").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AddRefreshResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
    }

