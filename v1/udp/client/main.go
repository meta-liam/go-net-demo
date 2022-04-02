package main

// go run v1/tcp/udp/client/main.go
// https://colobu.com/2014/12/02/go-socket-programming-UDP/
import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "37", "port")

//go run timeclient.go -host time.nist.gov
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("liam")) //write
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	data := make([]byte, 4)
	readLen, err := conn.Read(data) //read
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}
	fmt.Println("[get msg]", string(data[:readLen]))

	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())
	os.Exit(0)
}
