# ginweb based on [Gin](https://github.com/gin-gonic/gin)

## Installation

```sh
$ go get -u github.com/gin-gonic/gin
```

## Examples & Documentation

```
assets 静态资源
conf 配置文件
controllers 控制器
entity 实体结构
logs 日志文件
middleware 中间件
models 参数结构
libs 方法
pkg 包方法
routers 路由
service 业务逻辑
views 视图目录
tests 测试目录
vendor 依赖
ctrl 控制脚本
main.go 启动文件
```

```sh
$ ./ctrl build

$ curl -X POST foo:bar@127.0.0.1:8080/admin -H "Content-Type:application/json" -d "{\"value\":\"123\"}"

$ curl -X GET http://127.0.0.1:8080/user/foo
```

## TODOs

- tests
- documentation
