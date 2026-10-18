package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/errs"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/request"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/response"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type HttpClient struct {
	goClient       *http.Client
	ignoreSSL      bool
	certFile       string
	clientCertFile string
	clientKeyFile  string
}

type RetryFunc func(request *request.HttpRequest,
	returnType *interface{}) (*response.HttpResponse, error)

func NewHttpClient() *HttpClient {
	return &HttpClient{
		ignoreSSL: true,
		goClient:  defaultGoClient(),
	}
}

func defaultGoClient() *http.Client {
	return &http.Client{
		Transport: defaultTransport(),
	}
}

func defaultTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}

// doRequest do the request.
func (hc *HttpClient) doRequest(request *http.Request) (*http.Response, error) {
	return hc.goClient.Do(request)
}

func (hc *HttpClient) Execute(request *request.HttpRequest,
	returnType interface{}) (*response.HttpResponse, error) {

	req := hc.buildRequest(request)
	resp, err := hc.doRequest(req)
	if err != nil {
		return nil, errs.NewServerRequestError(err.Error(), err)
	}
	if resp == nil {
		return nil, errs.NewServerResponseError("response is nil", nil, -1, nil, "")
	}
	if err = handleResponse(resp, returnType); err != nil {
		return nil, err
	}
	return &response.HttpResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Data:       resp,
	}, nil
}

func handleResponse(resp *http.Response, returnType interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
		return errs.NewServerResponseError(fmt.Sprintf("response status is: %s,status code is:%d",
			resp.Status, resp.StatusCode), err, resp.StatusCode, resp.Header, "can not read response body")
	}
	// successful
	if isSuccessful(resp.StatusCode) {
		if returnType == nil || resp.StatusCode == 204 {
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
			return nil
		}
		// If we succeed, return the data, otherwise pass on to deserialize error.
		err = deserialize(returnType, body, resp.Header.Get("Content-Type"))
		if err != nil {
			return errs.NewServerResponseError(fmt.Sprintf("can't deserialize respnse body with content-type: %s",
				resp.Header.Get("Content-Type")), err, resp.StatusCode, resp.Header, string(body))
		}
	} else {
		if body == nil {
			return errs.NewServerResponseError(fmt.Sprintf("response status is: %s,status code is:%d",
				resp.Status, resp.StatusCode), nil, resp.StatusCode, resp.Header, "response body is nil")
		}
		respBody := string(body)
		return errs.NewServerResponseError(fmt.Sprintf("response status: %s,http status code:%d,response body:%s",
			resp.Status, resp.StatusCode, respBody), nil, resp.StatusCode, resp.Header, respBody)
	}
	return nil
}

func isSuccessful(code int) bool {
	return code >= 200 && code < 300
}

func (hc *HttpClient) buildRequest(httpRequest *request.HttpRequest) (request *http.Request) {
	req, _ := prepareRequest(httpRequest)
	return req
}

func (hc *HttpClient) ApplySSLSettings(ignoreSSL bool, certFile string, clientCertFile string, clientKeyFile string) error {
	if ignoreSSL && hc.ignoreSSL {
		return nil
	}
	if ignoreSSL && !hc.ignoreSSL {
		hc.ignoreSSL = ignoreSSL
		transport := hc.goClient.Transport.(*http.Transport)
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: ignoreSSL,
		}
		hc.goClient.Transport = transport
		return nil
	}
	if !ignoreSSL && !hc.ignoreSSL {
		if certFile == hc.certFile && clientCertFile == hc.clientCertFile &&
			clientKeyFile == hc.clientKeyFile {
			return nil
		}
	}

	var certP *x509.CertPool = nil
	if len(certFile) > 0 {
		certP = &x509.CertPool{}
		pemCert, err := ioutil.ReadFile(certFile)
		if err != nil {
			return errs.NewSslHandShakeError(fmt.Sprintf("can't read certfile: %s", certFile), err)
		}
		certP.AppendCertsFromPEM(pemCert)
	}
	transport := hc.goClient.Transport.(*http.Transport)
	var clientCert *tls.Certificate = nil
	if len(clientKeyFile) > 0 && len(clientCertFile) > 0 {
		cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
		if err != nil {
			return errs.NewSslHandShakeError(fmt.Sprintf("can't load keypair of client cert file: %s, client key file: %s",
				clientCertFile, clientKeyFile), err)
		}
		clientCert = &cert
	}

	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: ignoreSSL,
	}
	if certP != nil {
		transport.TLSClientConfig.RootCAs = certP
	}
	if clientCert != nil {
		transport.TLSClientConfig.Certificates = []tls.Certificate{*clientCert}
	}
	hc.goClient.Transport = transport
	hc.certFile = certFile
	hc.clientCertFile = clientCertFile
	hc.clientKeyFile = clientKeyFile
	return nil
}

func (hc *HttpClient) SetReadTimeout(timeout int32) {
	if hc.goClient == nil {
		hc.goClient = defaultGoClient()
	}
	hc.goClient.Timeout = time.Duration(timeout) * time.Second
}

func (hc *HttpClient) SetConnectTimeout(timeout int32) {
	if hc.goClient == nil {
		hc.goClient = defaultGoClient()
	}
	transport := hc.goClient.Transport.(*http.Transport)
	transport.DialContext = hc.setDialContext(timeout)
	hc.goClient.Transport = transport
}

func (hc *HttpClient) setDialContext(timeout int32) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		return (&net.Dialer{
			Timeout:   time.Duration(timeout) * time.Second,
			DualStack: true,
		}).DialContext(ctx, network, address)
	}
}

func (hc *HttpClient) SetIgnoreSSL(ignoreSSL bool) {
	if ignoreSSL != hc.ignoreSSL {
		transport := defaultTransport()
		transport.TLSClientConfig.InsecureSkipVerify = ignoreSSL
		hc.goClient.Transport = transport
		hc.ignoreSSL = ignoreSSL
	}
}

func (hc *HttpClient) SetCertFile(certFile string) error {
	if certFile != hc.certFile {
		return hc.ApplySSLSettings(hc.ignoreSSL, certFile, hc.clientCertFile, hc.clientKeyFile)
	}
	return nil
}

func (hc *HttpClient) SetClientKeyPairFile(clientCertFile string, clientKeyFile string) error {
	if clientCertFile != hc.clientCertFile || clientKeyFile != hc.clientKeyFile {
		return hc.ApplySSLSettings(hc.ignoreSSL, hc.certFile,
			hc.clientCertFile, hc.clientKeyFile)
	}
	return nil
}

// prepareRequest build the request
func prepareRequest(httpRequest *request.HttpRequest) (req *http.Request, err error) {
	var body *bytes.Buffer
	rawBody := httpRequest.Body
	contentType := httpRequest.ContentType
	// Detect rawBody
	if rawBody != nil {
		contentType = utils.DetectContentType(rawBody)
		body, err = utils.SetBody(rawBody, contentType)
		if err != nil {
			return nil, errs.NewServerRequestError(fmt.Sprintf("set request body error: %s", err.Error()), err)
		}
	}

	// Setup url and query parameters, url contains query strings
	rawUrl := httpRequest.Url
	if len(rawUrl) == 0 {
		return nil, errs.NewServerRequestError("request url is empty", nil)
	}
	realUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, errs.NewServerRequestError(fmt.Sprintf("can't parse request url: %s, "+
			"error is: %s", rawUrl, err.Error()), err)
	}

	// Generate a new http.Request
	method := httpRequest.Method
	if body != nil {
		req, err = http.NewRequest(method, realUrl.String(), body)
	} else {
		req, err = http.NewRequest(method, realUrl.String(), nil)
	}
	if err != nil {
		return nil, errs.NewServerRequestError(fmt.Sprintf("can't create http request, "+
			"method=%s, url=%s, error: %s", method, realUrl.String(), err.Error()), err)
	}

	// Add request headers
	headers := httpRequest.HeaderParams
	for name, value := range headers {
		req.Header.Add(name, value)
	}
	if len(contentType) == 0 {
		contentType = "application/json; charset=utf-8"
	}
	req.Header.Add("Content-Type", contentType)
	return req, nil
}

func deserialize(respType interface{}, respBody []byte, contentType string) (err error) {
	if respBody == nil {
		//return errs.NewGenericResponseError("response body is nil", nil, respBody)
		return errs.NewServerResponseError("response body is nil", nil, -1, nil, "")
	}
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(respBody, respType); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "json") {
		if err = json.Unmarshal(respBody, respType); err != nil {
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}
