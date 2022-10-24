![](https://img.shields.io/badge/version-v1.0.0-green.svg) &nbsp; ![](https://img.shields.io/badge/builder-success-green.svg) &nbsp;

> 📢📢📢 IOC依赖注入，对项目全局实例化对象进行管理，自动对结构体属性进行初始化，摆脱随处可见的 new

## 一、前言

### 🚀🚀安装

- Get

```shell
go get github.com/archine/ioc@v1.0.0
```

- Mod

```shell
# go.mod文件加入下面的一条
github.com/archine/ioc v1.0.0
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
	"github.com/archine/ioc"
)

type UserController struct {
	// 属性直接定义即可，属性名必须为大小，类型必须为指针(不然会对象拷贝，实际并不是同一个实例，避免出现这个问题,IOC内部直接拒绝非指针的属性进行注入)
	Service2 *Service

	// 也可以采用匿名的方式
	*Service
}

type Service struct {
	Age int
}

// CreateBean 由于该 Service被其他地方所引用，所以会自动触发该方法进行实例化，只会触发一次
func (s *Service) CreateBean() ioc.Bean {
	return &Service{20}
}

func main() {
	u := &UserController{}

	// 对实例进行依赖注入
	ioc.Inject(u)
}
```

### 2、手动实例化到IOC

在一些场景中，可能无法同时满足自动实例化的两个条件。这时我们可以通过手动放进 IOC 容器

```go
package main

import (
	"fmt"
	"github.com/archine/ioc"
)

type UserController struct {
	Service2 *Service
	*Service
}

type Service struct {
	Age int
}

func main() {
	u := &UserController{}

	// 手动放入实例到IOC
	ioc.SetBeans(&Service{20})

	ioc.Inject(u)
	fmt.Printf("%p", u.Service)
}
```

### 3、选择指定的 Bean
在一些场景中，我们实例中的属性为一个 interface，在项目中也存在多个该 interface 的实现类，且都在 IOC 容器中，这时，我们可以选择注入指定的实现类对应的实例
```go
package main

import (
	"fmt"
	"github.com/archine/ioc"
)

type UserController struct {
	Svc BaseService `@autowired:"main.Service1"` // 指定需要注入的 Bean Name，名称为 包名.结构体名
}

type BaseService interface {
	GetAge() int
}

// serivce1

type Service1 struct {
}

func (s *Service1) GetAge() int {
	return 10
}

// service2

type Service2 struct {
}

func (s *Service2) GetAge() int {
	return 20
}

func main() {
	u := &UserController{}
	// 手动放入两个实现类的实例
	ioc.SetBeans(&Service1{}, &Service2{})

	ioc.Inject(u)

	// 会输出 service1 的值 10
	fmt.Println(u.Svc.GetAge())
}
```
---