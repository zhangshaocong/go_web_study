package main

import "fmt"

type person struct {
	name string
	age int
}
//比较两人年龄，返回年龄大的那个人，并返回年龄差
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1 , p1.age-p2.age
	}
	return p2, p2.age-p1.age
}
func main()  {
	//方式一
	var tom person
	tom.name,tom.age = "tom", 18
	//方式二：根据key赋值
	bob := person{age: 25,name: "bob"}
	//方式三：按照顺序初始化
	paul := person{"paul",43}

	tb_Older, tb_diff := Older(tom, bob)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	tp_Older, tp_diff := Older(tom,paul)

	fmt.Printf("Of %s and %s , %s is older by %d years \n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	bp_Older, bp_diff := Older(bob,paul)
	fmt.Printf("Of %s and %s , %s is older by %d years \n",
		bob.name, paul.name, bp_Older.name, bp_diff)


}