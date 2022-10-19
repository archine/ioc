![](https://img.shields.io/badge/version-v1.0.0-green.svg) &nbsp; ![](https://img.shields.io/badge/builder-success-green.svg) &nbsp;

> 📢📢📢 IOC依赖注入，对项目全局实例化对象进行管理，自动对结构体属性进行初始化，摆脱随处可见的 new

## 一、前言

### 1、基础条件

引入前需检查本地环境是否设置了 `GOPRIVATE`变量, 如未配置需先进行配置，可通过配置操作系统的变量或者通过在命令行执行下方的命令，否则无法拉取包

```shell
go env -w GOPRIVATE=gitlab.avatarworks.com

go env -w GOINSECURE=gitlab.avatarworks.com
```

### 2、🚀🚀安装

- Get

```shell
go get gitlab.avatarworks.com/servers/component/hj-ioc@v1.0.0
```

- Mod

```shell
# go.mod文件加入下面的一条
gitlab.avatarworks.com/servers/component/hj-ioc v1.0.0
# 命令行在该项目目录下执行
go mod tidy
```
## 二、使用说明
### 1、自动实例化
满足自动实例化，需要满足两个前置条件，一是必须实现 ``CreateBean()`` 方法，二是必须有地方对其进行引用，满足其一均不会进行自动实例化。看下面的使用案例

```go
package main

import (
	"fmt"
	ioc "gitlab.avatarworks.com/servers/component/hj-ioc"
)

type UserController struct {
	Service *Service
}

type Service struct {
}

func main() {
	u := &UserController{}
	ioc.SetBeans(&Service{30})
	ioc.Inject(u)
	fmt.Printf("%p", u.Service)
}
```
