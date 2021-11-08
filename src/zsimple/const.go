package main

import "fmt"

func main() {
	// # 字面量
	// 硬编码的字面量,只要这个常量在相应的类型的值域范围内，就可以作为该类型的常量
	//-1
	//3.1415936	浮点形式
	//3.2+12i	复数形式
	//true		布尔形式
	// "foo"	字符串常量

	// # 常量定义
	/*
		const Pi float64 = 3.1415926
		const zore = 0.0	// 无类型浮点常量
		const (
			size int64 = 1024
			eof = -1	//无类型整型常量
		)
		const u,v float64 = 0,3 //常量的多重赋值，u=0.0, v=3.0
		const a,b,c = 1,2,"a"	// a=1,b=2,c="a",无类型整型和字符串常量

		const mask = 1<<2	// 可以是一个在编译期运算的常量表达式，但不能是运行期才知道的返回结果
	*/

	// # 预定义常量
	// go中预定义了这些常量 true,false,iota
	// iota可以认为是一个可以被编译器修改的常量，在每个const关键字出现时被重置为0，然后在下个const出现之前，每出现一次iota，其所代表的数字就回会自动增加1

	//const (
	//	c0 = iota
	//	c1 = iota
	//	c2 = iota
	//)
	//fmt.Printf("c0:%d0,c1:%d,c2:%d",c0,c1,c2)

	//const (
	//	a = 1 << iota
	//	b = 1 << iota
	//	c = 1 << iota
	//)

	//const c0 = iota		// 0
	//const c1 = iota		// 0
	//fmt.Printf("%d,%d",c0,c1)

	//const (
	//	c0 = iota	// 0
	//	c1			// 1
	//	c2			// 2
	//)
	//fmt.Printf("%d,%d,%d",c0,c1,c2)

	// # 枚举
	// 枚举是指一系列相关的常量，不支持enum枚举关键字
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出,小写开头包外不可见
	)
	fmt.Printf("%d,%d", Sunday, Monday)

}
