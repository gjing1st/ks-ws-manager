// Path: third_party/k8s
// FileName: req.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/4/26$ 19:06$

package client

import (
	"encoding/json"
	ks_error "github.com/gjing1st/ks-ws-manager/error"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
)

// HttpPost
// @Description 发送http-post请求
// @params reqUrl 请求的url
// @params req 请求参数
// @params res 要接收的返回参数
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 19:13
func (ws *WorkSpace) HttpPost(reqUrl string, reqData, res interface{}) error {
	token, err1 := ws.GetToken()
	if err1 != nil {
		return err1
	}
	client := resty.New()
	request := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(reqData)

	if res != nil {
		request.SetResult(res)
	}
	resp, err := request.Post(reqUrl)
	if err != nil || (resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated) {
		ks_error.DebugLog("请求k8s失败", err, "reqUrl=", reqUrl, "resp=", resp)
		_ = json.Unmarshal(resp.Body(), &res)
		return err
	}
	return nil
}

// HttpGet
// @Description 发送get请求
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 9:47
func (ws *WorkSpace) HttpGet(reqUrl string, reqData url.Values, res interface{}) error {
	token, err1 := ws.GetToken()
	if err1 != nil {
		return err1
	}
	client := resty.New()
	request := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token)
	if reqData != nil {
		request.SetQueryParamsFromValues(reqData)
	}
	if res != nil {
		request.SetResult(res)
	}
	resp, err := request.Get(reqUrl)
	if err != nil || (resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated) {
		ks_error.DebugLog("请求k8s失败", err, "reqUrl=", reqUrl, "resp=", resp)
		_ = json.Unmarshal(resp.Body(), &res)
		return err
	}
	return nil
}

// HttpDelete
// @Description 删除资源
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 12:16
func (ws *WorkSpace) HttpDelete(reqUrl string) error {
	token, err1 := ws.GetToken()
	if err1 != nil {
		return err1
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		Delete(reqUrl)

	if err != nil || (resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated) {
		ks_error.DebugLog("请求k8s失败", err, "reqUrl=", reqUrl, "resp=", resp)
		return err
	}
	return nil
}

func (ws *WorkSpace) HttpPut(reqUrl string, reqData, res interface{}) error {
	token, err1 := ws.GetToken()
	if err1 != nil {
		return err1
	}
	client := resty.New()
	request := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(reqData)

	if res != nil {
		request.SetResult(res)
	}
	resp, err := request.Put(reqUrl)
	if err != nil || (resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated) {
		ks_error.DebugLog("请求k8s失败", err, "reqUrl=", reqUrl, "resp=", resp)
		return err
	}
	return nil
}
