package main

import "fmt"

type testInt func(int) bool  //声明一个函数类型

func isOdd(integer int) bool {
	if integer % 2 == 0 {
		return false
	}
	return true
}

func isEvent(integer int) bool {
	if integer % 2 == 0 {
		return true
	}
	return false
}

//这里的f,把testInt当成了一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice{
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main()  {
	slice := []int{1,2,3,4,5,7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) //函数当做值来传递了
	fmt.Println("Odd elements of slice are :", odd)
	even := filter(slice, isEvent)
	fmt.Println("Event elements of slice are: ", even)
}