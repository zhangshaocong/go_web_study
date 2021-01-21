package main

import (
	"fmt"
	"regexp"
)

func main() {
	//正则判断ip
	//ip := "112.113.11.2"
	//m, err := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",ip)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(m)

	//resp, err := http.Get("http://www.baidu.com")
	//if err != nil {
	//	fmt.Println("Http error")
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("read error")
	//	return
	//}
	////fmt.Println(string(body))
	//src := string(body)
	////将html 标签全部替换成小写
	//re,_ := regexp.Compile("\\<[\\S\\s]+?\\>")
	//src = re.ReplaceAllStringFunc(src,strings.ToLower)
	////fmt.Println(src)
	////去除script
	//re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	//src = re.ReplaceAllString(src,"")
	////去除style
	//re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	//src = re.ReplaceAllString(src, "")
	////去除所有尖括号内的html代码并换成换行符
	//re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	//src = re.ReplaceAllString(src,"")
	////去除连续的换行符
	//re, _ = regexp.Compile("\\s{2,}")
	//src = re.ReplaceAllString(src,"\n")
	//fmt.Println(src)

	//正则查找
	a := "I am learning go language"
	re, _ := regexp.Compile("[a-z]{2,4}")

	//查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:",string(one))

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，否则返回指定长度
	all := re.FindAll([]byte(a), -1)
	//for _,v := range all{
	//	fmt.Println(string(v))
	//}
	fmt.Println("FindAll", all)

	//查找符合条件的index位置，开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println(index)

}