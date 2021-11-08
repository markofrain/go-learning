package main

import (
	"errors"
	"fmt"
)

// 小写字母开头的函数只在本包内可见，大写字母开头的函数才
// 能被其他包使用。
// 这个规则也适用于类型和变量的可见性
func main() {
	c, err := Add(-1, 2)
	if err == nil {
		fmt.Println(c)
	} else {
		fmt.Println("err", err)
	}

	//SomeInter("asd",1,2,"23")

	// 匿名函数可以被声明，并直接作为方法调用
	b := func(a int) {
		fmt.Println(a)
	}
	b(1)
	// 匿名函数何以定义返回函数,但必须在右大括号后增加小括号，然后再进行调用
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()

}

/**
同类型参数可只写一个类型。返回多个返回值需要使用括号
*/
func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil
}

/**
不定参数
*/
func SomeParam(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}

	// 不定参数传递,按原样传递
	SomeParam(args...)
	// 传递片段
	SomeParam(args[1:]...)
}

/**
任意类型的不定参数
*/
func SomeInter(format string, args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		case interface{}:
			fmt.Println(arg, "is a interface{}")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
