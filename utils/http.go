package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPResponse 表示通用的 HTTP 响应结构
type HTTPResponse struct {
	StatusCode int
	Body       []byte
	Err        error
}

// PostJSON 发起一个 JSON 格式的 POST 请求
func PostJSON(url string, data interface{}, headers map[string]string) (*HTTPResponse, error) {
	// 将请求数据编码为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode JSON: %v", err)
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置默认的 Content-Type
	req.Header.Set("Content-Type", "application/json")

	// 添加自定义请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 创建 HTTP 客户端并设置超时时间
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// 返回封装的响应
	return &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
		Err:        nil,
	}, nil
}

// Get 发起一个 GET 请求
func Get(url string, headers map[string]string) (*HTTPResponse, error) {
	// 创建 GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 添加自定义请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 创建 HTTP 客户端并设置超时时间
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// 返回封装的响应
	return &HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
		Err:        nil,
	}, nil
}
