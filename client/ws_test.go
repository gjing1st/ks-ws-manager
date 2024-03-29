// Path: client
// FileName: ws_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/29$ 16:20$

package client

import (
	"fmt"
	"testing"
)

func TestDeployWSApp(t *testing.T) {
	ks := NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = WorkSpace{ks}
	data := struct {
		ConfigYml string `json:"config.yml"`
		MysqlConf string `json:"mysql_conf.yml"`
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
