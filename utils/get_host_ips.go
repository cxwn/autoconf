package utils

import (
	"fmt"
	"log"
	"net"
	"os"
)

// GetHostIP 增加主机ip检测单元，返回主机包含主机IP的数组，每台主机的网卡可能配置有多个 IP.
func GetHostIP() []string {
	netInterfaces, err := net.Interfaces()
	ips := make([]string, 0)
	var ipv4 string
	if err != nil {
		fmt.Println("获取网卡信息失败，未成功获取本机IP，程序将退出，错误信息：", err.Error())
		os.Exit(169) //获取IP发生错误的错误代码：错误代码169
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addr, _ := netInterfaces[i].Addrs()
			for _, address := range addr {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ipv4 = ipnet.IP.String()
						log.Printf("请核对，主机或 Pod IP 为：%s\n", ipv4)
						ips = append(ips, ipv4)
					}
				}
			}
		}
	}
	return ips
}
