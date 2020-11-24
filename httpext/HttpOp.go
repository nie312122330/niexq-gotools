package httpext

import (
	"bytes"
	"github.com/nie312122330/niexq-gotools/jsonext"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GetText请求
// url：         请求地址
func GetText(url string,timeOut time.Duration) (string,error) {
	b,err:=Get(url,timeOut)
	if err != nil {
		return "",err
	}
	return string(b),nil
}
// 发送GET请求
// url：         请求地址
func Get(url string,timeOut time.Duration) ([]byte,error) {
	client := &http.Client{Timeout: timeOut}
	resp, err := client.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return body,nil
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
func PostJson(url string, data interface{}, timeOut time.Duration) (string,error) {
	return Post(url,data,"application/json",timeOut)
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string,timeOut time.Duration) (string,error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: timeOut}
	jsonStr ,_:=jsonext.ToJsonBytes(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(*jsonStr))
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result),nil
}