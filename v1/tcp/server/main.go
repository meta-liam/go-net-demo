package main

// go run v1/tcp/server/main.go
// go build -ldflags '-w -s' -o build/tcp/server_main v1/tcp/server/main.go
// 参考：https://colobu.com/2014/12/02/go-socket-programming-TCP/

/**
功能：tcp 全网监听
*/
import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	Server()
}

var host = flag.String("host", "", "host") // 允许所有ip交换数据
var port = flag.String("port", "20001", "port")

const network = "tcp4"

//const address = ":20001" //":20001" 允许所以ip交换数据

func Server() {
	address := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Println("[tcp]run server,Listening on " + address)
	tcpAddr, err := net.ResolveTCPAddr(network, address)
	if err != nil {
		fmt.Println("[tcp.error1]", err)
		os.Exit(1)
	}
	// listener对应ServerSocket
	serverSocket, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		fmt.Println("[tcp.error2]", err)
		os.Exit(1)
	}
	for {
		// 每次连接建立返回一个Connection，Connection对应Socket
		conn, err := serverSocket.AcceptTCP()
		//fmt.Println("[tcp]connection established...")
		fmt.Printf("[tcp]Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		if err != nil {
			fmt.Println("[tcp.error3]", err)
			os.Exit(1)
		}
		// 开辟Goroutine去处理新的连接
		go server(conn)
	}
}

// 最普通的版本
func server(socket *net.TCPConn) {
	defer func(tcpConn *net.TCPConn) {
		err := tcpConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(socket)
	for {
		request := make([]byte, 1024)
		readLen, err := socket.Read(request)
		if err == io.EOF {
			fmt.Println("[tcp]连接关闭")
			return
		}
		msg := string(request[:readLen])
		fmt.Println("[tcp]get msg::", msg)
		msg = "echo:" + msg
		_, _ = socket.Write([]byte(msg))
	}
}
