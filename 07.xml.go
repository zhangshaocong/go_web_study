package main

import (
	"encoding/xml"
	"fmt"
	"os"
)
type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:",innerxml"`
}
//
//type server struct {
//	XMLName xml.Name `xml:"server"`
//	ServerName string `xml:"serverName"`
//	ServerIP string `xml:"serverIP"`
//}

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
}
type server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}
func main()  {
	//file, err := os.Open("07server.xml")
	//if err != nil {
	//	fmt.Print("errror %v \n", err)
	//	return
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Printf("error : %v", err)
	//	return
	//}
	//v := Recurlyservers{}
	//err = xml.Unmarshal(data,&v)
	//if err != nil {
	//	fmt.Printf("error :%v", err)
	//	return
	//}
	//fmt.Println(v)

	//生成xml
	x := &Servers{Version: "1"}
	x.Svs = append(x.Svs, server{ServerName: "Beijing", ServerIP: "127.0.0.3"})
	x.Svs = append(x.Svs, server{"nanjing","127.0.0.4"})
	output, err := xml.MarshalIndent(x," ", " ")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
