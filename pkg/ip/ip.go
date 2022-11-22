package ip

import (
	"fmt"
	"net"
)

// GetNetworkCardIp 根据网卡名字返回ip地址
func GetNetworkCardIp(networkCard string) (string, error) {
	netInterface, err := net.InterfaceByName(networkCard)
	if err != nil {
		return "", err
	}
	if netInterface.Flags&net.FlagUp == 0 {
		return "", fmt.Errorf("net error")
	}
	addrs, err := netInterface.Addrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}
