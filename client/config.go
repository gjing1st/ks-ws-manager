// Path:
// FileName: manager.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/29$ 11:21$

package client

import ks_error "github.com/gjing1st/ks-ws-manager/error"

type k8sConfig struct {
	ksAddr            string //ks-apiserver地址
	username          string //ks管理员用户名
	password          string //密码
	configmapName     string //要创建的配置字典名称
	harborAddr        string //harbor仓库地址
	harborProject     string //harbor仓库项目名称
	repoName          string //ks中添加的应用仓库名称
	repoWorkspace     string //该应用仓库位于哪个企业空间
	apisixAdminaddr   string `default:"http://apisix-admin.default"`
	apisixGatewayAddr string `default:"http://apisix-admin.default"`
	apisixApiKey      string `default:"edd1c9f034335f136f87ad84b625c8f1"`
	apisixConsumerKey string `default:"i9ybgzkq88tisj0d"`
}

func NewK8SConfig() (c *k8sConfig) {
	c = &k8sConfig{
		ksAddr:            "http://ks-apiserver.kubesphere-system",
		username:          "admin",
		password:          "P@88w0rd",
		configmapName:     "tna-backend-conf",
		harborAddr:        "http://core.harbor.dked:30002",
		harborProject:     "tna",
		repoName:          "harbor-helm",
		repoWorkspace:     "tna",
		apisixAdminaddr:   "http://apisix-admin.default:9180",
		apisixGatewayAddr: "http://apisix-admin.default",
		apisixApiKey:      "edd1c9f034335f136f87ad84b625c8f1",
		apisixConsumerKey: "i9ybgzkq88tisj0d",
	}
	return c
}

func (c *k8sConfig) SetKsAddr(ksAddr string) *k8sConfig {
	c.ksAddr = ksAddr
	return c
}
func (c *k8sConfig) SetUsername(username string) *k8sConfig {
	c.username = username
	return c
}
func (c *k8sConfig) SetPassword(password string) *k8sConfig {
	c.password = password
	return c
}
func (c *k8sConfig) SetConfigmapName(configmapName string) *k8sConfig {
	c.configmapName = configmapName
	return c
}
func (c *k8sConfig) SetHarborAddr(harborAddr string) *k8sConfig {
	c.harborAddr = harborAddr
	return c
}
func (c *k8sConfig) SetHarborProject(harborProject string) *k8sConfig {
	c.harborProject = harborProject
	return c
}
func (c *k8sConfig) SetRepoName(repoName string) *k8sConfig {
	c.repoName = repoName
	return c
}
func (c *k8sConfig) SetRepoWorkspace(repoWorkspace string) *k8sConfig {
	c.repoWorkspace = repoWorkspace
	return c
}
func (c *k8sConfig) SetApisixAdminaddr(apisixAdminaddr string) *k8sConfig {
	c.apisixAdminaddr = apisixAdminaddr
	return c
}
func (c *k8sConfig) SetApisixGatewayAddr(apisixGatewayAddr string) *k8sConfig {
	c.apisixGatewayAddr = apisixGatewayAddr
	return c
}
func (c *k8sConfig) SetApisixConsumerKey(apisixConsumerKey string) *k8sConfig {
	c.apisixConsumerKey = apisixConsumerKey
	return c
}

func (c *k8sConfig) SetDebug(debug bool) *k8sConfig {
	ks_error.Debug = debug
	return c
}

func (c *k8sConfig) GetKsAddr() string {
	return c.ksAddr
}
func (c *k8sConfig) GetUsername() string {
	return c.username
}
func (c *k8sConfig) GetPassword() string {
	return c.password
}
func (c *k8sConfig) GetConfigmapName() string {
	return c.configmapName
}
func (c *k8sConfig) GetHarborAddr() string {
	return c.harborAddr
}
func (c *k8sConfig) GetHarborProject() string {
	return c.harborProject
}
func (c *k8sConfig) GetRepoName() string {
	return c.repoName
}
func (c *k8sConfig) GetRepoWorkspace() string {
	return c.repoWorkspace
}
