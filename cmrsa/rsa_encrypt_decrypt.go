package cmrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

//  PublicKeyEncrypt RSA 公钥加密方法，第一个参数为需要加密的字符串，第二个参数为 RSA 公钥字符串。加密后返回一个 Base64 编码的字符串及错误。
func PublicKeyEncrypt(text string, publicKey string) (string, error) {
	var publicKeyBytes []byte = []byte(publicKey)

	block, _ := pem.Decode(publicKeyBytes)

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		panic(err)
	}

	// 类型断言
	pubKey := publicKeyInterface.(*rsa.PublicKey)

	// 对明文进行加密
	encryptText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(text))
	if err != nil {
		panic(err)
	}

	// 返回base64编码字符串
	return base64.StdEncoding.EncodeToString(encryptText), nil
}

// PrivateKeyDecrypt 私钥解密方法，第一个参数为base64编码的加密字符串，第二个参数为 RSA 私钥字符串。解密后返回原始字符串。
func PrivateKeyDecrypt(encryptText string, privateKey string) (string, error) {
	decryptTextBytes, err := base64.StdEncoding.DecodeString(encryptText)

	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode([]byte(privateKey))

	//X509解码
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	//对密文进行解密
	decryptText, _ := rsa.DecryptPKCS1v15(rand.Reader, priKey, decryptTextBytes)

	//返回明文
	return string(decryptText), nil
}
