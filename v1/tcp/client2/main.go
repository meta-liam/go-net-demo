package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

// go run v1/tcp/client2/main.go
// go build -ldflags '-w -s' -o build/tcp/client2_main ./v1/tcp/client2/main.go

/*
功能：写完以后再收数据
*/
var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "20001", "port")

func main() {
	address := fmt.Sprintf("%s:%s", *host, *port)
	flag.Parse()
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to " + address)
	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}
func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		db := "hello " + strconv.Itoa(i) + "\r\n"
		fmt.Print("[write]", db)
		_, e := conn.Write([]byte(db))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}
}
func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for i := 1; i <= 10; i++ {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		fmt.Print("[read]", line)
	}
}
