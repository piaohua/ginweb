package main

import (
	"flag"
	"runtime"
	"log"

	"ginweb/routers"
	"ginweb/service"

	"github.com/golang/glog"
)

var (
	configPath = flag.String("config", "./conf/config.yaml", "If non-empty, start with this config")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	defer glog.Flush()
	//加载配置
	service.Load(*configPath)

	//初始化
	config := service.GetConfig()
	glog.Infof("config: %#v", config)
	log.Printf("config: %#v", config)

    // 设置路由
	r := routers.SetupRouter()

	// 设置gin开发模式
	gin.SetMode(config.GinMode)

	//判断是否开启跨域模式
	if service.GetConfigCors {
		router.Use(utils.Cors())
	}

	//判断是否开启jwt验证
	if service.GetConfigJWT() {
		router.Use(utils.JWTAuth())
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(service.GetConfigAddr())

	//TODO 关闭服务
	//TODO 热更新
}
