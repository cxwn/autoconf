package main

import (
	_ "autoconf/cli"
	"autoconf/cmrsa"
    _ "autoconf/curd"
    "autoconf/oasr"
	"fmt"
	_"os"
)

// 384 位的公钥长度，可加密包含数字、字母、符号的最长21位的密码，尝过该长度程序会报错，一般情况下256位加密即可满足要求，加密位数过高会导致效率低下。
var privateKey string = `-----BEGIN RSA PRIVATE KEY-----
MIHxAgEAAjEAw+3qUW6PYKDuF478HYzC7fV0at5NRSjHoLKEPqPNiBlg3QxV42bk
crVu+qZWJfORAgMBAAECMGuWMivjDQIffH4dOt2zFLr8JKAmT8HhQL8PW4Nw4dS6
b7rsv4OiRHDE2AjDP0Lj0QIZANZS6vwWFvySDLV8tf0WX/8yTlYcAxM7lQIZAOoH
UFbNq5QZRQy0F+P385hYRq6whlj5DQIYaRjGBBGs+fOAaeqar12ue0ym31DjLSY1
AhgWEYJ97P8VBB0CyajHEoaiAxEHQBYrJbECGDzN9H7nRPr+2naCeLXW5NKs49vo
LdRirA==
-----END RSA PRIVATE KEY-----`

var publicKey string = cmrsa.GetPublicKeyFromPrivateKey(privateKey)

func main() {

	/*===================
	encryptText, _ := cmrsa.PublicKeyEncrypt(`{bPAMgE$}j0Z8QFz`, publicKey)
	fmt.Print(encryptText)
	decryptText, _ := cmrsa.PrivateKeyDecrypt(encryptText, privateKey)
	print("\n")
	fmt.Print(decryptText)
	===================*/

	fmt.Println("若配置中涉及加密字段，请使用 RSA 算法进行加密。该版本的RSA加密公钥为：")
	fmt.Println(publicKey)

	// 离线转写自动化配置。
	oasr.ConfOASR(privateKey, publicKey)

	/*==================================cli===========================
	var con curd.Config
	cli.InitCli(&con)
	if con.Name == "" {
		fmt.Println("配置文件不能为空，程序即将退出，请指定配置文件再试！")
		os.Exit(188)
	}
	if con.Prefix == "" {
		fmt.Println("环境变量前缀不能为空，程序即将退出，请指定正确的环境变量前缀再试！")
		os.Exit(186)
	}

	fmt.Println("正在对系统进行自动化配置，请稍候...")
	if con.Update(privateKey) {
		fmt.Println("系统自动化配置完成，谢谢！")
	} else {
		fmt.Println("系统自动化配置失败，请检查！")
	}
	====================================cli===========================*/

}
