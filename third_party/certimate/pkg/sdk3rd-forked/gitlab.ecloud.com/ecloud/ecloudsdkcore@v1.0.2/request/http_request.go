package request

import (
	"fmt"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/consts"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/param"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
	"reflect"
	"strings"
)

type HttpRequest struct {
	Method       string            `json:"method,omitempty"`
	Protocol     string            `json:"protocol,omitempty"`
	Url          string            `json:"url,omitempty"`
	Path         string            `json:"path,omitempty"`
	Action       string            `json:"action,omitempty"`
	Product      string            `json:"product,omitempty"`
	Version      string            `json:"version,omitempty"`
	SdkVersion   string            `json:"sdkVersion,omitempty"`
	Source       string            `json:"source,omitempty"`
	QueryString  string            `json:"queryString,omitempty"`
	ContentType  string            `json:"contentType,omitempty"`
	Endpoint     string            `json:"endpoint,omitempty"`
	Body         interface{}       `json:"body,omitempty"`
	PathParams   map[string]string `json:"pathParams,omitempty"`
	QueryParams  map[string]string `json:"queryParams,omitempty"`
	HeaderParams map[string]string `json:"headerParams,omitempty"`
}

type HttpRequestPosition string

const (
	BODY   HttpRequestPosition = "Body"
	QUERY  HttpRequestPosition = "Query"
	PATH   HttpRequestPosition = "Path"
	HEADER HttpRequestPosition = "Header"
)

func NewHttpRequest() *HttpRequest {
	return &HttpRequest{
		Method:       "POST",
		QueryParams:  make(map[string]string),
		HeaderParams: make(map[string]string),
		PathParams:   make(map[string]string),
		ContentType:  "application/json; charset=utf-8",
	}
}

func DefaultHttpRequest() *HttpRequest {
	return NewHttpRequest()
}

func (httpReq *HttpRequest) Reset() {
	httpReq.Action = ""
	httpReq.QueryParams = make(map[string]string)
	httpReq.HeaderParams = make(map[string]string)
	httpReq.PathParams = make(map[string]string)
	httpReq.Body = nil
	httpReq.Method = "POST"
}

func (httpReq *HttpRequest) BuildHeaderParams(headerParams map[string]string) {
	if httpReq.HeaderParams == nil {
		httpReq.HeaderParams = make(map[string]string)
	}
	for name, value := range headerParams {
		httpReq.HeaderParams[name] = value
	}
}

func (httpReq *HttpRequest) AddDefaultHeaders() {
	if httpReq.HeaderParams == nil {
		httpReq.HeaderParams = make(map[string]string)
	}
	x := fmt.Sprintf("action:%s;product:%s;version:%s;sdkversion:%s;language:Golang",
		httpReq.Action, httpReq.Product, httpReq.Version, httpReq.SdkVersion)
	if len(httpReq.Source) > 0 {
		x += fmt.Sprintf(";source:%s", httpReq.Source)
	}
	httpReq.HeaderParams["x-openapi-sdk"] = x
	httpReq.HeaderParams["User-Agent"] = "OpenAPI/2.0/Golang"
}

func (httpReq *HttpRequest) BuildQueryParams(queryParams map[string]string) {
	if httpReq.QueryParams == nil {
		httpReq.QueryParams = make(map[string]string)
	}
	for name, value := range queryParams {
		httpReq.QueryParams[name] = value
	}
}

func (httpReq *HttpRequest) BuildPathParams(pathParams map[string]string) {
	if httpReq.PathParams == nil {
		httpReq.PathParams = make(map[string]string)
	}
	for name, value := range pathParams {
		httpReq.PathParams[name] = value
	}
}

func (httpReq *HttpRequest) BuildApiParams(params *param.Params, c *config.Config, rc *config.RuntimeConfig) {
	if params == nil {
		return
	}

	if len(params.ContentType) > 0 {
		httpReq.ContentType = params.ContentType
	}
	if len(params.Method) > 0 {
		httpReq.Method = params.Method
	}
	if len(params.Protocol) > 0 {
		httpReq.Protocol = params.Protocol
	}
	ignoreGateway := utils.DefaultBool(rc.IgnoreGateway, c.IgnoreGateway)
	if utils.BoolValue(ignoreGateway) {
		httpReq.Path = params.Uri
	} else {
		httpReq.Path = params.GatewayUri
	}
	if len(params.Action) > 0 {
		httpReq.Action = params.Action
	}
	if httpReq.PathParams != nil && len(httpReq.PathParams) > 0 {
		httpReq.BuildPathParamsString()
	}
	if c.GlobalHeaderParams != nil && len(c.GlobalHeaderParams) > 0 {
		httpReq.BuildHeaderParams(c.GlobalHeaderParams)
	}
	poolId := utils.StringValue(c.PoolId)
	if len(poolId) > 0 {
		httpReq.AddHeaders(map[string]string{"Pool-Id": poolId})
	}
	httpReq.AddDefaultHeaders()
	if c.GlobalHeaderParams != nil && len(c.GlobalHeaderParams) > 0 {
		httpReq.BuildHeaderParams(c.GlobalHeaderParams)
	}
	if rc.RuntimeHeaderParams != nil && len(rc.RuntimeHeaderParams) > 0 {
		httpReq.BuildHeaderParams(rc.RuntimeHeaderParams)
	}
	authType := c.AuthType
	if utils.IsSet(rc.AuthType) {
		authType = rc.AuthType
	}
	if *authType != consts.AK {
		httpReq.BuildQueryParamsString()
	}
	centralTransportEnabled := utils.DefaultBool(rc.CentralTransportEnabled, c.CentralTransportEnabled)
	if !utils.BoolValue(centralTransportEnabled) && len(httpReq.Endpoint) > 0 {
		httpReq.Url = httpReq.Endpoint
	} else {
		httpReq.Url = utils.DefaultEndpoint
	}
	httpProxy := utils.DefaultString(rc.HttpProxy, c.HttpProxy)
	if httpProxy != nil {
		httpReq.Url = utils.StringValue(httpProxy)
	}
	httpsProxy := utils.DefaultString(rc.HttpsProxy, c.HttpsProxy)
	if httpsProxy != nil {
		httpReq.Url = utils.StringValue(httpsProxy)
	}
}

func (httpReq *HttpRequest) BuildFinalUrl() {
	url := httpReq.Url
	if len(url) == 0 {
		return
	}
	protocol := httpReq.Protocol
	if len(protocol) > 0 &&
		!strings.HasPrefix(url, "http") &&
		!strings.HasPrefix(url, "https") {
		url = protocol + "://" + url
	}
	path := httpReq.Path
	if len(path) > 0 {
		url = url + path
	}
	queryString := httpReq.QueryString
	if len(queryString) > 0 {
		url = url + queryString
	}
	httpReq.Url = url
}

func (httpReq *HttpRequest) BuildPathParamsString() {
	for name, value := range httpReq.PathParams {
		httpReq.Path = strings.ReplaceAll(httpReq.Path, fmt.Sprintf("{%s}", name), value)
	}
	if !strings.HasPrefix(httpReq.Path, "/") {
		httpReq.Path = "/" + httpReq.Path
	}
	if strings.HasSuffix(httpReq.Path, "/") {
		httpReq.Path = httpReq.Path[0 : len(httpReq.Path)-1]
	}
}

func (httpReq *HttpRequest) BuildQueryParamsString() {
	build := strings.Builder{}
	paramCount := len(httpReq.QueryParams)
	for name, value := range httpReq.QueryParams {
		build.WriteString(utils.PercentEncode(name))
		build.WriteString("=")
		build.WriteString(utils.PercentEncode(value))
		paramCount--
		if paramCount > 0 {
			build.WriteString("&")
		}
	}
	httpReq.QueryString = build.String()
}

func (httpReq *HttpRequest) ConvertRequest(request interface{}) {
	if request == nil {
		return
	}
	reqType := reflect.TypeOf(request)
	if reqType.Kind() == reflect.Ptr {
		reqType = reqType.Elem()
	}
	reqValue := reflect.ValueOf(request)
	if reqValue.Kind() == reflect.Ptr {
		reqValue = reqValue.Elem()
	}
	var flag = false
	for i := 0; i < reqType.NumField(); i++ {
		fieldType := reqType.Field(i)
		value := reqValue.FieldByName(fieldType.Name)
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue
			}
			value = value.Elem()
		}
		propertyType := fieldType.Type
		if propertyType.Kind() == reflect.Ptr {
			propertyType = propertyType.Elem()
		}

		_, flag = propertyType.FieldByName(string(BODY))
		if flag {
			httpReq.Body = value.Interface()
			continue
		}
		_, flag = propertyType.FieldByName(string(HEADER))
		if flag {
			httpReq.BuildHeaderParams(utils.StructToMap(value.Interface(), ignorePositionField))
			continue
		}
		_, flag = propertyType.FieldByName(string(QUERY))
		if flag {
			httpReq.BuildQueryParams(utils.StructToMap(value.Interface(), ignorePositionField))
			continue
		}
		_, flag = propertyType.FieldByName(string(PATH))
		if flag {
			httpReq.BuildPathParams(utils.StructToMap(value.Interface(), ignorePositionField))
			continue
		}
	}
}

func (httpReq *HttpRequest) AddHeaders(headers map[string]string) {
	if httpReq.HeaderParams == nil {
		httpReq.HeaderParams = map[string]string{}
	}
	for name, value := range headers {
		httpReq.HeaderParams[name] = value
	}
}

var _typeOfHeader = reflect.TypeOf(position.Header{})
var _typeOfBody = reflect.TypeOf(position.Body{})
var _typeOfQuery = reflect.TypeOf(position.Query{})
var _typeOfPath = reflect.TypeOf(position.Path{})

func ignorePositionField(obj interface{}, field reflect.StructField, value reflect.Value) bool {
	if utils.ValueIsEmpty(value) {
		return true
	}
	typeOfValue := value.Type()
	return typeOfValue == _typeOfBody || typeOfValue == _typeOfQuery ||
		typeOfValue == _typeOfHeader || typeOfValue == _typeOfPath
}
