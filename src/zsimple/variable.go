package main

import "fmt"

func main() {

	// # 声明变量
	/*var v1 int
	var v2 string
	var v3 [10]int // 数组
	var v4 []int // 数组切片
	var v5 struct{
		f int
	}
	var v6 *int  //指针
	var v7 map[string]int //map,key为string类型，value为int类型
	var v8 func(a int) int  // 函数
	v1 = 12
	fmt.Println(v1)

	// 避免重复写var
	var (
		v9 int
		v10 string
	)*/

	// # 变量初始化

	/*var v1 int = 10
	var v2 =11
	v3 := 11 // 同时进行变量声明和初始化工作
	*/

	// # 变量赋值

	// 多重赋值
	var i, j int
	i, j = j, i

	// # 匿名变量
	// 如果函数返回多个值，但只想要其中一个，可以使用_忽略
	_, _, nickName := ret_some_value()
	fmt.Println(nickName)

}

/**
返回多参数
*/
func ret_some_value() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibo Marubo"
}
