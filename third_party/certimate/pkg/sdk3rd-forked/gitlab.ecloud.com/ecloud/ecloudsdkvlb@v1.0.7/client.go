// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package ecloudsdkvlb

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkvlb/model"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/param"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/region"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type Client struct {
	apiClient   *ecloudsdkcore.APIClient
	config      *config.Config
	httpRequest *request.HttpRequest
	allRegions  []region.Region
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
    product    string = "vlb"
    version           = "v1"
    sdkVersion        = "1.0.7"
)

func (c *Client) initRegions() []region.Region {
    var regions []region.Region
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-jiangsu-1"),
              		PoolId:   utils.String("CIDC-RP-25"),
              		Endpoint: utils.String("https://api-wuxi-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-guangdong-1"),
              		PoolId:   utils.String("CIDC-RP-26"),
              		Endpoint: utils.String("https://api-dongguan-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-sichuan-1"),
              		PoolId:   utils.String("CIDC-RP-27"),
              		Endpoint: utils.String("https://api-yaan-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-henan-1"),
              		PoolId:   utils.String("CIDC-RP-28"),
              		Endpoint: utils.String("https://api-zhengzhou-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-beijing-1"),
              		PoolId:   utils.String("CIDC-RP-29"),
              		Endpoint: utils.String("https://api-beijing-2.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-hunan-1"),
              		PoolId:   utils.String("CIDC-RP-30"),
              		Endpoint: utils.String("https://api-zhuzhou-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-shandong-1"),
              		PoolId:   utils.String("CIDC-RP-31"),
              		Endpoint: utils.String("https://api-jinan-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-shaanxi-1"),
              		PoolId:   utils.String("CIDC-RP-32"),
              		Endpoint: utils.String("https://api-xian-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-shanghai-1"),
              		PoolId:   utils.String("CIDC-RP-33"),
              		Endpoint: utils.String("https://api-shanghai-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-chongqing-1"),
              		PoolId:   utils.String("CIDC-RP-34"),
              		Endpoint: utils.String("https://api-chongqing-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-zhejiang-1"),
              		PoolId:   utils.String("CIDC-RP-35"),
              		Endpoint: utils.String("https://api-ningbo-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-tianjin-1"),
              		PoolId:   utils.String("CIDC-RP-36"),
              		Endpoint: utils.String("https://api-tianjin-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-jilin-1"),
              		PoolId:   utils.String("CIDC-RP-37"),
              		Endpoint: utils.String("https://api-jilin-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-hubei-1"),
              		PoolId:   utils.String("CIDC-RP-38"),
              		Endpoint: utils.String("https://api-hubei-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-jiangxi-1"),
              		PoolId:   utils.String("CIDC-RP-39"),
              		Endpoint: utils.String("https://api-jiangxi-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-gansu-1"),
              		PoolId:   utils.String("CIDC-RP-40"),
              		Endpoint: utils.String("https://api-gansu-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-shangxi-1"),
              		PoolId:   utils.String("CIDC-RP-41"),
              		Endpoint: utils.String("https://api-shanxi-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-liaoning-1"),
              		PoolId:   utils.String("CIDC-RP-42"),
              		Endpoint: utils.String("https://api-liaoning-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-yunnan-1"),
              		PoolId:   utils.String("CIDC-RP-43"),
              		Endpoint: utils.String("https://api-yunnan-2.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-hebei-1"),
              		PoolId:   utils.String("CIDC-RP-44"),
              		Endpoint: utils.String("https://api-hebei-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-fujian-1"),
              		PoolId:   utils.String("CIDC-RP-45"),
              		Endpoint: utils.String("https://api-fujian-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-guangxi-1"),
              		PoolId:   utils.String("CIDC-RP-46"),
              		Endpoint: utils.String("https://api-guangxi-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-anhui-1"),
              		PoolId:   utils.String("CIDC-RP-47"),
              		Endpoint: utils.String("https://api-anhui-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-neimenggu-1"),
              		PoolId:   utils.String("CIDC-RP-48"),
              		Endpoint: utils.String("https://api-huhehaote-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-guzhou-1"),
              		PoolId:   utils.String("CIDC-RP-49"),
              		Endpoint: utils.String("https://api-guiyang-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-hainan-1"),
              		PoolId:   utils.String("CIDC-RP-53"),
              		Endpoint: utils.String("https://api-hainan-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-xinjiang-1"),
              		PoolId:   utils.String("CIDC-RP-54"),
              		Endpoint: utils.String("https://api-xinjiang-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-heilongjiang-1"),
              		PoolId:   utils.String("CIDC-RP-55"),
              		Endpoint: utils.String("https://api-heilongjiang-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-ningxia-1"),
              		PoolId:   utils.String("CIDC-RP-60"),
              		Endpoint: utils.String("https://api-ningxia-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-qinghai-1"),
              		PoolId:   utils.String("CIDC-RP-61"),
              		Endpoint: utils.String("https://api-qinghai-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-xizang-1"),
              		PoolId:   utils.String("CIDC-RP-62"),
              		Endpoint: utils.String("https://api-xizang-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-hubei-2"),
              		PoolId:   utils.String("CIDC-RP-64"),
              		Endpoint: utils.String("https://api-wuhan-1.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-xizang-2"),
              		PoolId:   utils.String("CIDC-RP-68"),
              		Endpoint: utils.String("https://api-xizang-2.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-qinghai-2"),
              		PoolId:   utils.String("CIDC-RP-69"),
              		Endpoint: utils.String("https://api-qinghai-2.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String("cn-ningxia-2"),
              		PoolId:   utils.String("CIDC-RP-71"),
              		Endpoint: utils.String("https://api-ningxia-2.cmecloud.cn:8443"),
              	})
	regions = append(regions, region.Region{
              		RegionId: utils.String(""),
              		PoolId:   utils.String("CIDC-CORE-00"),
              		Endpoint: utils.String("https://ecloud.10086.cn"),
              	})
	return regions
}

func (c *Client) setEndpoint(config *config.Config, httpRequest *request.HttpRequest) {
	if utils.IsUnSet(config.RegionId) && utils.IsUnSet(config.PoolId) {
		httpRequest.Endpoint = utils.DefaultEndpoint
	} else if utils.IsSet(config.RegionId) {
		c.findAndSetEndpointByRegionId(config.RegionId, httpRequest)
	} else {
		c.findAndSetEndpointByPoolId(config.PoolId, httpRequest)
	}
}

func (c *Client) findAndSetEndpointByRegionId(regionId *string, httpRequest *request.HttpRequest) {
	for _, r := range c.allRegions {
		if utils.StringValue(r.RegionId) == utils.StringValue(regionId) {
			endpoint := r.Endpoint
			if utils.IsUnSet(endpoint) {
				httpRequest.Endpoint = utils.DefaultEndpoint
			} else {
				httpRequest.Endpoint = *endpoint
			}
			break
		}
	}
}

func (c *Client) findAndSetEndpointByPoolId(poolId *string, httpRequest *request.HttpRequest) {
	for _, r := range c.allRegions {
		if utils.StringValue(r.PoolId) == utils.StringValue(poolId) {
			endpoint := r.Endpoint
			if utils.IsUnSet(endpoint) {
				httpRequest.Endpoint = utils.DefaultEndpoint
			} else {
				httpRequest.Endpoint = *endpoint
			}
			break
		}
	}
}

// GetAsyncResult getAsyncResult
func (c *Client) GetAsyncResult(request *model.GetAsyncResultRequest) (*model.GetAsyncResultResponse, error) {
    return c.GetAsyncResultWithConfig(request, nil)
}

// GetAsyncResultWithConfig getAsyncResult
func (c *Client) GetAsyncResultWithConfig(request *model.GetAsyncResultRequest, runtimeConfig *config.RuntimeConfig) (*model.GetAsyncResultResponse, error) {
    params := param.NewParamsBuilder().
        Action("getAsyncResult").
        Uri("/asyncResult/v1").
        GatewayUri("/api/openapi-vlb/lb-console/asyncResult/v1").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetAsyncResultResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalancListenerName 修改监听器
func (c *Client) UpdateLoadBalancListenerName(request *model.UpdateLoadBalancListenerNameRequest) (*model.UpdateLoadBalancListenerNameResponse, error) {
    return c.UpdateLoadBalancListenerNameWithConfig(request, nil)
}

// UpdateLoadBalancListenerNameWithConfig 修改监听器
func (c *Client) UpdateLoadBalancListenerNameWithConfig(request *model.UpdateLoadBalancListenerNameRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalancListenerNameResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalancListenerName").
        Uri("/acl/v3/listener").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalancListenerNameResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// QueryLoadBalanceQuota 查询弹性负载均衡配额
func (c *Client) QueryLoadBalanceQuota() (*model.QueryLoadBalanceQuotaResponse, error) {
    return c.QueryLoadBalanceQuotaWithConfig(nil)
}

// QueryLoadBalanceQuotaWithConfig 查询弹性负载均衡配额
func (c *Client) QueryLoadBalanceQuotaWithConfig(runtimeConfig *config.RuntimeConfig) (*model.QueryLoadBalanceQuotaResponse, error) {
    params := param.NewParamsBuilder().
        Action("queryLoadBalanceQuota").
        Uri("/acl/v3/loadBalancer/quota").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/quota").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.QueryLoadBalanceQuotaResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateMembers 批量添加端口至已有后端服务器
func (c *Client) CreateMembers(request *model.CreateMembersRequest) (*model.CreateMembersResponse, error) {
    return c.CreateMembersWithConfig(request, nil)
}

// CreateMembersWithConfig 批量添加端口至已有后端服务器
func (c *Client) CreateMembersWithConfig(request *model.CreateMembersRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateMembersResponse, error) {
    params := param.NewParamsBuilder().
        Action("createMembers").
        Uri("/acl/v3/member/members/{memberId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/members/{memberId}").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateMembersResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// BindTags LB绑定标签
func (c *Client) BindTags(request *model.BindTagsRequest) (*model.BindTagsResponse, error) {
    return c.BindTagsWithConfig(request, nil)
}

// BindTagsWithConfig LB绑定标签
func (c *Client) BindTagsWithConfig(request *model.BindTagsRequest, runtimeConfig *config.RuntimeConfig) (*model.BindTagsResponse, error) {
    params := param.NewParamsBuilder().
        Action("bindTags").
        Uri("/acl/v3/loadBalancer/bindTags").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/bindTags").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.BindTagsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadbalanceByTag 查询负载均衡器列表
func (c *Client) ListLoadbalanceByTag(request *model.ListLoadbalanceByTagRequest) (*model.ListLoadbalanceByTagResponse, error) {
    return c.ListLoadbalanceByTagWithConfig(request, nil)
}

// ListLoadbalanceByTagWithConfig 查询负载均衡器列表
func (c *Client) ListLoadbalanceByTagWithConfig(request *model.ListLoadbalanceByTagRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadbalanceByTagResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadbalanceByTag").
        Uri("/acl/v3/loadBalancer/loadBalancers").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/loadBalancers").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadbalanceByTagResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteTLSCustomizeProtocol 删除TLS安全策略
func (c *Client) DeleteTLSCustomizeProtocol(request *model.DeleteTLSCustomizeProtocolRequest) (*model.DeleteTLSCustomizeProtocolResponse, error) {
    return c.DeleteTLSCustomizeProtocolWithConfig(request, nil)
}

// DeleteTLSCustomizeProtocolWithConfig 删除TLS安全策略
func (c *Client) DeleteTLSCustomizeProtocolWithConfig(request *model.DeleteTLSCustomizeProtocolRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteTLSCustomizeProtocolResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteTLSCustomizeProtocol").
        Uri("/acl/v3/tls/{tlsCustomizeProtocolId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/tls/{tlsCustomizeProtocolId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteTLSCustomizeProtocolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ElbOrderDelete api退订弹性负载均衡器
func (c *Client) ElbOrderDelete(request *model.ElbOrderDeleteRequest) (*model.ElbOrderDeleteResponse, error) {
    return c.ElbOrderDeleteWithConfig(request, nil)
}

// ElbOrderDeleteWithConfig api退订弹性负载均衡器
func (c *Client) ElbOrderDeleteWithConfig(request *model.ElbOrderDeleteRequest, runtimeConfig *config.RuntimeConfig) (*model.ElbOrderDeleteResponse, error) {
    params := param.NewParamsBuilder().
        Action("elbOrderDelete").
        Uri("/acl/v3/order/loadBalancer").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/order/loadBalancer").
        Protocol("https").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.ElbOrderDeleteResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceHTTPListener 修改HTTP监听器
func (c *Client) UpdateLoadBalanceHTTPListener(request *model.UpdateLoadBalanceHTTPListenerRequest) (*model.UpdateLoadBalanceHTTPListenerResponse, error) {
    return c.UpdateLoadBalanceHTTPListenerWithConfig(request, nil)
}

// UpdateLoadBalanceHTTPListenerWithConfig 修改HTTP监听器
func (c *Client) UpdateLoadBalanceHTTPListenerWithConfig(request *model.UpdateLoadBalanceHTTPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceHTTPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceHTTPListener").
        Uri("/protocol/v3/listener/http").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/http").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceHTTPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetLoadBalanceListenerDetailResp 查询负载均衡监听器详情
func (c *Client) GetLoadBalanceListenerDetailResp(request *model.GetLoadBalanceListenerDetailRespRequest) (*model.GetLoadBalanceListenerDetailRespResponse, error) {
    return c.GetLoadBalanceListenerDetailRespWithConfig(request, nil)
}

// GetLoadBalanceListenerDetailRespWithConfig 查询负载均衡监听器详情
func (c *Client) GetLoadBalanceListenerDetailRespWithConfig(request *model.GetLoadBalanceListenerDetailRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetLoadBalanceListenerDetailRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getLoadBalanceListenerDetailResp").
        Uri("/acl/v3/listener/{listenerId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/{listenerId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetLoadBalanceListenerDetailRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadbalanceName 更改弹性负载均衡名称
func (c *Client) UpdateLoadbalanceName(request *model.UpdateLoadbalanceNameRequest) (*model.UpdateLoadbalanceNameResponse, error) {
    return c.UpdateLoadbalanceNameWithConfig(request, nil)
}

// UpdateLoadbalanceNameWithConfig 更改弹性负载均衡名称
func (c *Client) UpdateLoadbalanceNameWithConfig(request *model.UpdateLoadbalanceNameRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadbalanceNameResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadbalanceName").
        Uri("/acl/v3/loadBalancer/name").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/name").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadbalanceNameResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UnbindFip 解绑公网IP
func (c *Client) UnbindFip(request *model.UnbindFipRequest) (*model.UnbindFipResponse, error) {
    return c.UnbindFipWithConfig(request, nil)
}

// UnbindFipWithConfig 解绑公网IP
func (c *Client) UnbindFipWithConfig(request *model.UnbindFipRequest, runtimeConfig *config.RuntimeConfig) (*model.UnbindFipResponse, error) {
    params := param.NewParamsBuilder().
        Action("unbindFip").
        Uri("/acl/v3/loadBalancer/unbindfIp/{loadBalanceId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/unbindfIp/{loadBalanceId}").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UnbindFipResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateLoadBalanceHTTPSListener 创建HTTPS监听器
func (c *Client) CreateLoadBalanceHTTPSListener(request *model.CreateLoadBalanceHTTPSListenerRequest) (*model.CreateLoadBalanceHTTPSListenerResponse, error) {
    return c.CreateLoadBalanceHTTPSListenerWithConfig(request, nil)
}

// CreateLoadBalanceHTTPSListenerWithConfig 创建HTTPS监听器
func (c *Client) CreateLoadBalanceHTTPSListenerWithConfig(request *model.CreateLoadBalanceHTTPSListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateLoadBalanceHTTPSListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("createLoadBalanceHTTPSListener").
        Uri("/protocol/v3/listener/https").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/https").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateLoadBalanceHTTPSListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListControlGroup 查询访问控制组列表
func (c *Client) ListControlGroup(request *model.ListControlGroupRequest) (*model.ListControlGroupResponse, error) {
    return c.ListControlGroupWithConfig(request, nil)
}

// ListControlGroupWithConfig 查询访问控制组列表
func (c *Client) ListControlGroupWithConfig(request *model.ListControlGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.ListControlGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("listControlGroup").
        Uri("/acl/v3/accessControlGroup").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListControlGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceHTTPSListenerPool 修改HTTPS监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceHTTPSListenerPool(request *model.UpdateLoadBalanceHTTPSListenerPoolRequest) (*model.UpdateLoadBalanceHTTPSListenerPoolResponse, error) {
    return c.UpdateLoadBalanceHTTPSListenerPoolWithConfig(request, nil)
}

// UpdateLoadBalanceHTTPSListenerPoolWithConfig 修改HTTPS监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceHTTPSListenerPoolWithConfig(request *model.UpdateLoadBalanceHTTPSListenerPoolRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceHTTPSListenerPoolResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceHTTPSListenerPool").
        Uri("/protocol/v3/listener/https/updatePool").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/https/updatePool").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceHTTPSListenerPoolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// AsyncCreatePool 启用异步线程创建后端服务器组（包括HealthMonitor）
func (c *Client) AsyncCreatePool(request *model.AsyncCreatePoolRequest) (*model.AsyncCreatePoolResponse, error) {
    return c.AsyncCreatePoolWithConfig(request, nil)
}

// AsyncCreatePoolWithConfig 启用异步线程创建后端服务器组（包括HealthMonitor）
func (c *Client) AsyncCreatePoolWithConfig(request *model.AsyncCreatePoolRequest, runtimeConfig *config.RuntimeConfig) (*model.AsyncCreatePoolResponse, error) {
    params := param.NewParamsBuilder().
        Action("asyncCreatePool").
        Uri("/acl/v3/pool").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AsyncCreatePoolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadBalanceHTTPListener 查询HTTP监听器
func (c *Client) ListLoadBalanceHTTPListener(request *model.ListLoadBalanceHTTPListenerRequest) (*model.ListLoadBalanceHTTPListenerResponse, error) {
    return c.ListLoadBalanceHTTPListenerWithConfig(request, nil)
}

// ListLoadBalanceHTTPListenerWithConfig 查询HTTP监听器
func (c *Client) ListLoadBalanceHTTPListenerWithConfig(request *model.ListLoadBalanceHTTPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadBalanceHTTPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadBalanceHTTPListener").
        Uri("/protocol/v3/listener/{loadBalanceId}/listeners/http").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/{loadBalanceId}/listeners/http").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadBalanceHTTPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadBalancePoolMember 查询负载均衡后端服务器列表
func (c *Client) ListLoadBalancePoolMember(request *model.ListLoadBalancePoolMemberRequest) (*model.ListLoadBalancePoolMemberResponse, error) {
    return c.ListLoadBalancePoolMemberWithConfig(request, nil)
}

// ListLoadBalancePoolMemberWithConfig 查询负载均衡后端服务器列表
func (c *Client) ListLoadBalancePoolMemberWithConfig(request *model.ListLoadBalancePoolMemberRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadBalancePoolMemberResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadBalancePoolMember").
        Uri("/acl/v3/member/{poolId}/listMembers").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{poolId}/listMembers").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadBalancePoolMemberResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceTCPListener 修改TCP监听器
func (c *Client) UpdateLoadBalanceTCPListener(request *model.UpdateLoadBalanceTCPListenerRequest) (*model.UpdateLoadBalanceTCPListenerResponse, error) {
    return c.UpdateLoadBalanceTCPListenerWithConfig(request, nil)
}

// UpdateLoadBalanceTCPListenerWithConfig 修改TCP监听器
func (c *Client) UpdateLoadBalanceTCPListenerWithConfig(request *model.UpdateLoadBalanceTCPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceTCPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceTCPListener").
        Uri("/protocol/v3/listener/tcp").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/tcp").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceTCPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ElbOrderCreateAsync api订购弹性负载均衡器
func (c *Client) ElbOrderCreateAsync(request *model.ElbOrderCreateAsyncRequest) (*model.ElbOrderCreateAsyncResponse, error) {
    return c.ElbOrderCreateAsyncWithConfig(request, nil)
}

// ElbOrderCreateAsyncWithConfig api订购弹性负载均衡器
func (c *Client) ElbOrderCreateAsyncWithConfig(request *model.ElbOrderCreateAsyncRequest, runtimeConfig *config.RuntimeConfig) (*model.ElbOrderCreateAsyncResponse, error) {
    params := param.NewParamsBuilder().
        Action("elbOrderCreateAsync").
        Uri("/acl/v3/order/loadBalancer").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/order/loadBalancer").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ElbOrderCreateAsyncResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeletePoolAsync 异步删除后端服务器组
func (c *Client) DeletePoolAsync(request *model.DeletePoolAsyncRequest) (*model.DeletePoolAsyncResponse, error) {
    return c.DeletePoolAsyncWithConfig(request, nil)
}

// DeletePoolAsyncWithConfig 异步删除后端服务器组
func (c *Client) DeletePoolAsyncWithConfig(request *model.DeletePoolAsyncRequest, runtimeConfig *config.RuntimeConfig) (*model.DeletePoolAsyncResponse, error) {
    params := param.NewParamsBuilder().
        Action("deletePoolAsync").
        Uri("/acl/v3/pool/deleteAsync/{poolId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool/deleteAsync/{poolId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeletePoolAsyncResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteListenerAsync 异步删除负载均衡监听器
func (c *Client) DeleteListenerAsync(request *model.DeleteListenerAsyncRequest) (*model.DeleteListenerAsyncResponse, error) {
    return c.DeleteListenerAsyncWithConfig(request, nil)
}

// DeleteListenerAsyncWithConfig 异步删除负载均衡监听器
func (c *Client) DeleteListenerAsyncWithConfig(request *model.DeleteListenerAsyncRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteListenerAsyncResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteListenerAsync").
        Uri("/acl/v3/listener/deleteAsync/{listenerId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/deleteAsync/{listenerId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteListenerAsyncResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetPoolResp 查询后端服务器组详情
func (c *Client) GetPoolResp(request *model.GetPoolRespRequest) (*model.GetPoolRespResponse, error) {
    return c.GetPoolRespWithConfig(request, nil)
}

// GetPoolRespWithConfig 查询后端服务器组详情
func (c *Client) GetPoolRespWithConfig(request *model.GetPoolRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetPoolRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getPoolResp").
        Uri("/acl/v3/pool/{poolId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool/{poolId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetPoolRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceTCPListenerPool 修改TCP监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceTCPListenerPool(request *model.UpdateLoadBalanceTCPListenerPoolRequest) (*model.UpdateLoadBalanceTCPListenerPoolResponse, error) {
    return c.UpdateLoadBalanceTCPListenerPoolWithConfig(request, nil)
}

// UpdateLoadBalanceTCPListenerPoolWithConfig 修改TCP监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceTCPListenerPoolWithConfig(request *model.UpdateLoadBalanceTCPListenerPoolRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceTCPListenerPoolResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceTCPListenerPool").
        Uri("/protocol/v3/listener/tcp/updatePool").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/tcp/updatePool").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceTCPListenerPoolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceHTTPSListener 修改HTTPS监听器
func (c *Client) UpdateLoadBalanceHTTPSListener(request *model.UpdateLoadBalanceHTTPSListenerRequest) (*model.UpdateLoadBalanceHTTPSListenerResponse, error) {
    return c.UpdateLoadBalanceHTTPSListenerWithConfig(request, nil)
}

// UpdateLoadBalanceHTTPSListenerWithConfig 修改HTTPS监听器
func (c *Client) UpdateLoadBalanceHTTPSListenerWithConfig(request *model.UpdateLoadBalanceHTTPSListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceHTTPSListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceHTTPSListener").
        Uri("/protocol/v3/listener/https").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/https").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceHTTPSListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ElbOrderCreateAsyncWithHttpStatus 订购弹性负载均衡器(修复HTTP状态码)
func (c *Client) ElbOrderCreateAsyncWithHttpStatus(request *model.ElbOrderCreateAsyncWithHttpStatusRequest) (*model.ElbOrderCreateAsyncWithHttpStatusResponse, error) {
    return c.ElbOrderCreateAsyncWithHttpStatusWithConfig(request, nil)
}

// ElbOrderCreateAsyncWithHttpStatusWithConfig 订购弹性负载均衡器(修复HTTP状态码)
func (c *Client) ElbOrderCreateAsyncWithHttpStatusWithConfig(request *model.ElbOrderCreateAsyncWithHttpStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.ElbOrderCreateAsyncWithHttpStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("elbOrderCreateAsyncWithHttpStatus").
        Uri("/acl/v3/order/create/loadBalancer").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/order/create/loadBalancer").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ElbOrderCreateAsyncWithHttpStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteControlGroup 删除访问控制策略以及组内的所有条目
func (c *Client) DeleteControlGroup(request *model.DeleteControlGroupRequest) (*model.DeleteControlGroupResponse, error) {
    return c.DeleteControlGroupWithConfig(request, nil)
}

// DeleteControlGroupWithConfig 删除访问控制策略以及组内的所有条目
func (c *Client) DeleteControlGroupWithConfig(request *model.DeleteControlGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteControlGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteControlGroup").
        Uri("/acl/v3/accessControlGroup/{accessControlGroupId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup/{accessControlGroupId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteControlGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// AsyncCreateWithHttpStatus 创建负载均衡监听器(优化HTTP状态码)
func (c *Client) AsyncCreateWithHttpStatus(request *model.AsyncCreateWithHttpStatusRequest) (*model.AsyncCreateWithHttpStatusResponse, error) {
    return c.AsyncCreateWithHttpStatusWithConfig(request, nil)
}

// AsyncCreateWithHttpStatusWithConfig 创建负载均衡监听器(优化HTTP状态码)
func (c *Client) AsyncCreateWithHttpStatusWithConfig(request *model.AsyncCreateWithHttpStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.AsyncCreateWithHttpStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("asyncCreateWithHttpStatus").
        Uri("/acl/v3/listener/create").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/create").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AsyncCreateWithHttpStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceUDPListener 修改UDP监听器
func (c *Client) UpdateLoadBalanceUDPListener(request *model.UpdateLoadBalanceUDPListenerRequest) (*model.UpdateLoadBalanceUDPListenerResponse, error) {
    return c.UpdateLoadBalanceUDPListenerWithConfig(request, nil)
}

// UpdateLoadBalanceUDPListenerWithConfig 修改UDP监听器
func (c *Client) UpdateLoadBalanceUDPListenerWithConfig(request *model.UpdateLoadBalanceUDPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceUDPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceUDPListener").
        Uri("/protocol/v3/listener/udp").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/udp").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceUDPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListControlGroupDetail 查询访问控制组详情
func (c *Client) ListControlGroupDetail(request *model.ListControlGroupDetailRequest) (*model.ListControlGroupDetailResponse, error) {
    return c.ListControlGroupDetailWithConfig(request, nil)
}

// ListControlGroupDetailWithConfig 查询访问控制组详情
func (c *Client) ListControlGroupDetailWithConfig(request *model.ListControlGroupDetailRequest, runtimeConfig *config.RuntimeConfig) (*model.ListControlGroupDetailResponse, error) {
    params := param.NewParamsBuilder().
        Action("listControlGroupDetail").
        Uri("/acl/v3/accessControlGroup/{accessControlGroupId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup/{accessControlGroupId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListControlGroupDetailResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetTLSCustomizeProtocolResp 查询TLS安全策略详情
func (c *Client) GetTLSCustomizeProtocolResp(request *model.GetTLSCustomizeProtocolRespRequest) (*model.GetTLSCustomizeProtocolRespResponse, error) {
    return c.GetTLSCustomizeProtocolRespWithConfig(request, nil)
}

// GetTLSCustomizeProtocolRespWithConfig 查询TLS安全策略详情
func (c *Client) GetTLSCustomizeProtocolRespWithConfig(request *model.GetTLSCustomizeProtocolRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetTLSCustomizeProtocolRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getTLSCustomizeProtocolResp").
        Uri("/acl/v3/tls/{tlsCustomizeProtocolId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/tls/{tlsCustomizeProtocolId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetTLSCustomizeProtocolRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetLoadBalancePoolMemberResp 查询负载均衡后端服务器详情
func (c *Client) GetLoadBalancePoolMemberResp(request *model.GetLoadBalancePoolMemberRespRequest) (*model.GetLoadBalancePoolMemberRespResponse, error) {
    return c.GetLoadBalancePoolMemberRespWithConfig(request, nil)
}

// GetLoadBalancePoolMemberRespWithConfig 查询负载均衡后端服务器详情
func (c *Client) GetLoadBalancePoolMemberRespWithConfig(request *model.GetLoadBalancePoolMemberRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetLoadBalancePoolMemberRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getLoadBalancePoolMemberResp").
        Uri("/acl/v3/member/{memberId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{memberId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetLoadBalancePoolMemberRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateControlGroupRule 创建访问控制组规则
func (c *Client) CreateControlGroupRule(request *model.CreateControlGroupRuleRequest) (*model.CreateControlGroupRuleResponse, error) {
    return c.CreateControlGroupRuleWithConfig(request, nil)
}

// CreateControlGroupRuleWithConfig 创建访问控制组规则
func (c *Client) CreateControlGroupRuleWithConfig(request *model.CreateControlGroupRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateControlGroupRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("createControlGroupRule").
        Uri("/acl/v3/accessControlGroupRule").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroupRule").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateControlGroupRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ValidateControlGroup 删除校验访问控制组
func (c *Client) ValidateControlGroup(request *model.ValidateControlGroupRequest) (*model.ValidateControlGroupResponse, error) {
    return c.ValidateControlGroupWithConfig(request, nil)
}

// ValidateControlGroupWithConfig 删除校验访问控制组
func (c *Client) ValidateControlGroupWithConfig(request *model.ValidateControlGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.ValidateControlGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("validateControlGroup").
        Uri("/acl/v3/accessControlGroup/deletable/{accessControlGroupId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup/deletable/{accessControlGroupId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ValidateControlGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetLoadbalanceExpirationCertification 获取负载均衡过期证书
func (c *Client) GetLoadbalanceExpirationCertification() (*model.GetLoadbalanceExpirationCertificationResponse, error) {
    return c.GetLoadbalanceExpirationCertificationWithConfig(nil)
}

// GetLoadbalanceExpirationCertificationWithConfig 获取负载均衡过期证书
func (c *Client) GetLoadbalanceExpirationCertificationWithConfig(runtimeConfig *config.RuntimeConfig) (*model.GetLoadbalanceExpirationCertificationResponse, error) {
    params := param.NewParamsBuilder().
        Action("getLoadbalanceExpirationCertification").
        Uri("/acl/v3/certification/expiration").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/certification/expiration").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.GetLoadbalanceExpirationCertificationResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceHTTPListenerPool 修改HTTP监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceHTTPListenerPool(request *model.UpdateLoadBalanceHTTPListenerPoolRequest) (*model.UpdateLoadBalanceHTTPListenerPoolResponse, error) {
    return c.UpdateLoadBalanceHTTPListenerPoolWithConfig(request, nil)
}

// UpdateLoadBalanceHTTPListenerPoolWithConfig 修改HTTP监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceHTTPListenerPoolWithConfig(request *model.UpdateLoadBalanceHTTPListenerPoolRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceHTTPListenerPoolResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceHTTPListenerPool").
        Uri("/protocol/v3/listener/http/updatePool").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/http/updatePool").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceHTTPListenerPoolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateControlGroupRule 更新访问控制组规则
func (c *Client) UpdateControlGroupRule(request *model.UpdateControlGroupRuleRequest) (*model.UpdateControlGroupRuleResponse, error) {
    return c.UpdateControlGroupRuleWithConfig(request, nil)
}

// UpdateControlGroupRuleWithConfig 更新访问控制组规则
func (c *Client) UpdateControlGroupRuleWithConfig(request *model.UpdateControlGroupRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateControlGroupRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateControlGroupRule").
        Uri("/acl/v3/accessControlGroupRule").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroupRule").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateControlGroupRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdatePoolName 启用异步线程修改后端服务器组
func (c *Client) UpdatePoolName(request *model.UpdatePoolNameRequest) (*model.UpdatePoolNameResponse, error) {
    return c.UpdatePoolNameWithConfig(request, nil)
}

// UpdatePoolNameWithConfig 启用异步线程修改后端服务器组
func (c *Client) UpdatePoolNameWithConfig(request *model.UpdatePoolNameRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdatePoolNameResponse, error) {
    params := param.NewParamsBuilder().
        Action("updatePoolName").
        Uri("/acl/v3/pool").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdatePoolNameResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteLoadBalancePoolMember 删除负载均衡后端服务器
func (c *Client) DeleteLoadBalancePoolMember(request *model.DeleteLoadBalancePoolMemberRequest) (*model.DeleteLoadBalancePoolMemberResponse, error) {
    return c.DeleteLoadBalancePoolMemberWithConfig(request, nil)
}

// DeleteLoadBalancePoolMemberWithConfig 删除负载均衡后端服务器
func (c *Client) DeleteLoadBalancePoolMemberWithConfig(request *model.DeleteLoadBalancePoolMemberRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteLoadBalancePoolMemberResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteLoadBalancePoolMember").
        Uri("/acl/v3/member/{poolId}/member/{memberId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{poolId}/member/{memberId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteLoadBalancePoolMemberResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ElbOrderRenew api续订弹性负载均衡器
func (c *Client) ElbOrderRenew(request *model.ElbOrderRenewRequest) (*model.ElbOrderRenewResponse, error) {
    return c.ElbOrderRenewWithConfig(request, nil)
}

// ElbOrderRenewWithConfig api续订弹性负载均衡器
func (c *Client) ElbOrderRenewWithConfig(request *model.ElbOrderRenewRequest, runtimeConfig *config.RuntimeConfig) (*model.ElbOrderRenewResponse, error) {
    params := param.NewParamsBuilder().
        Action("elbOrderRenew").
        Uri("/acl/v3/order/loadBalancer/renew").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/order/loadBalancer/renew").
        Protocol("https").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.ElbOrderRenewResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadBalanceUDPListener 查询UDP监听器
func (c *Client) ListLoadBalanceUDPListener(request *model.ListLoadBalanceUDPListenerRequest) (*model.ListLoadBalanceUDPListenerResponse, error) {
    return c.ListLoadBalanceUDPListenerWithConfig(request, nil)
}

// ListLoadBalanceUDPListenerWithConfig 查询UDP监听器
func (c *Client) ListLoadBalanceUDPListenerWithConfig(request *model.ListLoadBalanceUDPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadBalanceUDPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadBalanceUDPListener").
        Uri("/protocol/v3/listener/{loadBalanceId}/listeners/udp").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/{loadBalanceId}/listeners/udp").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadBalanceUDPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateAliAcl 创建阿里域访问控制组
func (c *Client) CreateAliAcl(request *model.CreateAliAclRequest) (*model.CreateAliAclResponse, error) {
    return c.CreateAliAclWithConfig(request, nil)
}

// CreateAliAclWithConfig 创建阿里域访问控制组
func (c *Client) CreateAliAclWithConfig(request *model.CreateAliAclRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateAliAclResponse, error) {
    params := param.NewParamsBuilder().
        Action("createAliAcl").
        Uri("/acl/v3/accessControlGroup/ALI").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup/ALI").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateAliAclResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadBalanceTCPListener 查询TCP监听器
func (c *Client) ListLoadBalanceTCPListener(request *model.ListLoadBalanceTCPListenerRequest) (*model.ListLoadBalanceTCPListenerResponse, error) {
    return c.ListLoadBalanceTCPListenerWithConfig(request, nil)
}

// ListLoadBalanceTCPListenerWithConfig 查询TCP监听器
func (c *Client) ListLoadBalanceTCPListenerWithConfig(request *model.ListLoadBalanceTCPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadBalanceTCPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadBalanceTCPListener").
        Uri("/protocol/v3/listener/{loadBalanceId}/listeners/tcp").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/{loadBalanceId}/listeners/tcp").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadBalanceTCPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateTCAcl 创建NFV域访问控制组
func (c *Client) CreateTCAcl(request *model.CreateTCAclRequest) (*model.CreateTCAclResponse, error) {
    return c.CreateTCAclWithConfig(request, nil)
}

// CreateTCAclWithConfig 创建NFV域访问控制组
func (c *Client) CreateTCAclWithConfig(request *model.CreateTCAclRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateTCAclResponse, error) {
    params := param.NewParamsBuilder().
        Action("createTCAcl").
        Uri("/acl/v3/accessControlGroup/NFV").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup/NFV").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateTCAclResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListTLSCustomizeProtocolResp 查询TLS安全策略列表
func (c *Client) ListTLSCustomizeProtocolResp(request *model.ListTLSCustomizeProtocolRespRequest) (*model.ListTLSCustomizeProtocolRespResponse, error) {
    return c.ListTLSCustomizeProtocolRespWithConfig(request, nil)
}

// ListTLSCustomizeProtocolRespWithConfig 查询TLS安全策略列表
func (c *Client) ListTLSCustomizeProtocolRespWithConfig(request *model.ListTLSCustomizeProtocolRespRequest, runtimeConfig *config.RuntimeConfig) (*model.ListTLSCustomizeProtocolRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("listTLSCustomizeProtocolResp").
        Uri("/acl/v3/tls").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/tls").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListTLSCustomizeProtocolRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// AsyncUpdateWithHttpStatus 修改负载均衡监听器(优化HTTP状态码)
func (c *Client) AsyncUpdateWithHttpStatus(request *model.AsyncUpdateWithHttpStatusRequest) (*model.AsyncUpdateWithHttpStatusResponse, error) {
    return c.AsyncUpdateWithHttpStatusWithConfig(request, nil)
}

// AsyncUpdateWithHttpStatusWithConfig 修改负载均衡监听器(优化HTTP状态码)
func (c *Client) AsyncUpdateWithHttpStatusWithConfig(request *model.AsyncUpdateWithHttpStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.AsyncUpdateWithHttpStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("asyncUpdateWithHttpStatus").
        Uri("/acl/v3/listener/update").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/update").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.AsyncUpdateWithHttpStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateTLSCustomizeProtocol 更新TLS安全策略
func (c *Client) UpdateTLSCustomizeProtocol(request *model.UpdateTLSCustomizeProtocolRequest) (*model.UpdateTLSCustomizeProtocolResponse, error) {
    return c.UpdateTLSCustomizeProtocolWithConfig(request, nil)
}

// UpdateTLSCustomizeProtocolWithConfig 更新TLS安全策略
func (c *Client) UpdateTLSCustomizeProtocolWithConfig(request *model.UpdateTLSCustomizeProtocolRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateTLSCustomizeProtocolResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateTLSCustomizeProtocol").
        Uri("/acl/v3/tls").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/tls").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateTLSCustomizeProtocolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// AsyncCreatePolicy 创建七层转发策略
func (c *Client) AsyncCreatePolicy(request *model.AsyncCreatePolicyRequest) (*model.AsyncCreatePolicyResponse, error) {
    return c.AsyncCreatePolicyWithConfig(request, nil)
}

// AsyncCreatePolicyWithConfig 创建七层转发策略
func (c *Client) AsyncCreatePolicyWithConfig(request *model.AsyncCreatePolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.AsyncCreatePolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("asyncCreatePolicy").
        Uri("/acl/v3/l7Policy").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/l7Policy").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.AsyncCreatePolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// EnableLoadBalancer 启用停用弹性负载均衡
func (c *Client) EnableLoadBalancer(request *model.EnableLoadBalancerRequest) (*model.EnableLoadBalancerResponse, error) {
    return c.EnableLoadBalancerWithConfig(request, nil)
}

// EnableLoadBalancerWithConfig 启用停用弹性负载均衡
func (c *Client) EnableLoadBalancerWithConfig(request *model.EnableLoadBalancerRequest, runtimeConfig *config.RuntimeConfig) (*model.EnableLoadBalancerResponse, error) {
    params := param.NewParamsBuilder().
        Action("enableLoadBalancer").
        Uri("/acl/v3/loadBalancer/enable/{loadBalanceId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/enable/{loadBalanceId}").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.EnableLoadBalancerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateLoadBalanceListenerAsync 启用异步线程创建负载均衡监听器
func (c *Client) CreateLoadBalanceListenerAsync(request *model.CreateLoadBalanceListenerAsyncRequest) (*model.CreateLoadBalanceListenerAsyncResponse, error) {
    return c.CreateLoadBalanceListenerAsyncWithConfig(request, nil)
}

// CreateLoadBalanceListenerAsyncWithConfig 启用异步线程创建负载均衡监听器
func (c *Client) CreateLoadBalanceListenerAsyncWithConfig(request *model.CreateLoadBalanceListenerAsyncRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateLoadBalanceListenerAsyncResponse, error) {
    params := param.NewParamsBuilder().
        Action("createLoadBalanceListenerAsync").
        Uri("/acl/v3/listener").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateLoadBalanceListenerAsyncResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetBatchHealthStatusResp 查询负载均衡后端服务器的健康检查状态
func (c *Client) GetBatchHealthStatusResp(request *model.GetBatchHealthStatusRespRequest) (*model.GetBatchHealthStatusRespResponse, error) {
    return c.GetBatchHealthStatusRespWithConfig(request, nil)
}

// GetBatchHealthStatusRespWithConfig 查询负载均衡后端服务器的健康检查状态
func (c *Client) GetBatchHealthStatusRespWithConfig(request *model.GetBatchHealthStatusRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetBatchHealthStatusRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getBatchHealthStatusResp").
        Uri("/acl/v3/member/batchHealthStatus/{poolId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/batchHealthStatus/{poolId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetBatchHealthStatusRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalancePoolMember 更新负载均member
func (c *Client) UpdateLoadBalancePoolMember(request *model.UpdateLoadBalancePoolMemberRequest) (*model.UpdateLoadBalancePoolMemberResponse, error) {
    return c.UpdateLoadBalancePoolMemberWithConfig(request, nil)
}

// UpdateLoadBalancePoolMemberWithConfig 更新负载均member
func (c *Client) UpdateLoadBalancePoolMemberWithConfig(request *model.UpdateLoadBalancePoolMemberRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalancePoolMemberResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalancePoolMember").
        Uri("/acl/v3/member/{poolId}/member/{memberId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{poolId}/member/{memberId}").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalancePoolMemberResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CertificationBindTags 证书绑定标签
func (c *Client) CertificationBindTags(request *model.CertificationBindTagsRequest) (*model.CertificationBindTagsResponse, error) {
    return c.CertificationBindTagsWithConfig(request, nil)
}

// CertificationBindTagsWithConfig 证书绑定标签
func (c *Client) CertificationBindTagsWithConfig(request *model.CertificationBindTagsRequest, runtimeConfig *config.RuntimeConfig) (*model.CertificationBindTagsResponse, error) {
    params := param.NewParamsBuilder().
        Action("certificationBindTags").
        Uri("/customer/v3/certification/bindTags").
        GatewayUri("/api/openapi-vlb/lb-console-openstack-lb-barbican/customer/v3/certification/bindTags").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CertificationBindTagsResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetL7PolicyResp 查询七层转发策略详情
func (c *Client) GetL7PolicyResp(request *model.GetL7PolicyRespRequest) (*model.GetL7PolicyRespResponse, error) {
    return c.GetL7PolicyRespWithConfig(request, nil)
}

// GetL7PolicyRespWithConfig 查询七层转发策略详情
func (c *Client) GetL7PolicyRespWithConfig(request *model.GetL7PolicyRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetL7PolicyRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getL7PolicyResp").
        Uri("/acl/v3/l7Policy/{l7PolicyId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/l7Policy/{l7PolicyId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetL7PolicyRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadBalanceListenerResp 查询负载均衡监听器列表
func (c *Client) ListLoadBalanceListenerResp(request *model.ListLoadBalanceListenerRespRequest) (*model.ListLoadBalanceListenerRespResponse, error) {
    return c.ListLoadBalanceListenerRespWithConfig(request, nil)
}

// ListLoadBalanceListenerRespWithConfig 查询负载均衡监听器列表
func (c *Client) ListLoadBalanceListenerRespWithConfig(request *model.ListLoadBalanceListenerRespRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadBalanceListenerRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadBalanceListenerResp").
        Uri("/acl/v3/listener/{loadBalanceId}/listeners").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/{loadBalanceId}/listeners").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadBalanceListenerRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// AsyncUpdateL7Policy 修改七层转发策略
func (c *Client) AsyncUpdateL7Policy(request *model.AsyncUpdateL7PolicyRequest) (*model.AsyncUpdateL7PolicyResponse, error) {
    return c.AsyncUpdateL7PolicyWithConfig(request, nil)
}

// AsyncUpdateL7PolicyWithConfig 修改七层转发策略
func (c *Client) AsyncUpdateL7PolicyWithConfig(request *model.AsyncUpdateL7PolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.AsyncUpdateL7PolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("asyncUpdateL7Policy").
        Uri("/acl/v3/l7Policy").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/l7Policy").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.AsyncUpdateL7PolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateTLSName 更新TLS安全策略
func (c *Client) UpdateTLSName(request *model.UpdateTLSNameRequest) (*model.UpdateTLSNameResponse, error) {
    return c.UpdateTLSNameWithConfig(request, nil)
}

// UpdateTLSNameWithConfig 更新TLS安全策略
func (c *Client) UpdateTLSNameWithConfig(request *model.UpdateTLSNameRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateTLSNameResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateTLSName").
        Uri("/acl/v3/tls").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/tls").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateTLSNameResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListL7PolicyResp 查询七层转发策略列表
func (c *Client) ListL7PolicyResp(request *model.ListL7PolicyRespRequest) (*model.ListL7PolicyRespResponse, error) {
    return c.ListL7PolicyRespWithConfig(request, nil)
}

// ListL7PolicyRespWithConfig 查询七层转发策略列表
func (c *Client) ListL7PolicyRespWithConfig(request *model.ListL7PolicyRespRequest, runtimeConfig *config.RuntimeConfig) (*model.ListL7PolicyRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("listL7PolicyResp").
        Uri("/acl/v3/l7Policy/l7Policies").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/l7Policy/l7Policies").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListL7PolicyRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ElbOrderDeleteWithHttpStatus 退订弹性负载均衡器(修复HTTP状态码)
func (c *Client) ElbOrderDeleteWithHttpStatus(request *model.ElbOrderDeleteWithHttpStatusRequest) (*model.ElbOrderDeleteWithHttpStatusResponse, error) {
    return c.ElbOrderDeleteWithHttpStatusWithConfig(request, nil)
}

// ElbOrderDeleteWithHttpStatusWithConfig 退订弹性负载均衡器(修复HTTP状态码)
func (c *Client) ElbOrderDeleteWithHttpStatusWithConfig(request *model.ElbOrderDeleteWithHttpStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.ElbOrderDeleteWithHttpStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("elbOrderDeleteWithHttpStatus").
        Uri("/acl/v3/order/delete/loadBalancer").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/order/delete/loadBalancer").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.ElbOrderDeleteWithHttpStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListNewServerPortsForMember 可添加的业务主机（云主机、物理机）端口列表
func (c *Client) ListNewServerPortsForMember(request *model.ListNewServerPortsForMemberRequest) (*model.ListNewServerPortsForMemberResponse, error) {
    return c.ListNewServerPortsForMemberWithConfig(request, nil)
}

// ListNewServerPortsForMemberWithConfig 可添加的业务主机（云主机、物理机）端口列表
func (c *Client) ListNewServerPortsForMemberWithConfig(request *model.ListNewServerPortsForMemberRequest, runtimeConfig *config.RuntimeConfig) (*model.ListNewServerPortsForMemberResponse, error) {
    params := param.NewParamsBuilder().
        Action("listNewServerPortsForMember").
        Uri("/acl/v3/member/serverPorts").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/serverPorts").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListNewServerPortsForMemberResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteL7Policy 删除七层转发策略
func (c *Client) DeleteL7Policy(request *model.DeleteL7PolicyRequest) (*model.DeleteL7PolicyResponse, error) {
    return c.DeleteL7PolicyWithConfig(request, nil)
}

// DeleteL7PolicyWithConfig 删除七层转发策略
func (c *Client) DeleteL7PolicyWithConfig(request *model.DeleteL7PolicyRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteL7PolicyResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteL7Policy").
        Uri("/acl/v3/l7Policy/{l7PolicyId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/l7Policy/{l7PolicyId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteL7PolicyResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadbalanceCertificationResp 查询负载均衡证书列表
func (c *Client) ListLoadbalanceCertificationResp() (*model.ListLoadbalanceCertificationRespResponse, error) {
    return c.ListLoadbalanceCertificationRespWithConfig(nil)
}

// ListLoadbalanceCertificationRespWithConfig 查询负载均衡证书列表
func (c *Client) ListLoadbalanceCertificationRespWithConfig(runtimeConfig *config.RuntimeConfig) (*model.ListLoadbalanceCertificationRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadbalanceCertificationResp").
        Uri("/acl/v3/certification").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/certification").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Build()
    returnValue := &model.ListLoadbalanceCertificationRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateLoadBalanceHTTPListener 创建HTTP监听器
func (c *Client) CreateLoadBalanceHTTPListener(request *model.CreateLoadBalanceHTTPListenerRequest) (*model.CreateLoadBalanceHTTPListenerResponse, error) {
    return c.CreateLoadBalanceHTTPListenerWithConfig(request, nil)
}

// CreateLoadBalanceHTTPListenerWithConfig 创建HTTP监听器
func (c *Client) CreateLoadBalanceHTTPListenerWithConfig(request *model.CreateLoadBalanceHTTPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateLoadBalanceHTTPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("createLoadBalanceHTTPListener").
        Uri("/protocol/v3/listener/http").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/http").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateLoadBalanceHTTPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeletePolicyAsync 异步删除七层转发策略
func (c *Client) DeletePolicyAsync(request *model.DeletePolicyAsyncRequest) (*model.DeletePolicyAsyncResponse, error) {
    return c.DeletePolicyAsyncWithConfig(request, nil)
}

// DeletePolicyAsyncWithConfig 异步删除七层转发策略
func (c *Client) DeletePolicyAsyncWithConfig(request *model.DeletePolicyAsyncRequest, runtimeConfig *config.RuntimeConfig) (*model.DeletePolicyAsyncResponse, error) {
    params := param.NewParamsBuilder().
        Action("deletePolicyAsync").
        Uri("/acl/v3/l7Policy/deleteAsync/{l7PolicyId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/l7Policy/deleteAsync/{l7PolicyId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeletePolicyAsyncResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListPoolResp 后端服务器组列表
func (c *Client) ListPoolResp(request *model.ListPoolRespRequest) (*model.ListPoolRespResponse, error) {
    return c.ListPoolRespWithConfig(request, nil)
}

// ListPoolRespWithConfig 后端服务器组列表
func (c *Client) ListPoolRespWithConfig(request *model.ListPoolRespRequest, runtimeConfig *config.RuntimeConfig) (*model.ListPoolRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("listPoolResp").
        Uri("/acl/v3/pool/{loadBalanceId}/pools").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool/{loadBalanceId}/pools").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListPoolRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateLoadBalanceTCPListener 创建TCP监听器
func (c *Client) CreateLoadBalanceTCPListener(request *model.CreateLoadBalanceTCPListenerRequest) (*model.CreateLoadBalanceTCPListenerResponse, error) {
    return c.CreateLoadBalanceTCPListenerWithConfig(request, nil)
}

// CreateLoadBalanceTCPListenerWithConfig 创建TCP监听器
func (c *Client) CreateLoadBalanceTCPListenerWithConfig(request *model.CreateLoadBalanceTCPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateLoadBalanceTCPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("createLoadBalanceTCPListener").
        Uri("/protocol/v3/listener/tcp").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/tcp").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateLoadBalanceTCPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeletePool 删除后端服务器组
func (c *Client) DeletePool(request *model.DeletePoolRequest) (*model.DeletePoolResponse, error) {
    return c.DeletePoolWithConfig(request, nil)
}

// DeletePoolWithConfig 删除后端服务器组
func (c *Client) DeletePoolWithConfig(request *model.DeletePoolRequest, runtimeConfig *config.RuntimeConfig) (*model.DeletePoolResponse, error) {
    params := param.NewParamsBuilder().
        Action("deletePool").
        Uri("/acl/v3/pool/{poolId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool/{poolId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeletePoolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadBalanceHTTPSListener 查询HTTPS监听器
func (c *Client) ListLoadBalanceHTTPSListener(request *model.ListLoadBalanceHTTPSListenerRequest) (*model.ListLoadBalanceHTTPSListenerResponse, error) {
    return c.ListLoadBalanceHTTPSListenerWithConfig(request, nil)
}

// ListLoadBalanceHTTPSListenerWithConfig 查询HTTPS监听器
func (c *Client) ListLoadBalanceHTTPSListenerWithConfig(request *model.ListLoadBalanceHTTPSListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadBalanceHTTPSListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadBalanceHTTPSListener").
        Uri("/protocol/v3/listener/{loadBalanceId}/listeners/https").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/{loadBalanceId}/listeners/https").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadBalanceHTTPSListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateListener 修改监听器
func (c *Client) UpdateListener(request *model.UpdateListenerRequest) (*model.UpdateListenerResponse, error) {
    return c.UpdateListenerWithConfig(request, nil)
}

// UpdateListenerWithConfig 修改监听器
func (c *Client) UpdateListenerWithConfig(request *model.UpdateListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateListener").
        Uri("/acl/v3/listener").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// BindFip 绑定公网IP
func (c *Client) BindFip(request *model.BindFipRequest) (*model.BindFipResponse, error) {
    return c.BindFipWithConfig(request, nil)
}

// BindFipWithConfig 绑定公网IP
func (c *Client) BindFipWithConfig(request *model.BindFipRequest, runtimeConfig *config.RuntimeConfig) (*model.BindFipResponse, error) {
    params := param.NewParamsBuilder().
        Action("bindFip").
        Uri("/acl/v3/loadBalancer/bindfIp").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/bindfIp").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.BindFipResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// BatchDeleteControlGroup 批量删除访问控制组
func (c *Client) BatchDeleteControlGroup(request *model.BatchDeleteControlGroupRequest) (*model.BatchDeleteControlGroupResponse, error) {
    return c.BatchDeleteControlGroupWithConfig(request, nil)
}

// BatchDeleteControlGroupWithConfig 批量删除访问控制组
func (c *Client) BatchDeleteControlGroupWithConfig(request *model.BatchDeleteControlGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchDeleteControlGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchDeleteControlGroup").
        Uri("/acl/v3/accessControlGroup/batchDelete").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup/batchDelete").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.BatchDeleteControlGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetLoadBalanceDetailResp 根据负载均衡id,查询负载均衡器详情
func (c *Client) GetLoadBalanceDetailResp(request *model.GetLoadBalanceDetailRespRequest) (*model.GetLoadBalanceDetailRespResponse, error) {
    return c.GetLoadBalanceDetailRespWithConfig(request, nil)
}

// GetLoadBalanceDetailRespWithConfig 根据负载均衡id,查询负载均衡器详情
func (c *Client) GetLoadBalanceDetailRespWithConfig(request *model.GetLoadBalanceDetailRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetLoadBalanceDetailRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getLoadBalanceDetailResp").
        Uri("/acl/v3/loadBalancer/{loadBalanceId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/{loadBalanceId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetLoadBalanceDetailRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateControlGroup 更新访问控制组
func (c *Client) UpdateControlGroup(request *model.UpdateControlGroupRequest) (*model.UpdateControlGroupResponse, error) {
    return c.UpdateControlGroupWithConfig(request, nil)
}

// UpdateControlGroupWithConfig 更新访问控制组
func (c *Client) UpdateControlGroupWithConfig(request *model.UpdateControlGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateControlGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateControlGroup").
        Uri("/acl/v3/accessControlGroup").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateControlGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// AsyncUpdate 启用异步线程修改后端服务器组
func (c *Client) AsyncUpdate(request *model.AsyncUpdateRequest) (*model.AsyncUpdateResponse, error) {
    return c.AsyncUpdateWithConfig(request, nil)
}

// AsyncUpdateWithConfig 启用异步线程修改后端服务器组
func (c *Client) AsyncUpdateWithConfig(request *model.AsyncUpdateRequest, runtimeConfig *config.RuntimeConfig) (*model.AsyncUpdateResponse, error) {
    params := param.NewParamsBuilder().
        Action("asyncUpdate").
        Uri("/acl/v3/pool").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.AsyncUpdateResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateControlGroup 创建访问控制组
func (c *Client) CreateControlGroup(request *model.CreateControlGroupRequest) (*model.CreateControlGroupResponse, error) {
    return c.CreateControlGroupWithConfig(request, nil)
}

// CreateControlGroupWithConfig 创建访问控制组
func (c *Client) CreateControlGroupWithConfig(request *model.CreateControlGroupRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateControlGroupResponse, error) {
    params := param.NewParamsBuilder().
        Action("createControlGroup").
        Uri("/acl/v3/accessControlGroup").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroup").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateControlGroupResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// IDeleteListenerDeletables 校验监听器能否删除
func (c *Client) IDeleteListenerDeletables(request *model.IDeleteListenerDeletablesRequest) (*model.IDeleteListenerDeletablesResponse, error) {
    return c.IDeleteListenerDeletablesWithConfig(request, nil)
}

// IDeleteListenerDeletablesWithConfig 校验监听器能否删除
func (c *Client) IDeleteListenerDeletablesWithConfig(request *model.IDeleteListenerDeletablesRequest, runtimeConfig *config.RuntimeConfig) (*model.IDeleteListenerDeletablesResponse, error) {
    params := param.NewParamsBuilder().
        Action("iDeleteListenerDeletables").
        Uri("/acl/v3/listener/deletable").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/deletable").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.IDeleteListenerDeletablesResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteControlGroupRule 删除访问控制规则
func (c *Client) DeleteControlGroupRule(request *model.DeleteControlGroupRuleRequest) (*model.DeleteControlGroupRuleResponse, error) {
    return c.DeleteControlGroupRuleWithConfig(request, nil)
}

// DeleteControlGroupRuleWithConfig 删除访问控制规则
func (c *Client) DeleteControlGroupRuleWithConfig(request *model.DeleteControlGroupRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteControlGroupRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteControlGroupRule").
        Uri("/acl/v3/accessControlGroupRule/{accessControlGroupId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroupRule/{accessControlGroupId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteControlGroupRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// BatchDeleteLoadBalancePoolMember 批量删除负载均衡member
func (c *Client) BatchDeleteLoadBalancePoolMember(request *model.BatchDeleteLoadBalancePoolMemberRequest) (*model.BatchDeleteLoadBalancePoolMemberResponse, error) {
    return c.BatchDeleteLoadBalancePoolMemberWithConfig(request, nil)
}

// BatchDeleteLoadBalancePoolMemberWithConfig 批量删除负载均衡member
func (c *Client) BatchDeleteLoadBalancePoolMemberWithConfig(request *model.BatchDeleteLoadBalancePoolMemberRequest, runtimeConfig *config.RuntimeConfig) (*model.BatchDeleteLoadBalancePoolMemberResponse, error) {
    params := param.NewParamsBuilder().
        Action("batchDeleteLoadBalancePoolMember").
        Uri("/acl/v3/member/{poolId}/member/batchDelete").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{poolId}/member/batchDelete").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.BatchDeleteLoadBalancePoolMemberResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CertificationUnbindTag 解绑标签
func (c *Client) CertificationUnbindTag(request *model.CertificationUnbindTagRequest) (*model.CertificationUnbindTagResponse, error) {
    return c.CertificationUnbindTagWithConfig(request, nil)
}

// CertificationUnbindTagWithConfig 解绑标签
func (c *Client) CertificationUnbindTagWithConfig(request *model.CertificationUnbindTagRequest, runtimeConfig *config.RuntimeConfig) (*model.CertificationUnbindTagResponse, error) {
    params := param.NewParamsBuilder().
        Action("certificationUnbindTag").
        Uri("/customer/v3/certification/unbindTag/{containerUuid}").
        GatewayUri("/api/openapi-vlb/lb-console-openstack-lb-barbican/customer/v3/certification/unbindTag/{containerUuid}").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CertificationUnbindTagResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetHealthStatusResp 查询后端服务器的健康检查状态
func (c *Client) GetHealthStatusResp(request *model.GetHealthStatusRespRequest) (*model.GetHealthStatusRespResponse, error) {
    return c.GetHealthStatusRespWithConfig(request, nil)
}

// GetHealthStatusRespWithConfig 查询后端服务器的健康检查状态
func (c *Client) GetHealthStatusRespWithConfig(request *model.GetHealthStatusRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetHealthStatusRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getHealthStatusResp").
        Uri("/acl/v3/member/healthStatus/{poolId}/{memberId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/healthStatus/{poolId}/{memberId}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetHealthStatusRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateLoadBalancePoolMembers OpenAPI批量添加多个负载均衡后端服务器
func (c *Client) CreateLoadBalancePoolMembers(request *model.CreateLoadBalancePoolMembersRequest) (*model.CreateLoadBalancePoolMembersResponse, error) {
    return c.CreateLoadBalancePoolMembersWithConfig(request, nil)
}

// CreateLoadBalancePoolMembersWithConfig OpenAPI批量添加多个负载均衡后端服务器
func (c *Client) CreateLoadBalancePoolMembersWithConfig(request *model.CreateLoadBalancePoolMembersRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateLoadBalancePoolMembersResponse, error) {
    params := param.NewParamsBuilder().
        Action("createLoadBalancePoolMembers").
        Uri("/acl/v3/member/{poolId}/openApiMembers").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{poolId}/openApiMembers").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateLoadBalancePoolMembersResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// GetLoadbalanceCertificationDetailResp 获取负载均衡证书详情
func (c *Client) GetLoadbalanceCertificationDetailResp(request *model.GetLoadbalanceCertificationDetailRespRequest) (*model.GetLoadbalanceCertificationDetailRespResponse, error) {
    return c.GetLoadbalanceCertificationDetailRespWithConfig(request, nil)
}

// GetLoadbalanceCertificationDetailRespWithConfig 获取负载均衡证书详情
func (c *Client) GetLoadbalanceCertificationDetailRespWithConfig(request *model.GetLoadbalanceCertificationDetailRespRequest, runtimeConfig *config.RuntimeConfig) (*model.GetLoadbalanceCertificationDetailRespResponse, error) {
    params := param.NewParamsBuilder().
        Action("getLoadbalanceCertificationDetailResp").
        Uri("/acl/v3/certification/{containerUuid}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/certification/{containerUuid}").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.GetLoadbalanceCertificationDetailRespResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateLoadBalanceUDPListener 创建UDP监听器
func (c *Client) CreateLoadBalanceUDPListener(request *model.CreateLoadBalanceUDPListenerRequest) (*model.CreateLoadBalanceUDPListenerResponse, error) {
    return c.CreateLoadBalanceUDPListenerWithConfig(request, nil)
}

// CreateLoadBalanceUDPListenerWithConfig 创建UDP监听器
func (c *Client) CreateLoadBalanceUDPListenerWithConfig(request *model.CreateLoadBalanceUDPListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateLoadBalanceUDPListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("createLoadBalanceUDPListener").
        Uri("/protocol/v3/listener/udp").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/udp").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateLoadBalanceUDPListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadBalanceUDPListenerPool 修改UDP监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceUDPListenerPool(request *model.UpdateLoadBalanceUDPListenerPoolRequest) (*model.UpdateLoadBalanceUDPListenerPoolResponse, error) {
    return c.UpdateLoadBalanceUDPListenerPoolWithConfig(request, nil)
}

// UpdateLoadBalanceUDPListenerPoolWithConfig 修改UDP监听器绑定的后端服务器组
func (c *Client) UpdateLoadBalanceUDPListenerPoolWithConfig(request *model.UpdateLoadBalanceUDPListenerPoolRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadBalanceUDPListenerPoolResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadBalanceUDPListenerPool").
        Uri("/protocol/v3/listener/udp/updatePool").
        GatewayUri("/api/openapi-vlb/lb-console/protocol/v3/listener/udp/updatePool").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadBalanceUDPListenerPoolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListControlGroupRule 查询访问控制组规则
func (c *Client) ListControlGroupRule(request *model.ListControlGroupRuleRequest) (*model.ListControlGroupRuleResponse, error) {
    return c.ListControlGroupRuleWithConfig(request, nil)
}

// ListControlGroupRuleWithConfig 查询访问控制组规则
func (c *Client) ListControlGroupRuleWithConfig(request *model.ListControlGroupRuleRequest, runtimeConfig *config.RuntimeConfig) (*model.ListControlGroupRuleResponse, error) {
    params := param.NewParamsBuilder().
        Action("listControlGroupRule").
        Uri("/acl/v3/accessControlGroupRule").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/accessControlGroupRule").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListControlGroupRuleResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteLoadbalanceCertification 删除负载均衡证书
func (c *Client) DeleteLoadbalanceCertification(request *model.DeleteLoadbalanceCertificationRequest) (*model.DeleteLoadbalanceCertificationResponse, error) {
    return c.DeleteLoadbalanceCertificationWithConfig(request, nil)
}

// DeleteLoadbalanceCertificationWithConfig 删除负载均衡证书
func (c *Client) DeleteLoadbalanceCertificationWithConfig(request *model.DeleteLoadbalanceCertificationRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteLoadbalanceCertificationResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteLoadbalanceCertification").
        Uri("/acl/v3/certification/{containerUuid}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/certification/{containerUuid}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteLoadbalanceCertificationResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// IsPoolDeletables 校验pool能否删除
func (c *Client) IsPoolDeletables(request *model.IsPoolDeletablesRequest) (*model.IsPoolDeletablesResponse, error) {
    return c.IsPoolDeletablesWithConfig(request, nil)
}

// IsPoolDeletablesWithConfig 校验pool能否删除
func (c *Client) IsPoolDeletablesWithConfig(request *model.IsPoolDeletablesRequest, runtimeConfig *config.RuntimeConfig) (*model.IsPoolDeletablesResponse, error) {
    params := param.NewParamsBuilder().
        Action("isPoolDeletables").
        Uri("/acl/v3/pool/deletable").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/pool/deletable").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.IsPoolDeletablesResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteListener 删除负载均衡监听器
func (c *Client) DeleteListener(request *model.DeleteListenerRequest) (*model.DeleteListenerResponse, error) {
    return c.DeleteListenerWithConfig(request, nil)
}

// DeleteListenerWithConfig 删除负载均衡监听器
func (c *Client) DeleteListenerWithConfig(request *model.DeleteListenerRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteListenerResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteListener").
        Uri("/acl/v3/listener/{listenerId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/{listenerId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteListenerResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// ListLoadbalance 查询负载均衡器列表
func (c *Client) ListLoadbalance(request *model.ListLoadbalanceRequest) (*model.ListLoadbalanceResponse, error) {
    return c.ListLoadbalanceWithConfig(request, nil)
}

// ListLoadbalanceWithConfig 查询负载均衡器列表
func (c *Client) ListLoadbalanceWithConfig(request *model.ListLoadbalanceRequest, runtimeConfig *config.RuntimeConfig) (*model.ListLoadbalanceResponse, error) {
    params := param.NewParamsBuilder().
        Action("listLoadbalance").
        Uri("/acl/v3/loadBalancer/loadBalancers").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/loadBalancers").
        Protocol("http").
        ContentType("application/json").
        Method("GET").
        Request(request).
        Build()
    returnValue := &model.ListLoadbalanceResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteMemberWithHttpStatus 删除后端服务器(修复HTTP状态码)
func (c *Client) DeleteMemberWithHttpStatus(request *model.DeleteMemberWithHttpStatusRequest) (*model.DeleteMemberWithHttpStatusResponse, error) {
    return c.DeleteMemberWithHttpStatusWithConfig(request, nil)
}

// DeleteMemberWithHttpStatusWithConfig 删除后端服务器(修复HTTP状态码)
func (c *Client) DeleteMemberWithHttpStatusWithConfig(request *model.DeleteMemberWithHttpStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteMemberWithHttpStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteMemberWithHttpStatus").
        Uri("/acl/v3/member/{poolId}/delete/{memberId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{poolId}/delete/{memberId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteMemberWithHttpStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// DeleteListenerWithHttpStatus 删除负载均衡监听器(优化HTTP状态码)
func (c *Client) DeleteListenerWithHttpStatus(request *model.DeleteListenerWithHttpStatusRequest) (*model.DeleteListenerWithHttpStatusResponse, error) {
    return c.DeleteListenerWithHttpStatusWithConfig(request, nil)
}

// DeleteListenerWithHttpStatusWithConfig 删除负载均衡监听器(优化HTTP状态码)
func (c *Client) DeleteListenerWithHttpStatusWithConfig(request *model.DeleteListenerWithHttpStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.DeleteListenerWithHttpStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("deleteListenerWithHttpStatus").
        Uri("/acl/v3/listener/delete/{listenerId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/listener/delete/{listenerId}").
        Protocol("http").
        ContentType("application/json").
        Method("DELETE").
        Request(request).
        Build()
    returnValue := &model.DeleteListenerWithHttpStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateTLSCustomizeProtocol 创建TLS安全策略
func (c *Client) CreateTLSCustomizeProtocol(request *model.CreateTLSCustomizeProtocolRequest) (*model.CreateTLSCustomizeProtocolResponse, error) {
    return c.CreateTLSCustomizeProtocolWithConfig(request, nil)
}

// CreateTLSCustomizeProtocolWithConfig 创建TLS安全策略
func (c *Client) CreateTLSCustomizeProtocolWithConfig(request *model.CreateTLSCustomizeProtocolRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateTLSCustomizeProtocolResponse, error) {
    params := param.NewParamsBuilder().
        Action("createTLSCustomizeProtocol").
        Uri("/acl/v3/tls").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/tls").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateTLSCustomizeProtocolResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateMembersWithHttpStatus 批量添加后端服务器(修复HTTP状态码)
func (c *Client) CreateMembersWithHttpStatus(request *model.CreateMembersWithHttpStatusRequest) (*model.CreateMembersWithHttpStatusResponse, error) {
    return c.CreateMembersWithHttpStatusWithConfig(request, nil)
}

// CreateMembersWithHttpStatusWithConfig 批量添加后端服务器(修复HTTP状态码)
func (c *Client) CreateMembersWithHttpStatusWithConfig(request *model.CreateMembersWithHttpStatusRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateMembersWithHttpStatusResponse, error) {
    params := param.NewParamsBuilder().
        Action("createMembersWithHttpStatus").
        Uri("/acl/v3/member/{poolId}/create/members").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/member/{poolId}/create/members").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateMembersWithHttpStatusResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// LoadbalancerUnbindTag 解绑标签
func (c *Client) LoadbalancerUnbindTag(request *model.LoadbalancerUnbindTagRequest) (*model.LoadbalancerUnbindTagResponse, error) {
    return c.LoadbalancerUnbindTagWithConfig(request, nil)
}

// LoadbalancerUnbindTagWithConfig 解绑标签
func (c *Client) LoadbalancerUnbindTagWithConfig(request *model.LoadbalancerUnbindTagRequest, runtimeConfig *config.RuntimeConfig) (*model.LoadbalancerUnbindTagResponse, error) {
    params := param.NewParamsBuilder().
        Action("loadbalancerUnbindTag").
        Uri("/acl/v3/loadBalancer/unbindTag/{loadbalanceId}").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/loadBalancer/unbindTag/{loadbalanceId}").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.LoadbalancerUnbindTagResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// UpdateLoadbalanceCertification 更新负载均衡证书
func (c *Client) UpdateLoadbalanceCertification(request *model.UpdateLoadbalanceCertificationRequest) (*model.UpdateLoadbalanceCertificationResponse, error) {
    return c.UpdateLoadbalanceCertificationWithConfig(request, nil)
}

// UpdateLoadbalanceCertificationWithConfig 更新负载均衡证书
func (c *Client) UpdateLoadbalanceCertificationWithConfig(request *model.UpdateLoadbalanceCertificationRequest, runtimeConfig *config.RuntimeConfig) (*model.UpdateLoadbalanceCertificationResponse, error) {
    params := param.NewParamsBuilder().
        Action("updateLoadbalanceCertification").
        Uri("/acl/v3/certification").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/certification").
        Protocol("http").
        ContentType("application/json").
        Method("PUT").
        Request(request).
        Build()
    returnValue := &model.UpdateLoadbalanceCertificationResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

// CreateLoadbalanceCertification 创建负载均衡证书
func (c *Client) CreateLoadbalanceCertification(request *model.CreateLoadbalanceCertificationRequest) (*model.CreateLoadbalanceCertificationResponse, error) {
    return c.CreateLoadbalanceCertificationWithConfig(request, nil)
}

// CreateLoadbalanceCertificationWithConfig 创建负载均衡证书
func (c *Client) CreateLoadbalanceCertificationWithConfig(request *model.CreateLoadbalanceCertificationRequest, runtimeConfig *config.RuntimeConfig) (*model.CreateLoadbalanceCertificationResponse, error) {
    params := param.NewParamsBuilder().
        Action("createLoadbalanceCertification").
        Uri("/acl/v3/certification").
        GatewayUri("/api/openapi-vlb/lb-console/acl/v3/certification").
        Protocol("http").
        ContentType("application/json").
        Method("POST").
        Request(request).
        Build()
    returnValue := &model.CreateLoadbalanceCertificationResponse{}
    if _, err := c.apiClient.Excute(params, runtimeConfig, returnValue); err != nil {
        return nil, err
    } else {
        return returnValue, nil
    }
}

