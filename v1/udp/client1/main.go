package main

import (
	"fmt"
	"net"
	"time"
)

// go run go run v1/tcp/udp/client1/main.go

func main() {
	//addrs, _ := net.InterfaceAddrs()
	//fmt.Println("my addrs:", addrs)
	client()
}

const multicastAddr = "224.0.1.77" // "127.0.0.1" //
const clientPort = 20000           // 20001

var input string = "10.10.100.177"

func client() {
	time.Sleep(5 * time.Second)
	address := fmt.Sprintf("%s:%d", multicastAddr, clientPort)
	fmt.Println("[udp]:address:", address)
	//tcpAddr, _ := net.ResolveTCPAddr("udp4", address)
	conn, err := net.Dial("udp4", address)
	if err == nil {
		//conn.Close()
		fmt.Println("[udp]:Connection successful")
	} else {
		fmt.Println("[udp]:error:", err)
	}

	fmt.Println("------- write")
	//_, _ = fmt.Scanf("%s", &input)
	db := []byte(input)
	fmt.Println("udp:write 1:", input)
	fmt.Println("udp:write 2:", db)
	fmt.Println("udp:write 3:", string(db))
	_, _ = conn.Write(db)
	response := make([]byte, 1024)
	readLen, _ := conn.Read(response)
	fmt.Println("udp:read:::" + string(response[:readLen]))
	//_, _ = fmt.Scanf("%s", &input)
	time.Sleep(5 * time.Second)
	fmt.Println("-------close")
	conn.Close()
}

// go run v1/tcp/udp/client1/main.go
