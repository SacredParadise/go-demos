package main

import (
"fmt"
)

type Rect struct {                               //定义一个结构体
	width  float64
	length float64
}

func (rect Rect) area() float64 {               //定义一个方法，按值传递
	return rect.width * rect.length
}

func (rect *Rect) area1() float64 {            //定义一个方法，按指针传递
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

func main() {
	var rect = new(Rect)     //使用new函数创建一个结构体指针rect，也就是说rect的类型是*Rect
	rect.width = 100
	rect.length = 200
	fmt.Println("Width:", rect.width, "Length:", rect.length,"Area:", rect.area())  //通过结构体指针类型的变量调用area()方法
	fmt.Println("Width:", rect.width, "Length:", rect.length,"Area:", rect.area1())
}