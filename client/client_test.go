// Path: client
// FileName: client_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/29$ 16:07$

package client

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	ks := NewK8SConfig()
	ws := WorkSpace{ks}
	_ = ws.CreateWorkspaces("a")
}

func TestWorkSpace_OpenNodePort(t *testing.T) {
	ks := NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = WorkSpace{ks}
	nodePort, err := ws.OpenNodePort("tna4", "tna-cert")
	fmt.Println(err)
	fmt.Println("nodePort======", nodePort)
}

func TestWorkSpace_GetServiceYaml(t *testing.T) {
	ks := NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = WorkSpace{ks}
	ss := ws.GetServiceYaml("tna4", "tna-cert")
	fmt.Println(ss)
}
