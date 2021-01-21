package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP string
}
type Serverslice struct {
	Servers []Server
}
func main()  {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	//生成json
	var j Serverslice
	j.Servers = append(j.Servers, Server{ServerName: "VPN1", ServerIP: "123.12.123.23"})
	j.Servers = append(j.Servers, Server{ServerName: "VPN2", ServerIP: "123.12.123.24"})
	j.Servers = append(j.Servers, Server{ServerName: "VPN3", ServerIP: "123.12.123.25"})
	b, err:= json.Marshal(s)
	if err != nil {
		fmt.Println("error :%v \n", err)
	}
	fmt.Println(string(b))


}