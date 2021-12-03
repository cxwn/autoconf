package utils

import (
	"os"
	"strings"
)

// 获取系统所有环境变量名称（保持大写），并返回一个字符串切片。

func GetEnvs() []string {
	var envs []string
	for _, value := range os.Environ() {
		value := strings.TrimSpace(strings.Split(value, "=")[0]) // 删除环境变量中的空字符串。
		envs = append(envs, value)
	}
	return envs
}
