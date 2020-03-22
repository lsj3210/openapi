package utils

import "sync"

// G 全局参数
var G sync.Map

// GetConf 获取配置信息
func GetConf() Config {
	value, _ := G.Load("conf")
	return value.(Config)
}
