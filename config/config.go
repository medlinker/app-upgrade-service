package config

import (
	"github.com/sosop/libconfig"
)

var (
	iniConfig *libconfig.IniConfig
	mode      string
)

// InitConf read configuration
func InitConf(path string) {
	iniConfig = libconfig.NewIniConfig(path)
	mode = iniConfig.GetString("mode", "prod")
}

// GetString 获取配置
func GetString(key string, withMode bool, defaultValue ...string) string {
	if withMode {
		return iniConfig.GetString(mode+"::"+key, defaultValue...)
	}
	return iniConfig.GetString(key, defaultValue...)
}
