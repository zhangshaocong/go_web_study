package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:5200")

	tcpListener, _ := net.ListenTCP("tcp",tcpAddr)
	defer tcpListener.Close()

	fmt.Println("Server ready to read...")

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("accept error :", err)
			continue
		}
		fmt.Println("A client connected: " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)

	}
}
func tcpPipe(conn *net.TCPConn)  {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println(" Disconnected : " + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(message))
		//time.Sleep(time.Second * 3)
		//msg := time.Now().String() + conn.RemoteAddr().String() + " server say hello! \n"
		b := []byte(string(message))
		conn.Write(b)
	}
}
