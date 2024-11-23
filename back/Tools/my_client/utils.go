package myclient

import (
	"fmt"
	"net"
)

func GetAllLocalIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Println("IPv4:", ipNet.IP.String())
			} else {
				fmt.Println("IPv6:", ipNet.IP.String())
			}
		}
	}
}
