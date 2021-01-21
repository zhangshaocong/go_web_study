package main

import (
	"fmt"
	"strconv"
	"strings"
)
func checkErr(e error){
	if e != nil{
		fmt.Println(e)
	}
}
func main() {
	//字符串是否包含
	fmt.Println(strings.Contains("seasdfasd","asd"))
	fmt.Println(strings.Contains("seasdfasd","iop"))
	fmt.Println(strings.Contains("seasdfasd",""))
	fmt.Println(strings.Contains("",""))

	//字符串链接
	slic := []string{"foo","button","lili"}
	fmt.Println(strings.Join(slic,":"))

	//查找字符串首次出现的位置
	fmt.Println(strings.Index("checken","ken"))
	fmt.Println(strings.Index("checken","yes"))

	//字符串重复
	fmt.Println("ba" + strings.Repeat("ba",20))

	//字符串替换
	fmt.Println(strings.Replace("hello world","world","warcello",-1))

	//分割
	fmt.Printf("%q \n", strings.Split("a,b,c,d",","))

	//去除空格
	fmt.Println(strings.Trim(" !!! Actung !!! ","! "))

	//去除指定字符的空格，并返回slice
	fmt.Println(strings.Fields(" foo    bar   hello     "))

	str := make([]byte,0,100)
	str = strconv.AppendInt(str,4567,10)
	str = strconv.AppendBool(str,false)
	str = strconv.AppendQuote(str,"abcdefg")
	str = strconv.AppendQuoteRune(str,'单')
	fmt.Println(str)
	fmt.Println(string(str))

	a, err := strconv.ParseBool("false")
	checkErr(err)
	fmt.Println(a)
	b, err := strconv.ParseFloat("123.23",64)
	checkErr(err)
	fmt.Println(b)
	c, err := strconv.ParseInt("1234",10,64)
	checkErr(err)
	fmt.Println(c)
	d, err := strconv.ParseUint("12345", 10, 64)
	checkErr(err)
	fmt.Println(d)
	e, err := strconv.Atoi("1023")
	checkErr(err)
	fmt.Println(e)

}
