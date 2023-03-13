```
├── main.go
├── config
│   ├── config.go
│   └── env
│       ├── development.yaml
│       ├── production.yaml
│       └── staging.yaml
├── controllers
│   ├── user_controller.go
│   └── ...
├── models
│   ├── user_model.go
│   └── ...
├── services
│   ├── user_service.go
│   └── ...
├── repositories
│   ├── user_repository.go
│   └── ...
├── middlewares
│   ├── auth_middleware.go
│   └── ...
├── utils
│   ├── response.go
│   └── ...
└── routes
    ├── router.go
    └── ...
```

`main.go`是应用程序的入口文件，它初始化应用程序并启动Web服务器。

`config`目录包含应用程序的配置文件，它们根据不同的环境进行分组（如开发环境、生产环境、测试环境等）。

`controllers`目录包含处理HTTP请求的控制器，它们接收HTTP请求并返回HTTP响应。

`models`目录包含应用程序的数据模型，它们通常与数据库表对应。

`services`目录包含业务逻辑的服务层，它们处理控制器传递的数据并返回结果。

`repositories`目录包含与数据库交互的代码，它们提供了数据访问的接口。

`middlewares`目录包含应用程序的中间件，它们在HTTP请求处理之前或之后执行一些操作，如身份验证、日志记录等。

`utils`目录包含一些通用的工具函数或结构体，如HTTP响应结构体、日期时间格式化函数等。

`routes`目录包含应用程序的路由配置，它定义了HTTP请求的URL路径和对应的控制器、中间件等。
