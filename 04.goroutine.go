package main

import (
	"fmt"
	"runtime"
)

func say(s string)  {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main()  {
	go say("wold") //开启一个新的goroutines执行
	say("hello") //当前goroutines执行
}