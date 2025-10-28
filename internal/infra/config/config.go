// Package config 配置
package config

import (
	"github.com/spf13/viper"
)

var vp = viper.New()

const (
	KeyDatabase = "INFRA_DATABASE"
	KeyLogx     = "INFRA_LOGX"
	KeyCfpx     = "INFRA_CFPX_FILE"
	KeyEndpoint = "ENDPOINT"
)

// GetDatabase 数据库配置
func GetDatabase() string {
	return vp.GetString(KeyDatabase)
}

// GetLogx 获取日志配置
func GetLogx() string {
	return vp.GetString(KeyLogx)
}

// SetConfig 设置配置
func SetConfig(key string, val string) {
	vp.Set(key, val)
}

func GetCfpx() string {
	return vp.GetString(KeyCfpx)
}

// GetEndpoint 获取服务端地址
func GetEndpoint() string {
	endpoint := vp.GetString(KeyEndpoint)
	if endpoint == "" {
		endpoint = ":8080"
	}
	return endpoint
}
