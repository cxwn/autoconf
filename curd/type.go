package curd

type Config struct {
	Name           string // 文件名，包括文件路径，例如：/etc/redis/redis.conf。
	Comment        string // 本配置文件的注释形式，即注释符号，如使用“#”注释，则值为："#"，在读取配置文件时会忽略注释内容。
	Prefix         string // 配置文件对应的前缀。
	Connection     string // 环境变量中的链接符，有的配置文件采用“.”之类的连接符来连接配置文件中的变量。如：redis.port=6379。默认值为空字符串。
	EqualitySigned string // 配置文件中的等号形式，一般默认为“=”，少数配置文件为“:”。
}
