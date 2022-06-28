package ipGet

import (
	"fmt"
	"net"
	"strings"
	// "os"
)

// 获取本地的ipv6地址
func GetLocalIPv6() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// 检测是否为公网IPV6地址
			// 移动前缀为2409
			// 联通为2408
			// 电信为240e
			if strings.Contains(ipnet.IP.String(), "2409") || strings.Contains(ipnet.IP.String(), "2408") || strings.Contains(ipnet.IP.String(), "240e") {
				return ipnet.IP.String()
			}
		}
	}
	return ""

}
