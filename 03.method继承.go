package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

func (h Human) SayHi()  {
	fmt.Printf("Hi I am %s you can call me on %s \n", h.name,h.phone)
}
//实现Sing方法
func (h Human) Sing(lyrics string)  {
	fmt.Println("La la,la la la, la la", lyrics)
}
//实现喝酒的方法
func (h Human) Guzzle(beerStein string)  {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}
//重写匿名字段method方法
func (e Employee) SayHi()  {
	fmt.Printf("Hi I am %s ,I work at %s,you can call me on %s \n", e.name,e.company,e.phone)
}
//实现借钱method
func (s *Student) BorrowMoney(amount float32)  {
	s.loan += amount
}
//实现发工资method
func (e *Employee) SpendSalary(amount float32)  {
	e.money -= amount
}
// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string) //唱歌
	Guzzle(beerStein string) //喝酒
}
type YoungChap interface {
	SayHi()
	Sing(lyrics string) //唱歌
	BorrowMoney(amount float32) //借钱
}
type ElderGent interface {
	SayHi()
	Sing(lyrics string) //唱歌
	SpendSalary(amount float32)
}


func main()  {
	mark := Student{Human{"Warcello",26,"18611223344"},"Dazhuan",0.00}
	//paul := Student{Human{"Paul",15,"18611223311"},"Gaozhongf",100}
	//sam := Employee{Human{"Mac",12,"19011223344"},"Home",1000}
	//tom := Employee{Human{"Mac",28,"19011226644"},"Home",5000}

	var i Men
	//i能存储student
	i = mark
	fmt.Println("this is mark, astudent:")
	i.SayHi()
	i.Sing("Novemer rain ")
}
