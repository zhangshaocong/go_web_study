package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main()  {
	service := ":7000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
	checkError1(err)
	listener, err := net.ListenTCP("tcp",tcpAddr)
	checkError1(err)
	for {
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}
func handleClient(conn net.Conn)  {
	//设置2分钟连接超时
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte,128)
	defer conn.Close()
	for  {
		read_len , err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			break
		}  else {
			conn.Write([]byte(request[:read_len]))
		}
		request = make([]byte,128) //清除最后读取的内容
	}

}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}