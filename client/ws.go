// Path: client
// FileName: ws.go
// Created by dkedTeam
// Author: GJing
// Date: 2024/3/29$ 15:29$

package client

type DeployInNewWS interface {
	CreateWorkspaces(unitId string) error
	CreateProject(name, workspace string) error
	CreateRepos(workspace, repoName, projectName string) (*CreateRepoResponse, error)
	CreateConfigMap(projectName, data string) error
	CreateApp(workspace, namespace, appName string) error
	DeleteWS(name string) error
}

// DeployWSApp
// @Description 创建企业空间并部署应用(添加用户单位时请求k8s入口)
// @params  unitId string 用户单位标识
// @params  unitName string 用户单位名称
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 11:17
func (ws *WorkSpace) DeployWSApp(unitId, appName string, configmapData interface{}) (err error) {
	//1. 创建企业空间
	err = ws.CreateWorkspaces(unitId)
	if err != nil {
		//创建失败
		return
	}
	// 2.创建项目
	err = ws.CreateProject(unitId, unitId)
	if err != nil {
		//删除企业空间下所有资源，方便下次添加
		_ = ws.DeleteWS(unitId)
	}
	//统一使用主管单位应用仓库。需要配置主管单位配置文件k8sConfig.reponame和k8sConfig.repoworkspace,
	// 3.创建应用仓库
	//repo, err := ws.CreateRepos(unitId, repoName, config.K8sConfig.HarborProject)
	//if err != nil {
	//	//删除企业空间下所有资源，方便下次添加
	//	_ = ws.DeleteWS(unitId)
	//}
	//fmt.Println("repo", repo.RepoId)
	// 4.创建数据字典
	//err = ws.CreateConfigMap(unitId, assembleDataDictUseObject(unitId, unitName))
	err = ws.CreateConfigMap(unitId, ws.configmapName, configmapData)
	if err != nil {
		//删除企业空间下所有资源，方便下次添加
		_ = ws.DeleteWS(unitId)
	}
	// 5.创建应用
	err = ws.CreateApp(unitId, unitId, appName)
	if err != nil {
		//删除企业空间下所有资源，方便下次添加
		_ = ws.DeleteWS(unitId)
	}
	return err
}

// DropWorkSpace
// @Description 删除企业空间及所有资源
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/27 12:14
func (ws *WorkSpace) DropWorkSpace(unitId string) (err error) {
	//删除企业空间下所有资源，方便下次添加
	err = ws.DeleteWS(unitId)
	return
}

//func assembleDataDictUseObject(unitId, unitName string) string {
//	conf := config.Config
//	fmt.Println("~~~~~", conf)
//	conf.Database.DBName = unitId
//	confByte, _ := yaml.Marshal(conf)
//	return string(confByte)
//}

// DeployWSProject
// @Description 创建企业空间和项目
// @params wsName string 企业空间名称
// @params projectName string 项目/namespace名称
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2024/3/30 10:14
func (ws *WorkSpace) DeployWSProject(wsName, projectName string) (err error) {
	//1. 创建企业空间
	err = ws.CreateWorkspaces(wsName)
	if err != nil {
		//创建失败
		return
	}
	// 2.创建项目
	err = ws.CreateProject(projectName, wsName)
	if err != nil {
		//删除企业空间下所有资源，方便下次添加
		_ = ws.DeleteWS(wsName)
	}
	return err
}
