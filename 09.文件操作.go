package main

import (
	"fmt"
	"os"
)

func main()  {
	//os.Mkdir("warcello",0777)
	//os.MkdirAll("warcello/01/02", 0777)
	//err := os.Remove("warcello")
	//if err != nil {
	//	fmt.Println(err)
	//	os.RemoveAll("warcello/01/02")
	//}
	//写文件
	file := "warcello/warcello.txt"
	create, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
	}
	defer create.Close()
	for i := 0; i<10;i++ {
		create.WriteString("Just a test! \r\n")
		create.Write([]byte("Just a byte test !\r\n"))
	}

	//读文件
	//file := "warcello/warcello.txt"
	//re, err := os.Open(file)
	//if err != nil {
	//	fmt.Println(file,err)
	//}
	//defer re.Close()
	//buf := make([]byte, 1024)
	//for  {
	//	n,_ := re.Read(buf)
	//	if 0 == n {
	//		break
	//	}
	//	os.Stdout.Write(buf[:n])
	//}
	//os.Remove(file)

}