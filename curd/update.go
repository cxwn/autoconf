package curd

import (
	"autoconf/cmrsa"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// Update 调用系统接口，更新
func (conf Config) Update(privateKey string) bool {
	for _, value := range conf.Retrieve() {
		var newString string
		oldString := conf.Tidy()[strings.TrimPrefix(value, conf.Prefix)]
		if IsContainKeyword(strings.ToUpper(oldString)) && os.Getenv(value) != "" {
			decryptString, err := cmrsa.PrivateKeyDecrypt(os.Getenv(value), privateKey)
			if err == nil {
				newString = decryptString
			} else {
				fmt.Println("加密字段解码失败，请核实！")
				return false
			}
		} else {
			newString = os.Getenv(value)
		}
		err := updateConfiguration(conf, oldString, newString, conf.EqualitySigned)
		if err != nil {
			log.Printf("配置文件【%s】中的配置变量【%s】更新失败，错误信息为：【%s】，请检查！\n", conf.Name, conf.Tidy()[strings.TrimPrefix(value, conf.Prefix)], err)
			return false
		}
		fmt.Printf("配置文件【%s】中的配置变量【%s】更新成功，更新后的配置项值为(加密字段不解码)：【%s】！\n", conf.Name, conf.Tidy()[strings.TrimPrefix(value, conf.Prefix)], os.Getenv(value))
	}
	return true
}


// 调用系统 sed 命令处理配置文件。
/*
func updateConfiguration(file string, oldString string, newString string, equalitySigned string) error {

	// 如果被替换的配置项中含".","?","&"，则进行转义。
	if strings.Contains(oldString, `.`) {
		oldString = strings.Replace(oldString, ".", `\.`, -1)
	}

	if strings.Contains(newString, `.`) {
		newString = strings.Replace(newString, ".", `\.`, -1)
	}

	if strings.Contains(newString, "?") {
		newString = strings.Replace(newString, `?`, `\?`, -1)
	}

	if strings.Contains(newString, "&") {
		newString = strings.Replace(newString, `&`, `\&`, -1)
	}

	cmdline := "sed" + " " + "-i" + " " + `s%\(^\s*` + oldString + `\s*` + equalitySigned + `\s*\).*%\1` + newString + `%g` + " " + file
	cmdlines := strings.Split(cmdline, " ")
	cmd := exec.Command(cmdlines[0], cmdlines[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if cmd.Run() != nil {
		log.Fatal(cmd.Run())
		return cmd.Run()
	} else {
		return nil
	}
}
*/

// 更新配置文件。
func updateConfiguration(path Config, oldString string, newString string, equalitySigned string) error {

	var updated []string
	lines := path.ReadConfFileByLine()
	re := regexp.MustCompile(`^\s*`+ oldString +`\s*` + equalitySigned + `\s*`)

	for _, line := range lines {
		if re.MatchString(line) {
			line = re.FindString(line) + strings.TrimSpace(newString)
			updated = append(updated, line)
		} else {
			updated = append(updated, line)
		}
	}
	
	file, err := os.OpenFile(path.Name, os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("文件打开失败：" + err.Error())
		return err
	}

	file.Truncate(0)
	file.Seek(0, 0)

    if strings.HasSuffix(updated[len(updated)-1], `\n`) {
        file.WriteString(strings.Join(updated, "\n"))
    } else {
        file.WriteString(strings.Join(updated, "\n") + "\n")
    }


	defer file.Close();

	return nil
}
