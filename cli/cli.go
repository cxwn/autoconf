package cli

import (
	"autoconf/curd"
	flag "github.com/spf13/pflag"
	"strings"
)

func InitCli(con *curd.Config) {
	var name *string = flag.StringP("file", "f", "", "配置文件名，包括路径，不能为空")
	var annotation *string = flag.StringP("annotation", "a", "#", "注释符，仅适用于单行注释。注释符后的内容将不处理")
	var prefix *string = flag.StringP("prefix", "p", "", "环境变量前缀，不能为空")
	var connector *string = flag.StringP("connector", "c", "", "配置项内部连接符，默认无连接符。如：redis.port=6379，配置项内连接符为'.'")
	var equal *string = flag.StringP("equal", "e", "=", "赋值连接符,默认为'='。 yaml、json类型的配置文件，该符号应替换为':'")
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	flag.Lookup("annotation").NoOptDefVal = "#"
	flag.Lookup("equal").NoOptDefVal = "="
	flag.Lookup("connector").NoOptDefVal = ""
	flag.Parse()
	con.Name = *name
	con.Comment = *annotation
	con.Prefix = *prefix
	con.Connection = *connector
	con.EqualitySigned = *equal
}

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}
