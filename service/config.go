package service

import (
	"ginweb/entity"
	"ginweb/libs"
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

// GetConfigJWTSecret get JWTSecret
func GetConfigJWTSecret() string {
	return config.JWTSecret
}

// GetConfigMode get gin mode
func GetConfigMode() string {
	return config.GinMode
}

// GetConfigAddr get addr
func GetConfigAddr() string {
	if config.Envs == nil {
		return ""
	}
	if v, ok := config.Envs[config.Env]; ok {
		return v.Addr
	}
	return ""
}

// GetConfigEnv get env
func GetConfigEnv() *entity.Environments {
	if config.Envs == nil {
		return nil
	}
	if v, ok := config.Envs[config.Env]; ok {
        val := v
		return &val
	}
	return nil
}
