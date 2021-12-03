package curd

import (
	"strings"
)

// IsContainKeyword 函数，判断给定字符串 target 是否包含指定关键字。
func IsContainKeyword(target string) bool {
	keywords := [...]string{"PASS", "PASSWORD", "PASSWD", "PSWD"}
	for _, keyword := range keywords {
		if strings.Contains(target, keyword) {
			return true
		}
	}
	return false
}
