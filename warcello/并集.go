package main

import (
	"fmt"
	"strings"
)

func union(slice1,slice2 []string) []string {
	m := make(map[string]int)
	for _, v:= range slice1 {
		m[v]++
	}
	for _, v:= range slice2 {
		times,_:= m[v]
		if times == 0 {
			slice1 = append(slice1,v)
		}
	}
	return slice1
}
func intersect(slice1, slice2 []string) []string  {
	m := make(map[string]int)
	n := make([]string,0)
	for _,v := range slice1 {
		m[v]++
	}
	fmt.Println("m:::",m)

	for _,v := range slice2 {
		fmt.Println("v:::",v)
		times,_:= m[v]
		fmt.Println("times:",times)
		if times == 1 {
			fmt.Println(n)
			n = append(n,v)
		}
	}
	return n

}
func matchNum(from,own string) string  {
	m := make(map[string]int)
	n := make([]string,0)

	from_cov_slice1 := strings.Split(from,",")
	own_cov_slice2 := strings.Split(own,",")
	for _,v := range from_cov_slice1 {
		m[v]++
	}
	for _,v := range own_cov_slice2 {
		times,_:= m[v]
		if times == 0 {
			n = append(n,v)
		}
	}
	return strings.Join(n,",")
}

func other(from ,own string) string {
	own_slice := strings.Split(own,",")
	res := make([]string,0)
	for _,v :=range own_slice {
		if s := strings.Contains(from,string(v)); !s {
			res = append(res,v)
		}
	}
	return strings.Join(res,",")
}

func main()  {
	form:="2,4,55,70"
	sql :="2,4,6,9"
	ma := matchNum(form,sql)
	fmt.Println(ma)
	other := other(form,sql)
	fmt.Println(other)
}