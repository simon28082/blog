# blog
study record blog


## Install
```go
go mod download
go mod vendor
```

# 域驱动设计
## 目录结构
```
├── config                  配置目录
│   ├── *.yaml
├── database                数据库目录
│   └── migrations          迁移目录
├── docs                    文档
├── examples                示例
├── go.mod                  
├── go.sum
├── internal                内部代码包（无法对外部公用）
│   ├── app                 应用目录
│   │   ├── handlers        router处理目录，相当于controller
│   │   ├── provider.go     
│   │   ├── routes.go
│   │   └── transformer     资源转换层
│   ├── bootstrap           应用启动目录
│   │   └── bootstrap.go
│   ├── domain              域目录，所有业务逻辑及服务
│   │   ├── auth            
│   │   └── sms
│   └── pkg                 内部公共包
│       ├── app.go
│       └── README.md
├── main.go
├── pkg                     外部公共包，可共享于其它程序
├── README.md
├── testdata                测试数据目录
└── vendor                  第三方包
```

- `internal/domain` 主要是提供当前系统的所有业务逻辑，`Model`，仓库数据等
- `internal/app` 应用的入口
- 只有`app`中可以调用`domain`，而`domain`不能被`app`程序调用
- `service`可以调用`repository`，但`repository`不能调用`service`
- `app`只能调用`domain`，`domain`不能调用`app`
- `app`和`domain`层之间使用`struct`数据结构交互
- 严格来说只要`domain`完成了，基本上`app`只要调用`domain`多接口就全部完成了
- 每个service都使用`struct`申明，不使用函数


# swagger 参数格式
// 参数 字段名 类型（formData|query|header） 字段类型 是否必填(true|false) 字段说明
