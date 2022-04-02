package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// go run v1/tcp/udp/server/main.go
var host = flag.String("host", "", "host")
var port = flag.String("port", "37", "port")

func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	msg := string(data[:n])
	fmt.Println("[udp]get msg::", msg)

	daytime := time.Now().Unix()
	fmt.Println("[link]", n, remoteAddr)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(daytime))
	conn.WriteToUDP(b, remoteAddr)
}
