// Path: client
// FileName: client_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/29$ 16:07$

package client

import "testing"

func TestCreate(t *testing.T) {
	ks := NewK8SConfig()
	ws := WorkSpace{ks}
	ws.CreateWorkspaces("a")
}
