package main

import (
	"fmt"
	"reflect"
)

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func main() {
	// 初始化

	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	// 在Go语言中，未进行显式初始化的变量都会被初始化为该类型的零值，例如bool类型的零值为false，int类型的零值为0，string类型的零值为空字符串。
	// 在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名，表示“构造函数”
	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
	fmt.Println(rect4)
	fmt.Println(reflect.TypeOf(rect1))
}
