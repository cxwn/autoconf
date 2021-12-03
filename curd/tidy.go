package curd

// 对读取到的配置文件进行整理：剔除注释，提取变量配置文件中可配置项的名称，并将连接符中的连接符替换为下划线、字符转换为全大写，插入到字符以变量的全大写为键、变量为值的 map 中。

import (
	"strings"
)

func (conf Config) Tidy() map[string]string {
	targets := make(map[string]string)
	lines := conf.ReadConfFileByLine()
	for _, line := range lines {
		if strings.Contains(line, conf.EqualitySigned) { // 仅选取包含赋值号的行进行处理
			if !strings.Contains(line, conf.Comment) {
				key := strings.TrimSpace(strings.Split(line, conf.EqualitySigned)[0])
				value := strings.TrimSpace(strings.Split(line, conf.EqualitySigned)[0])
				if strings.Contains(key, conf.Connection) && conf.Connection != "" {
					count := strings.Count(key, conf.Connection)
					key = strings.ToUpper(key)
					key = strings.Replace(key, conf.Connection, "_", count)
					targets[key] = value
				} else {
					key = strings.ToUpper(key)
					targets[key] = value
				}
			} else if strings.Contains(line, conf.Comment) {
				available := strings.TrimSpace(strings.Split(line, conf.Comment)[0])
				if available != "" && strings.Contains(available, conf.EqualitySigned) {
					key := strings.TrimSpace(strings.Split(available, conf.EqualitySigned)[0])
					value := strings.TrimSpace(strings.Split(available, conf.EqualitySigned)[1])
					if strings.Contains(key, conf.Connection) && conf.Connection != "" {
						count := strings.Count(key, conf.Connection)
						key = strings.ToUpper(key)
						key = strings.Replace(key, conf.Connection, "_", count)
						targets[key] = value
					} else {
						key = strings.ToUpper(key)
						targets[key] = value
					}
				}
			}
		}
	}
	return targets
}
