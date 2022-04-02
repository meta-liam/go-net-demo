package main

import (
	"fmt"
	"net"
	"strings"
)

// go run v1/tcp/ip/main.go
// https://www.jianshu.com/p/301aabc06972

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func main() {
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
	fmt.Println("------------------------")
	getIP2()
}

func getIP2() {
	fmt.Println("------------------------2")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("address:")
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("IP:", ipnet.IP.String())
			}
		}
	}
}
