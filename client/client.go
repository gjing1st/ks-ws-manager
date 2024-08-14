// Path: client
// FileName: client.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/29$ 15:27$

package client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	ks_error "github.com/gjing1st/ks-ws-manager/error"
	"github.com/gjing1st/ks-ws-manager/util"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// 缓存k8s-token
var tokenMap sync.Map

// WorkSpace 企业空间
type WorkSpace struct {
	*k8sConfig
}

func (ws *WorkSpace) SetConfig(ks *k8sConfig) *WorkSpace {
	ws.k8sConfig = ks
	return ws
}

// Set
// @Description  内存变量过期 类redis
// @param: key 变量名
// @param: value 变量值
// @param: exp 过期时间
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 19:00
func Set(key, value interface{}, exp time.Duration) {
	tokenMap.Store(key, value)
	time.AfterFunc(exp, func() {
		tokenMap.Delete(key)
	})
}

// GetToken
// @Description 获取ks token
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 18:45
func (ws *WorkSpace) GetToken() (token string, err error) {
	//查看缓存中是否已有token
	if tokenI, ok := tokenMap.Load("access_token"); ok {
		return util.String(tokenI), nil
	}
	var res *TokenResponse
	reqData := url.Values{}
	reqData.Add("grant_type", "password")
	reqData.Add("username", ws.username)
	reqData.Add("password", ws.password)
	reqData.Add("client_id", "kubesphere")
	reqData.Add("client_secret", "kubesphere")
	reqUrl := ws.ksAddr + "/oauth/token"
	client := resty.New()
	resp, err1 := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&res).
		SetBody(reqData.Encode()).
		Post(reqUrl)
	ks_error.DebugLog("reqData.Encode()----------", reqData.Encode())
	ks_error.DebugLog("reqUrl----------", reqUrl)

	if err1 != nil || resp.StatusCode() != http.StatusOK {
		ks_error.DebugLog("k8s获取token失败", err)
		_ = json.Unmarshal(resp.Body(), &res)
		if err1 == nil {
			err1 = errors.New("k8s获取token失败")
		}
		return "", err1
	}
	//写入缓存，过期时间减半
	//此处已加入Bearer 获取后直接使用
	if res != nil && res.AccessToken != "" {
		token = "Bearer " + res.AccessToken
	}
	if res == nil {
		return "", errors.New("请求失败")
	}
	Set("access_token", token, time.Second*time.Duration(res.ExpiresIn/2))
	return util.String(token), nil
}

// CreateWorkspaces
// @Description: 创建企业空间
// @param: unitId string 要创建的企业空间名称(用户单位唯一标识)
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 19:10
func (ws *WorkSpace) CreateWorkspaces(name string) error {
	reqUrl := ws.ksAddr + "/kapis/tenant.kubesphere.io/v1alpha2/workspaces"
	reqData := NewCreateWorkspacesRequest(name)
	return ws.HttpPost(reqUrl, reqData, nil)

}
func (ws *WorkSpace) CreateUser(name, passwd, role string) error {
	reqUrl := ws.ksAddr + "/kapis/iam.kubesphere.io/v1alpha2/users"
	reqData := NewCreateUserRequest(name, passwd, role)
	err := ws.HttpPost(reqUrl, reqData, nil)
	if err != nil {
		ks_error.DebugLog("创建用户失败", err)
	}
	return err
}

// CreateProject
// @Description: 创建企业空间下的项目
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 19:17
func (ws *WorkSpace) CreateProject(name, workspace string) error {
	reqData := NewCreateProjectRequest(name, workspace)
	reqUrl := ws.ksAddr + "/kapis/tenant.kubesphere.io/v1alpha2/workspaces/" + workspace + "/namespaces"
	return ws.HttpPost(reqUrl, reqData, nil)
}

// CreateRepos
// @Description: 添加应用仓库
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 19:20
// 该函数停止调用，统一使用tcs-master主管单位中的应用仓库，避免每个用户单位创建企业空间
func (ws *WorkSpace) CreateRepos(workspace, repoName, projectName string) (*CreateRepoResponse, error) {
	reqData := NewCreateRepoRequest(ws.harborAddr, repoName, projectName)
	reqUrl := ws.ksAddr + "/kapis/openpitrix.io/v1/workspaces/" + workspace + "/repos"
	res := &CreateRepoResponse{}
	err := ws.HttpPost(reqUrl, reqData, res)
	return res, err
}

// CreateConfigMap
// @Description 创建数据字典
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 19:32
func (ws *WorkSpace) CreateConfigMap(projectName, configmapName string, data interface{}) error {
	reqData := NewCreateConfigMap(projectName, data, configmapName)
	reqUrl := ws.ksAddr + "/api/v1/namespaces/" + projectName + "/configmaps"
	return ws.HttpPost(reqUrl, reqData, nil)
}

// GetAppList
// @Description 获取仓库中的应用列表
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 21:55
func (ws *WorkSpace) GetAppList(repoId string) (res *ReposAppsResponse, err error) {
	//reqData := GetRepoAppListRequest{
	//	OrderBy:    "create_time",
	//	Conditions: "status=active,repo_id=" + repoId,
	//	Reverse:    true,
	//}
	reqData := url.Values{}
	reqData.Add("orderBy", "create_time")
	reqData.Add("conditions", "status=active,repo_id="+repoId)
	reqData.Add("reverse", "true")
	res = &ReposAppsResponse{}
	reqUrl := ws.ksAddr + "/kapis/openpitrix.io/v1/apps"
	err = ws.HttpGet(reqUrl, reqData, res)
	return
}

// CreateApp
// @Description 部署应用，该方法涉及其他方法，使用应用列表中的第一个应用仓库
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 10:46
func (ws *WorkSpace) CreateApp(workspace, namespace, appName string) error {
	//获取仓库列表
	repos, err := ws.GetRepoList(workspace)
	if err != nil {
		ks_error.DebugLog("获取应用仓库列表失败", err)
		return err
	}
	if len(repos.Items) <= 0 {
		errs := "当前项目未添加应用仓库"
		ks_error.DebugLog("获取应用仓库列表失败", errs)
		return errors.New(errs)
	}
	repoId := repos.Items[0].RepoId
	//更新应用仓库
	_, _ = ws.UpdateRepo(workspace, repoId)
	time.Sleep(time.Second)
	//获取应用仓库中的应用
	apps, err := ws.GetAppList(repoId)
	if err != nil {
		ks_error.DebugLog("获取应用仓库中的应用失败", err)
		return err
	}
	if len(apps.Items) <= 0 {
		errs := "harbor中的该项目没有对应的helm应用"
		ks_error.DebugLog("获取应用仓库中的应用失败", errs)
		return errors.New(errs)
	}
	appid := apps.Items[0].Appid
	versionId := apps.Items[0].LatestAppVersion.VersionId
	//获取version helm yaml内容
	conf, err := ws.Files(appid, versionId)
	if err != nil {
		ks_error.DebugLog("获取应用对应的helm文件信息错误", err)
		return err
	}
	//开始创建应用，部署实际项目
	err = ws.CreateProjectApp(workspace, namespace, appid, versionId, conf, appName)
	if err != nil {
		ks_error.DebugLog("创建应用失败", err)
		return err
	}
	return err
}

// GetRepoList
// @Description 应用仓库列表
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 10:03
func (ws *WorkSpace) GetRepoList(workspaces string) (res *AppRepoResponse, err error) {
	workspaces = ws.repoWorkspace
	reqUrl := ws.ksAddr + "/kapis/openpitrix.io/v1/workspaces/" + workspaces + "/repos"
	res = &AppRepoResponse{}
	err = ws.HttpGet(reqUrl, nil, res)

	return
}

// UpdateRepo
// @Description
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 10:08
func (ws *WorkSpace) UpdateRepo(workspaces, repoId string) (message string, err error) {
	workspaces = ws.repoWorkspace
	reqUrl := ws.ksAddr + "/kapis/openpitrix.io/v1/workspaces/" + workspaces + "/repos/" + repoId + "/action"
	var reqData UpdateRequest
	reqData.Action = "index"
	res := &MessageResponse{}
	err = ws.HttpPost(reqUrl, reqData, res)
	if err != nil {
		message = res.Message
	}
	return
}

// Files
// @Description 获取版本文件，主要提取values.yaml文件数据
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 10:33
func (ws *WorkSpace) Files(appid, versionId string) (valuesYaml string, err error) {
	res := &FilesResponse{}
	reqUrl := ws.ksAddr + "/kapis/openpitrix.io/v1/apps/" + appid + "/versions/" + versionId + "/files"
	err = ws.HttpGet(reqUrl, nil, res)
	if err != nil {
		ks_error.DebugLog("获取版本文件失败", err)
		return
	}
	vaByte, err := base64.StdEncoding.DecodeString(res.Files.ValuesYaml)
	if err != nil {
		ks_error.DebugLog("获取版本文件失败", err)
		return "", err
	}
	valuesYaml = string(vaByte)
	return
}

// CreateProjectApp
// @Description 创建实际项目应用。企业空间-项目-应用
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 10:39
func (ws *WorkSpace) CreateProjectApp(workspace, namespace, appid, versionId, conf, name string) error {
	reqData := CreateProjectAppRequest{
		appid,
		conf,
		name,
		versionId,
	}
	reqUrl := ws.ksAddr + fmt.Sprintf("/kapis/openpitrix.io/v1/workspaces/%s/namespaces/%s/applications", workspace, namespace)
	err := ws.HttpPost(reqUrl, reqData, nil)
	if err != nil {
		ks_error.DebugLog("创建实际项目应用失败", err)

	}
	return err
}

// DeleteWS
// @Description 删除企业空间及其下资源
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 11:37
func (ws *WorkSpace) DeleteWS(name string) (err error) {
	reqUrl := ws.ksAddr + "/kapis/tenant.kubesphere.io/v1alpha3/workspacetemplates/" + name
	err = ws.HttpDelete(reqUrl)
	if err != nil {
		ks_error.DebugLog("删除企业空间及其下资源失败", err)
	}
	return err
}

// OpenNodePort
// @Description 开启nodePort外部访问模式
// @params  projectName string 项目名称/namespace名称
// @params  serviceName string 服务名称
// @params  serviceYml string 服务的yaml文件，此处为ws.GetServiceYaml返回的数据
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/4/2 16:40
func (ws *WorkSpace) OpenNodePort(projectName, serviceName string) (nodePort uint16, err error) {
	serviceYml := ws.GetServiceYaml(projectName, serviceName)
	//替换开启nodePort
	serviceYml = strings.Replace(serviceYml, `"type":"ClusterIP"`, `"type": "NodePort"`, -1)
	reqUrl := ws.ksAddr + fmt.Sprintf("/api/v1/namespaces/%s/services/%s", projectName, serviceName)
	res := &ServiceYmlResp{}
	err = ws.HttpPut(reqUrl, serviceYml, res)
	if err != nil {
		ks_error.DebugLog("开启nodeport失败", err, "namespaces", projectName, "serviceName", serviceName)
		return
	}
	if len(res.ServiceSpec.Ports) > 0 {
		nodePort = res.ServiceSpec.Ports[0].NodePort
	}
	return
}

// GetServiceYaml
// @Description 获取服务的yaml文件
// @params  projectName string 项目名称/namespace名称
// @params  serviceName string 服务名称
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/4/2 16:40
func (ws *WorkSpace) GetServiceYaml(projectName, serviceName string) string {
	reqUrl := ws.ksAddr + fmt.Sprintf("/api/v1/namespaces/%s/services/%s", projectName, serviceName)
	var res interface{}
	err := ws.HttpGet(reqUrl, nil, &res)
	if err != nil {
		ks_error.DebugLog("开启nodeport失败", err, "namespaces", projectName, "serviceName", serviceName)
	}
	resp, _ := json.Marshal(res)
	return string(resp)
}
