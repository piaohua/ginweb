package service

import (
	"log"
	"os"
	"time"

	"ginweb/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// xormEngines
var xormEngines = make(map[string]*entity.XormEngine)

// SetEngine set Engine
func SetEngine(k string, e *entity.XormEngine) {
	xormEngines[k] = e
}

// GetEngine get engine
func GetEngine(key string) (e *entity.XormEngine) {
	if v, ok := xormEngines[key]; ok {
		return v
	}
	return xormEngines["default"]
}

// NewXorm new xorm
func NewXorm(key string, engine *entity.XormEngine) {
	if len(key) == 0 {
		log.Panic("key-empty")
	}
	InitXorm(engine)
	SetEngine(key, engine)
}

// InitXorm init xorm
func InitXorm(e *entity.XormEngine) {
	var err error
	var engine *xorm.Engine
	engine, err = xorm.NewEngine(e.DriverName, e.DataSourceName)
	if err != nil {
		log.Panicf("new engine err: %v\n", err)
	}

	engine.ShowSQL(e.ShowSQL) //在控制台打印出生成的SQL语句
	//日志默认显示级别为INFO
	engine.Logger().SetLevel(core.LogLevel(e.LoggerLevel)) //在控制台打印调试及以上的信息

	sql2log(e) //日志写入文件

	engine.SetMaxIdleConns(e.MaxIdleConns) //设置连接池的空闲数大小,default is 2
	engine.SetMaxOpenConns(e.MaxOpenConns) //设置最大打开连接数

	engine.SetMapper(core.GonicMapper{})

	//设置时区
	engine.TZLocation, err = time.LoadLocation(e.Location)
	if err != nil {
		log.Panicf("set location err: %v\n", err)
	}

	e.Engine = engine
	go flushDaemon(e)
}

// sync2 同步结构
func sync2(e *entity.XormEngine, beans ...interface{}) {
	err := e.Sync2(beans)
	if err != nil {
		log.Panicf("sync2 err: %v\n", err)
	}
}

func flushDaemon(e *entity.XormEngine) {
	for _ = range time.NewTicker(30 * time.Second).C {
		err := e.Ping()
		if err != nil {
			log.Printf("sync2 err: %v\n", err)
			//TODO
		}
	}
}

//sql save to file
func sql2log(e *entity.XormEngine) {
	if len(e.LoggerFile) == 0 {
		return
	}
	f, err := os.Create(e.LoggerFile)
	if err != nil {
		log.Panicf("sql2log err: %v\n", err)
	}
	e.SetLogger(xorm.NewSimpleLogger(f))
}
