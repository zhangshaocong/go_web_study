package main

import (
	"fmt"
	"math"
)

//定义长方形的长跟高struct
type Rectangle struct {
	width,height float64
}
type Circle struct {
	radius float64
}
//计算长方形面积
func (r Rectangle) area() float64  {
	return r.width * r.height
}
//计算圆的面积
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//
//func area(r Rectangle) float64 {
//	return r.width * r.height
//}

func main()  {
	r1 := Rectangle{12,2}
	r2 := Rectangle{9,4}
	fmt.Println("Area of r1 is", r1.area())
	fmt.Println("Area of r2 is", r2.area())
	c1 := Circle{10}
	c2 := Circle{12}
	fmt.Println("Area of c1 is", c1.area())
	fmt.Println("Area of c2 is", c2.area())

}