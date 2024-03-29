package client

import (
	"fmt"
	"time"
)

const (
	repoName      = "harbor-helm"
	appName       = "tna"
	repoWorkSpace = "tna"
)

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
type Result struct {
	Code int
	Err  error
}

// AppListResponse k8s应用列表
type AppListResponse struct {
	Items []struct {
		Name    string `json:"name"`
		Cluster struct {
			AppId      string `json:"app_id"`
			ClusterId  string `json:"cluster_id"`
			CreateTime string `json:"create_time"`
			Env        string `json:"env"`
			Name       string `json:"name"`
			Owner      string `json:"owner"`
			RuntimeId  string `json:"runtime_id"`
			Status     string `json:"status"`
			StatusTime string `json:"status_time"`
			VersionId  string `json:"version_id"`
			Zone       string `json:"zone"`
		} `json:"cluster"`
		Version struct {
			AppId     string `json:"app_id"`
			Name      string `json:"name"`
			VersionId string `json:"version_id"`
		} `json:"version"`
		App struct {
			AppId       string `json:"app_id"`
			CategorySet string `json:"category_set"`
			ChartName   string `json:"chart_name"`
			Name        string `json:"name"`
		} `json:"app"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

// VersionResponse 版本信息
type VersionResponse struct {
	Items []struct {
		Active      bool   `json:"active"`
		AppId       string `json:"app_id"`
		CreateTime  string `json:"create_time"`
		Name        string `json:"name"`
		Owner       string `json:"owner"`
		PackageName string `json:"package_name"`
		Status      string `json:"status"`
		StatusTime  string `json:"status_time"`
		UpdateTime  string `json:"update_time"`
		VersionId   string `json:"version_id"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

type FilesResponse struct {
	Files struct {
		Helmignore string `json:".helmignore"`
		ValuesYaml string `json:"values.yaml"`
	} `json:"files"`
	VersionId string `json:"version_id"`
}

type UpVersionReq struct {
	AppId     string `json:"app_id"`
	ClusterId string `json:"cluster_id"`
	Cluster   string `json:"cluster"`
	Conf      string `json:"conf"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Owner     string `json:"owner"`
	VersionId string `json:"version_id"`
	Workspace string `json:"workspace"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type K8sAppAndVersion struct {
	App     *AppListResponse
	Version *VersionResponse
}

// AppRepoResponse 应用列表
type AppRepoResponse struct {
	Items []struct {
		RepoId string `json:"repo_id"`
		Url    string `json:"url"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

type UpdateRequest struct {
	Action string `json:"action"`
}

// CreateWorkspacesRequest 创建企业空间请求参数
type CreateWorkspacesRequest struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   `json:"metadata"`
	Spec       `json:"spec"`
}

type Metadata struct {
	Name        string `json:"name"`
	Annotations `json:"annotations"`
}
type Annotations struct {
	AliasName   string `json:"kubesphere.io/alias-name"`
	Creator     string `json:"kubesphere.io/creator"`
	Description string `json:"kubesphere.io/description"`
}
type Spec struct {
	Template `json:"template"`
}
type Template struct {
	TemplateSpec `json:"spec"`
}
type TemplateSpec struct {
	Manager string `json:"manager"`
}

// NewCreateWorkspacesRequest
// @description: 初始化创建企业空间请求参数
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/8/30 20:45
// @success:
func NewCreateWorkspacesRequest(name string) (req CreateWorkspacesRequest) {
	spec := Spec{
		Template{
			TemplateSpec{
				Manager: "admin",
			},
		},
	}
	metadata := Metadata{
		name,
		Annotations{
			AliasName:   name,
			Creator:     "admin",
			Description: "用户单位" + name + "的企业空间",
		},
	}
	req = CreateWorkspacesRequest{
		"tenant.kubesphere.io/v1alpha2",
		"WorkspaceTemplate",
		metadata,
		spec,
	}
	return req
}

type CreateProjectRequest struct {
	ApiVersion            string `json:"apiVersion"`
	Kind                  string `json:"kind"`
	CreateProjectMetadata `json:"metadata"`
}

type CreateProjectMetadata struct {
	Name                string `json:"name"`
	Annotations         `json:"annotations"`
	CreateProjectLabels `json:"labels"`
}
type CreateProjectLabels struct {
	Workspace string `json:"kubesphere.io/workspace"`
}

// NewCreateProjectRequest
// @description: 创建项目初始化请求参数
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/9/1 10:39
// @success:
func NewCreateProjectRequest(name, workspace string) (req CreateProjectRequest) {
	labels := CreateProjectLabels{
		workspace,
	}
	annotations := Annotations{
		name,
		"admin",
		"用户单位" + name + "的项目",
	}
	metadata := CreateProjectMetadata{
		name,
		annotations,
		labels,
	}
	req = CreateProjectRequest{
		"v1",
		"Namespace",
		metadata,
	}
	return
}

// CreateRepoRequest 创建应用仓库请求体
type CreateRepoRequest struct {
	Status     string   `json:"app_default_status"`
	Credential string   `json:"credential"`
	Name       string   `json:"name"`
	Providers  []string `json:"providers"`
	RepoType   string   `json:"repoType"`
	SyncPeriod string   `json:"sync_period"`
	Type       string   `json:"type"`
	Url        string   `json:"url"`
	Visibility string   `json:"visibility"`
}

type CreateRepoResponse struct {
	RepoId string `json:"repo_id"`
}

func NewCreateRepoRequest(url, repoName, projectName string) (req CreateRepoRequest) {
	providers := []string{"kubernetes"}
	req = CreateRepoRequest{
		"active",
		"{}",
		repoName,
		providers,
		"Helm",
		"30m",
		"http",
		fmt.Sprintf("%s/chartrepo/%s", url, projectName),
		"public",
	}
	return req
}

type ReposAppsResponse struct {
	Items []struct {
		Appid            string `json:"app_id"`
		Name             string `json:"name"`
		LatestAppVersion struct {
			Appid       string    `json:"app_id"`
			Name        string    `json:"name"`
			PackageName string    `json:"package_name"`
			UpdateTime  time.Time `json:"update_time"`
			VersionId   string    `json:"version_id"`
		} `json:"latest_app_version"`
		RepoId string `json:"repo_id"`
		Url    string `json:"url"`
	} `json:"items"`
	TotalCount int `json:"total_count"`
}

// CreateProjectAppRequest 创建实际应用项目请求参数
type CreateProjectAppRequest struct {
	Appid     string `json:"app_id"`
	Conf      string `json:"conf"`
	Name      string `json:"name"`
	VersionId string `json:"version_id"`
}
type ConfigMapMetadata struct {
	Annotations struct {
		KubesphereIoCreator string `json:"kubesphere.io/creator"`
	} `json:"annotations"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Labels    struct {
	} `json:"labels"`
}
type CreateConfigMapRequest struct {
	APIVersion        string `json:"apiVersion"`
	Kind              string `json:"kind"`
	ConfigMapMetadata `json:"metadata"`

	Spec struct {
		Template struct {
			Metadata struct {
				Labels struct {
				} `json:"labels"`
				Annotations struct {
					KubesphereIoCreator string `json:"kubesphere.io/creator"`
				} `json:"annotations"`
			} `json:"metadata"`
		} `json:"template"`
	} `json:"spec"`
	Data struct {
		ConfigYml string `json:"config.yml"`
	} `json:"data"`
}

// NewCreateConfigMap
// @Description 创建数据字典
// @params projectName string 项目名称
// @params data string 数据字典的值
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/26 19:48
func NewCreateConfigMap(projectName, data, configMapName string) (req CreateConfigMapRequest) {
	metadata := ConfigMapMetadata{
		Annotations: struct {
			KubesphereIoCreator string `json:"kubesphere.io/creator"`
		}(struct{ KubesphereIoCreator string }{KubesphereIoCreator: "admin"}),
		Name:      configMapName,
		Namespace: projectName,
	}
	req = CreateConfigMapRequest{
		APIVersion:        "v1",
		Kind:              "ConfigMap",
		ConfigMapMetadata: metadata,
		Data: struct {
			ConfigYml string `json:"config.yml"`
		}(struct{ ConfigYml string }{ConfigYml: data}),
	}
	return
}

type GetRepoAppListRequest struct {
	Conditions string `json:"conditions"`
	Reverse    bool   `json:"reverse"`
	OrderBy    string `json:"orderBy"`
}

type GetRepoAppListResponse struct {
}
