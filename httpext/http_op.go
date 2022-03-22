package httpext

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nie312122330/niexq-gotools/jsonext"
)

// GetText 发送GetText请求
// url：         请求地址
func GetText(url string, timeOut time.Duration) (string, error) {
	b, err := Get(url, timeOut)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Get 发送GET请求
// url：         请求地址
func Get(url string, timeOut time.Duration) ([]byte, error) {
	client := &http.Client{Timeout: timeOut, Transport: &http.Transport{DisableKeepAlives: true}}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("返回非%d错误", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// PostJSON 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
func PostJSON(url string, data interface{}, timeOut time.Duration) (string, error) {
	return Post(url, data, "application/json", timeOut)
}

// Post 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string, timeOut time.Duration) (string, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: timeOut, Transport: &http.Transport{DisableKeepAlives: true}}
	jsonStr, _ := jsonext.ToJSONBytes(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(*jsonStr))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("返回非%d错误", resp.StatusCode)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}
