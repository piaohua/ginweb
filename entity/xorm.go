package entity

import "github.com/go-xorm/xorm"

// XormEngine XormEngine struct
type XormEngine struct {
    *xorm.Engine `yaml:"-"`
    DriverName string `yaml:"driverName"`
    DataSourceName string `yaml:"dataSourceName"`
    ShowSQL bool `yaml:"showSQL"`
    LoggerLevel int `yaml:"loggerLevel"`
    MaxIdleConns int `yaml:"maxIdleConns"`
    MaxOpenConns int `yaml:"maxOpenConns"`
    Location string `yaml:"location"`
    LoggerFile string `yaml:"loggerFile"`
}
