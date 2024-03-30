# ks-ws-manger
基于kubersphere企业空间的多租户管理服务
##

## 使用
### 创建企业空间
- 需要传入唯一id，该id将被用做企业空间和项目名称
- 配置字典传入配置文件和对应内容(可定义一个结构体传入)参考示例
```go
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
		ConfigYml string `json:"config.yml"` //配置文件config.yml及其对应内容
		MysqlConf string `json:"mysql_conf.yml"` //mysql_conf.yml及其对应内容
	}{}
	data.ConfigYml = ``
	data.MysqlConf = ``
	err := ws.DeployWSApp("test-ws","app", data)
	fmt.Println("err", err)
}
```
### 删除企业空间
```go
import (
    "fmt"
    "github.com/gjing1st/ks-ws-manager/client"
    "testing"
)
func TestDropWS(t *testing.T) {
	//实例化ks相关信息
	ks := client.NewK8SConfig()
	ks.SetKsAddr("http://192.168.200.80:31511").SetDebug(true)
	var ws = client.WorkSpace{}
	ws.SetConfig(ks)
	err := ws.DropWorkSpace("test-ws")
	fmt.Println(err)
}
```