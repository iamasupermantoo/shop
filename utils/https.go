package utils

import (
	"bytes"
	"github.com/goccy/go-json"
	"golang.org/x/net/proxy"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client    *http.Client      //	client对象
	transport *http.Transport   //	代理对象
	headers   map[string]string // 请求头信息
	error     error             // 错误信息
}

// ClientProxyInfo 代理信息
type ClientProxyInfo struct {
	Host string //	主机
	User string // 用户
	Pass string // 密码
}

// NewClient 创建连接
func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		headers: map[string]string{},
	}
}

// SetHeaders 设置头信息
func (_Client *Client) SetHeaders(headers map[string]string) *Client {
	for key, val := range headers {
		_Client.headers[key] = val
	}
	return _Client
}

// Request 请求信息
func (_Client *Client) Request(method string, url string, params interface{}) ([]byte, error) {
	if _Client.error != nil {
		return nil, _Client.error
	}

	var bufferParams io.Reader
	if params != nil {
		dataBytes, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		bufferParams = bytes.NewBuffer(dataBytes)
	}

	// 创建请求链接
	req, err := http.NewRequest(method, url, bufferParams)
	if err != nil {
		return nil, err
	}

	// 设置请求头信息
	for key, val := range _Client.headers {
		req.Header.Set(key, val)
	}

	// 请求信息
	if _Client.transport != nil {
		_Client.client.Transport = _Client.transport
	}
	resp, err := _Client.client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 关闭连接
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	// 如果有设置代理, 那么关闭闲置连接
	if _Client.transport != nil {
		_Client.transport.CloseIdleConnections()
	}
	return bodyBytes, nil
}

// SetTimeout 设置超时时间
func (_Client *Client) SetTimeout(timeout time.Duration) *Client {
	_Client.client.Timeout = timeout
	return _Client
}

// SetTransport 设置代理
func (_Client *Client) SetTransport(transport *http.Transport) *Client {
	_Client.transport = transport
	return _Client
}

// SetSocket5 设置Socket5
func (_Client *Client) SetSocket5(proxyInfo *ClientProxyInfo) *Client {
	if proxyInfo != nil {
		dialer, err := proxy.SOCKS5("tcp", proxyInfo.Host, &proxy.Auth{User: proxyInfo.User, Password: proxyInfo.Pass}, proxy.Direct)
		if err != nil {
			_Client.error = err
		}

		// 转成上下文模式, 可以设置超时等模式
		if contextDialer, ok := dialer.(proxy.ContextDialer); ok {
			_Client.transport = &http.Transport{
				DialContext: contextDialer.DialContext,
			}
		}
	}
	return _Client
}

// GetHttpHost 获取http服务器地址
func GetHttpHost(rawUrl string) string {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return ""
	}
	return parsedURL.Scheme + "://" + parsedURL.Host
}
