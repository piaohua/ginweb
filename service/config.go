package service

import (
	"project/entity"
	"project/libs"
)

var config *entity.Config

// Load init config
func Load(filePath string) {
	config = new(entity.Config)
	err := libs.LoadYaml(filePath, config)
	if err != nil {
		panic(err)
	}
}

// GetConfig get config
func GetConfig() *entity.Config {
	return config
}

// GetConfigCors get cors
func GetConfigCors() bool {
	if config.Envs == nil {
		return false
	}
	if v, ok := config.Envs[config.Env]; ok {
		return v.Cors
	}
	return false
}

// GetConfigJWT get jwt
func GetConfigJWT() bool {
	return config.JWT
}

// GetConfigMode get gin mode
func GetConfigMode() bool {
	return config.GinMode
}
