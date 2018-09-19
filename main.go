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
	// 加载配置
	service.Load(*configPath)

	// 初始化
	config := service.GetConfig()
	glog.Infof("config: %#v", config)
	log.Printf("config: %#v", config)
    if config == nil {
		log.Panic("config empty")
    }

    // 设置路由
	r := routers.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
    addr := service.GetConfigAddr()
    log.Printf("Listen and Server in 0.0.0.0%s\n", addr)
    if addr == "" {
        log.Panic("Listen addr empty")
    }

	r.Run(addr)

	//TODO 关闭服务
	//TODO 热更新
}
