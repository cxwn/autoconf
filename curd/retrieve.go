package curd

import (
	"autoconf/utils"
	"log"
	"strings"
)

/*
在配置文件中找到需要更新的配置项，将配置项存于字符串切片中。
*/

// 判断在目标切片中是否存在指定字符串。
func isExist(list []string, str string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}

// Retrieve 返回需要更新的配置项，并将其值存放在一个字符串切片中。
func (conf Config) Retrieve() []string {
	var targets []string
	var keys []string
	for key := range conf.Tidy() {
		keys = append(keys, key) // 从配置文件中读取的可配置项。
	}

	for _, osEnv := range utils.GetEnvs() {
		if strings.HasPrefix(osEnv, conf.Prefix) {
			trimEnv := strings.TrimPrefix(osEnv, conf.Prefix)
			if isExist(keys, trimEnv) {
				targets = append(targets, osEnv)
			} else if !isExist(keys, trimEnv) {
				log.Printf("环境变量【%s】在配置文件【%s】中未查找到，请核实系统环境变量【%s】配置是否有误，配置将不在该配置文件中进行更新！\n", trimEnv, conf.Name, osEnv)
			}
		}
	}
	return targets
}
