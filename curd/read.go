package curd

import (
	"bufio"
	"log"
	"os"
)

/*
  按行读取配置文件，并将配置文件的每一行添加到一个字符串数组，最终返回该配置文件剔除注释、空白后的字符串数组。需要注意的是：最后一行可能没有换号符，导致最后一行无法读取。
*/

// ReadConfFileByLine 按行读取配置文件，并将读取到的结果按行存入一个字符串切片中。
func (conf Config) ReadConfFileByLine() []string {
	var lines []string

	file, err := os.Open(conf.Name)
	if err != nil {
		log.Printf("配置文件【%s】打开失败，请核实！\n", conf.Name)
		os.Exit(128) // 目标文件读取失败，错误代码：128
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
