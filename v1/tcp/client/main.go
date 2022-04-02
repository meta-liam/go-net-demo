package main

import (
	"flag"
	"fmt"
	"net"
)

// go run v1/tcp/client/main.go
// go build -ldflags '-w -s' -o build/tcp/client_main ./v1/tcp/client/main.go

func main() {
	client()
}

const network = "tcp4"

var host = flag.String("host", "localhost", "host") //localhost , 10.10.100.177
var port = flag.String("port", "20001", "port")

//const address = "10.10.100.177:20001" //"127.0.0.1:20001" //"127.0.0.1:8190"

func client() {
	address := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Println("run client: Connecting to ", address)
	tcpAddr, _ := net.ResolveTCPAddr(network, address)
	socket, _ := net.DialTCP(network, nil, tcpAddr)
	defer func(socket *net.TCPConn) {
		_ = socket.Close()
	}(socket)
	var input string
	fmt.Println("input for 5 loops")
	for i := 0; i < 5; i++ {
		_, _ = fmt.Scanf("%s", &input)
		fmt.Println("input:", input)
		_, _ = socket.Write([]byte(input))
		response := make([]byte, 1024)
		readLen, _ := socket.Read(response)
		fmt.Println(string(response[:readLen]))
	}
}
