package utils

import (
	"strconv"
	"strings"
)

type Subnet struct {
	SubInet string // 形如：192.168.12.0/24
}

// IsBelongToSubnet 判断某个 IP 是否属于指定子网
func IsBelongToSubnet(ip string, subnet Subnet) bool {
	decimalIpArray := make([]int, 4)   // IP 地址十进制表示
	binMaskArray := make([]string, 35) // 子网掩码，二进制字符串表示
	decimalMaskArray := make([]int, 4) // 子网掩码，十进制数字表示
	var sbits int                      // 子网位数
	dotCount := 0                      // IP 中点的个数
	ipAddress := strings.Split(ip, `.`)
	tempAddress := make([]string, 4)

	// 判断IP是否合法
	if len(ipAddress) != 4 {
		return false
	}

	for ikey := range ipAddress {
		decimalIpArray[ikey], _ = strconv.Atoi(ipAddress[ikey])
	}

	// 判断子网是否合法
	subnetInfo := strings.Split(subnet.SubInet, `/`)
	if len(subnetInfo) != 2 {
		return false
	} else {
		sbits, _ = strconv.Atoi(subnetInfo[1]) //子网位数
	}

	// 生成二进制的子网掩码
	for n := 0; n < 35; n++ {
		if (n+1)%9 == 0 {
			binMaskArray[n] = "."
		} else {
			binMaskArray[n] = "0"
		}
	}
	for m := 0; m < sbits+dotCount; m++ {
		if binMaskArray[m] == "." {
			dotCount++
		} else {
			binMaskArray[m] = "1"
		}
	}

	// 二进制子网掩码转换为十进制
	binMask := strings.Join(binMaskArray, "")
	for key, decimal := range strings.Split(binMask, ".") {
		decimalMaskArray[key] = binString2Dec(decimal)
	}

	for ipkey := range decimalIpArray {
		for dkey := range decimalMaskArray {
			if ipkey == dkey {
				tempAddress[ipkey] = strconv.Itoa(decimalIpArray[ipkey] & decimalMaskArray[ipkey])
			}
		}
	}

	return strings.Join(tempAddress, ".") == subnetInfo[0]
}

// binString2Dec 二进制字符串转十进制
func binString2Dec(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}
