package oasr

import (
	"autoconf/curd"
	"autoconf/utils"
	"fmt"
	"log"
	"os"
)

// 离线转写系统自动化配置。

func ConfOASR(privateKey string, publicKey string) {

	ips := utils.GetHostIP()
	if len(ips) > 1 {
		log.Printf("系统检测到多个 IP 地址，默认使用第一个 IP 地址【%s】进行配置，请确认！\n", ips[0])
	} else if len(ips) < 1 {
		log.Printf("未成功获取系统 IP 地址，配置失败！\n ")
		os.Exit(202) // IP 配置失败，系统退出代码为：202
	} else if len(ips) == 1 {
		log.Printf("此系统配置有一个 IP 地址： %s\n", ips[0])
	}
	fmt.Println("若配置中涉及加密字段，请使用 RSA 算法进行加密。该版本的RSA加密公钥为：")
	fmt.Println(publicKey)

	inis := []curd.Config{
		{
			Name:           "/apache-tomcat-8.5.5/webapps/WEB-INF/classes/redis.properties",
			Comment:        "#",
			Prefix:         "JTR_",
			Connection:     ".",
			EqualitySigned: "=",
		},
		{
			Name:           "/apache-tomcat-8.5.5/webapps/WEB-INF/classes/jdbc.properties",
			Comment:        "#",
			Prefix:         "JTJ_",
			Connection:     "",
			EqualitySigned: "=",
		},
		{
			Name:           "/apache-tomcat-8.5.5/webapps/WEB-INF/classes/log4j.properties",
			Comment:        "#",
			Prefix:         "JTL_",
			Connection:     ".",
			EqualitySigned: "=",
		},
		{
			Name:           "/ClusterSceneServer/configure.cfg",
			Comment:        "#",
			Prefix:         "CLU_",
			Connection:     "",
			EqualitySigned: "=",
		},
		{
			Name:           "/Offline_System_Client/cfg.ini",
			Comment:        "#",
			Prefix:         "OFC_",
			Connection:     "",
			EqualitySigned: "=",
		},
		{
			Name:           "/DetectOverlap/in.cfg",
			Comment:        "#",
			Prefix:         "DTC_",
			Connection:     "",
			EqualitySigned: "=",
		},
		{
			Name:           "/TBNR_release_time/bin/Decode.cfg",
			Comment:        "#",
			Prefix:         "TRT",
			Connection:     "",
			EqualitySigned: "=",
		},
		{
			Name:           "/TBNR_release_time/model/scripts/WFSTDecoder-inputMethod_dnn_onlyrec.cfg",
			Comment:        "#",
			Prefix:         "WSF_",
			Connection:     "",
			EqualitySigned: "=",
		},
		{
			Name:           "/xml_server/configure_xml.cfg",
			Comment:        "#",
			Prefix:         "XMS_",
			Connection:     "",
			EqualitySigned: "=",
		},
	}

	
	fmt.Println("正在对离线转写系统进行自动化配置，请稍候...")
	for _, ini := range inis {
		if ini.Update(privateKey) {
			fmt.Printf("系统配置%s自动化配置完成，谢谢！\n", ini.Name)
		} else {
			fmt.Printf("系统配置%s自动化配置失败，请检查！\n", ini.Name)
		}
	}
}
