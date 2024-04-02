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
	se := `
{
  "kind": "Service",
  "apiVersion": "v1",
  "metadata": {
    "name": "tna-cert",
    "namespace": "tna4",
    "labels": {
      "app.kubernetes.io/instance": "tna4",
      "app.kubernetes.io/managed-by": "Helm",
      "app.kubernetes.io/name": "tna-cert",
      "app.kubesphere.io/instance": "tna4",
      "helm.sh/chart": "tna-cert-4.0.0"
    },
    "annotations": {
      "kubesphere.io/creator": "admin",
      "meta.helm.sh/release-name": "tna4",
      "meta.helm.sh/release-namespace": "tna4"
    },
    "resourceVersion": "12726499"
  },
  "spec": {
    "ports": [
      {
        "name": "http",
        "protocol": "TCP",
        "port": 8100,
        "targetPort": 8100
      }
    ],
    "selector": {
      "app.kubernetes.io/instance": "tna4",
      "app.kubernetes.io/name": "tna-cert"
    },
    "clusterIP": "10.233.53.63",
    "clusterIPs": [
      "10.233.53.63"
    ],
    "type": "NodePort",
    "sessionAffinity": "None",
    "ipFamilies": [
      "IPv4"
    ],
    "ipFamilyPolicy": "SingleStack",
    "internalTrafficPolicy": "Cluster"
  }
}
`
	ks := NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = WorkSpace{ks}
	res, _ := ws.OpenNodePort("tna4", "tna-cert", se)
	fmt.Println(res)
}

func TestWorkSpace_GetServiceYaml(t *testing.T) {
	ks := NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = WorkSpace{ks}
	ss := ws.GetServiceYaml("tna4", "tna-cert")
	nodePort, err := ws.OpenNodePort("tna4", "tna-cert", ss)
	fmt.Println(err)
	fmt.Println("nodePort======", nodePort)
}
