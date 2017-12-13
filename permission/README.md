# 权限中心的接口调用说明

## 获取

```bash
go get -v github.com/antlinker/sdk/permission/...
```

## 使用范例

```go
package main

import (
	"context"
	"fmt"

	"github.com/antlinker/sdk/permission/client"
	"github.com/antlinker/sdk/permission/proto/permission"
)

func main() {
	client.InitRPCClient(&client.RPCConfig{
		Addr:   "127.0.0.1:50051",
		AppID:  "app_id",
		AppKey: "app_key",
	})

	userInfo, err := client.PermissionClient().LoginWithUserName(context.Background(), &permission.LoginWithUserNameRequest{
		Username: "admin",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfo)
}
```

## 接口说明

### 权限客户端（`PermissionClient`）

* **LoginWithUserName** - 使用用户名和密码登录
* **LoginWithUserID** - 使用用户ID和密码登录
* **QueryUserMenus** - 查询用户的功能菜单