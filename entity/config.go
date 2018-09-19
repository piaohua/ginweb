package entity

// Config config struct
type Config struct {
	Version string                  `yaml:"version"`
	GinMode string                  `yaml:"mode"`
	JWT     bool                    `yaml:"jwt"`
	JWTSecret   string              `yaml:"jwtSecret"`
	WXAppid     string              `yaml:"wxAppid"`
	WXAppSecret string              `yaml:"wxAppSecret"`
	Env     string                  `yaml:"env"`
	Envs    map[string]Environments `yaml:"environments"`
	Xorm    map[string]XormEngine   `yaml:"xorm"`
}

// Environments env struct
type Environments struct {
	URL  string `yaml:"url"`
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Cors bool   `yaml:"cors"`
}
