// Path:
// FileName: ws_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/29$ 17:15$

package ks_ws_manager

import (
	"fmt"
	"github.com/gjing1st/ks-ws-manager/client"
	"testing"
)

func TestDeployWSApp(t *testing.T) {
	//实例化ks相关信息
	ks := client.NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = client.WorkSpace{}
	ws.SetConfig(ks)
	//配置字典信息
	data := struct {
		ConfigYml string `json:"config.yml"`     //配置文件config.yml及其对应内容
		MysqlConf string `json:"mysql_conf.yml"` //mysql_conf.yml及其对应内容
	}{}
	data.ConfigYml = `
    port: 3306
    username: super_user
    maxidleconn: 10
    maxopenconn: 100`
	data.MysqlConf = `
	log:
	  # std|file
	  output: std
	  # trace|debug|info|warn
	  level: info
	  # 是否打印调用者信息
	  caller: true
	  # 日志目录
	  dir: ./log
	#web基础配置
	web:
	  port: 8801
	  #跨域开关
	  cors: true
	  #接口权限验证开关
	  auth: true`
	err := ws.DeployWSApp("test-ws", "app", data)
	fmt.Println("err", err)
}

func TestDropWS(t *testing.T) {
	//实例化ks相关信息
	ks := client.NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = client.WorkSpace{}
	ws.SetConfig(ks)
	err := ws.DropWorkSpace("test-ws")
	fmt.Println(err)
}
