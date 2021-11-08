package main

import (
	"fmt"
	"math"
)

func main() {

	//布尔类型:bool
	//整型: int8,int16,int32,int64
	//浮点类型: float32,float64
	//复数类型: complex64,complex128
	//字符串: string
	//字符类型: rune
	//错误类型: error
	// 复合类型
	//指针类型: pointer
	//数组: array
	//切片: slice
	//字典: map
	//通道: chan
	//结构体: struct
	//接口: interface

	// int ,uint uintptr等，这些类型不需要使用者对类型指明长度，一般用int和uint就行了

	// # 布尔类型

	//b1 := true
	//b2 := (1==2)
	//fmt.Printf("%t,%t",b1,b2)
	// 不能使用bool(1)直接转为bool类型，将会编译错误

	// # 浮点数比较
	//v := IsEquals(0.340001,0.34,0.00000001)
	//fmt.Println(v)

	// # 复数类型

	//var value1 complex64
	//value1 = 3.2 + 12i
	//value2 := 3.2 + 12i
	//value3 := complex(3.2,12)
	//fmt.Println(value1)
	//fmt.Println(value2)
	//fmt.Println(value3)
	//fmt.Printf("x:%f,%f",real(value3),imag(value3)) // real获取x部分数值，imag获取y部分数值

	// # 字符串
	// 在go中，字符串也是基本数据类型
	//var str string
	//str = "Hello World"
	//ch := str[0]
	//fmt.Printf("The length of \"%s\" is %d \n",str,len(str))
	//fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
	// 字符串中的内容可以以下表的形式获取，但是在初始化之后便不能通过数组下标的方式再进行赋值

	// # 数组
	//arr := [6]int{1,2,3,4,5,6}
	//fmt.Println(arr)
	//var arrSlice []int = arr[1:] //切片
	//fmt.Println(arrSlice)
	//mySlice1 := make([]int,5) //创建一个初始元素个数为5额数组切片，元素初始值为0
	//fmt.Println(mySlice1)
	//mySlice2 := make([]int,5,10) //创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	//fmt.Println(mySlice2)
	//myslice3 := []int{1,2,3,4,5} // 直接创建并初始化5个元素的数组切片
	//fmt.Println(myslice3)

	//// 遍历
	//for i := 0; i <len(myslice3); i++ {
	//	fmt.Println("mySlice[", i, "] =", myslice3[i])
	//}
	//for i,v := range myslice3 {
	//	fmt.Println("mySlice[", i, "] =", v)
	//}

	//// 动态增减元素
	////类似java中arrayList的实际元素大小和容量大小,可以进行动态扩容，
	//// 而make([]int,5,10)就定义了5个初始元素和长度为10的存储空间
	//// cap()内置函数和len()内置函数可以直接获取数组切片分配大小和实际元素大小
	//fmt.Printf("%d,%d \n",cap(mySlice2),len(mySlice2))
	//mySlice2 = append(mySlice2,1,2,3)	// append内置函数可以追加多个元素到数组切片中
	//fmt.Println(mySlice2)
	//mySlice2 = append(mySlice2,myslice3...)		// 也可以添加一个数组切片作为参数，但因为第二个参数以后为多参数，因此需要...进行拆分为多参
	//fmt.Println(mySlice2)

	// 基于数组切片创建数组切片，选择的oldSlice元素范围可以超过所包含的个数，只要不超过存储能力cap范围即可，其他多余的数值自动填上0

	// copy函数内容拷贝
	//slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := []int{5, 4, 3}
	//fmt.Println(slice1)
	//fmt.Println(slice2)
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//fmt.Println(slice1)
	//fmt.Println(slice2)
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	//fmt.Println(slice1)
	//fmt.Println(slice2)

	// # map
	// 示例
	//var personDB map[string] PersonInfo // 变量声明，personDB是变量名,string是键类型,PersonInfo是值类型
	//personDB = make(map[string] PersonInfo)	// 变量初始化,使用make创建map，第二个参数可以代表存储能力cap,make(map[string] PersonInfo,100)标识有100的存储能力
	//// 往这个map里插入几条数据
	//personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	//personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	//// 从这个map查找键为"1234"的信息
	//person, ok := personDB["1234"]
	//// ok是一个返回的bool型，返回true表示找到了对应的数据
	//if ok {
	//	fmt.Println("Found person", person.Name, "with ID 1234.")
	//} else {
	//	fmt.Println("Did not find person with ID 1234.")
	//}
	// 示例
	var test = map[string]PersonInfo{ // 创建并初始化
		"123": {"1", "2", "3"},
	}
	fmt.Println(test)
	value, ok := test["123"] //查询元素
	if ok {
		fmt.Println("find ", value)
	}
	delete(test, "123") // 删除元素
	fmt.Println(test)

}

/**
两个浮点数进行比较，p参数表示自定义精度
这个函数比较第一个参数和第二个参数的差值，然后和0比较大小，我们这里需要精度是两个参数的差的绝对值，不关注正负
*/
func IsEquals(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}
