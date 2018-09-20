package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"syscall"
	"time"

	"ginweb/routers"
	"ginweb/service"

	"github.com/fvbock/endless"
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
	// glog.Infof("config: %#v", config)
	// log.Printf("config: %#v", config)
	if config == nil {
		log.Panic("config empty")
	}

    // 连接数据库
    for k, v := range config.Xorm {
        val := v
        service.NewXorm(k, &val)
        log.Printf("driverName %s connected", v.DriverName)
    }

	// 设置路由
	r := routers.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	addr := service.GetConfigAddr()
	log.Printf("Listen and Server in 0.0.0.0%s\n", addr)
	if addr == "" {
		log.Panic("Listen addr empty")
	}

	//r.Run(addr)

	// 关闭服务
	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20 // 1M

	server := endless.NewServer(addr, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
