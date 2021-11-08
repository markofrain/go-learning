package main

type Integer int

func (a Integer) Less(b Integer) bool { // 面向对象
	return a < b
}
func IntegerLess(a Integer, b Integer) bool { //面向过程
	return a < b
}

/**
为Integer类型增加了Add方法，由于需要修改对象的值，所以需要用指针引用
*/
func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {

	// # 为类型添加方法

	//var a Integer = 1
	//if a.Less(2) {
	//	fmt.Println(a, "Less 2")
	//}

	//a.Less(2)
	//IntegerLess(a,2)

	// 面向对象只是换了一种语法形式来表达。C++语言的面向对象之所以让有些人迷惑的一大原因就在于其隐藏的this指针
	// python中self参数，他和this指针的作用是完全一样的

	// this到底从何而来。这主要是因为Integer类的Less()方法隐藏了第一个参数Integer* this

	// 如果要求的对象必须以指针传递，有时是一个额外成本，因为对象有时很小，用指针传递并不划算
	// 只有在你需要修改的对象的时候，才必须要指针。它不是Go语言的约束，而是自然约束
	// 给类型Integer定义了Add方法
	// a.Add(2)
	// fmt.Println(a)
	// 运行此代码则输出的是3,。如果没有指针的话，输出的依然是1。这是因为Go语言和C语言一样，类型是基于值传递的。要想改变变量的值只能传递指针

	// # 值语义和引用语义

	// 值语义和引用语义在于赋值
	// b = a
	// b.Modify()
	// 如果b的修改不会影响a的值，那么此类型属于值类型。如果会影响a的值，那么此类型是引用类型
	// Go语言中的大多数类型都基于值语义，包括：
	// 基本类型，如byte、int、bool、float32、float64和string等；
	// 复合类型，如数组（array）、结构体（struct）和指针（pointer）等。
	// 在Go中，数组的传递是按值语义传递(表现在为结构体赋值的时候，该数组会被完整地复制)
	//var a = [3]int{1, 2, 3}
	//var b = a
	//b[1]++
	//fmt.Println(a, b)
	// 要想表达引用，需要用指针
	//var a = [3]int{1, 2, 3}
	//var b = &a
	//b[1]++
	//fmt.Println(a, *b)
	// 这表明b=&a赋值语句是数组内容的引用。变量b的类型不是[3]int，而是*[3]int类型

	// Go语言中有4个类型比较特别，看起来像引用类型
	// 数组切片：指向数组（array）的一个区间。
	// map：极其常见的数据结构，提供键值查询能力。
	// channel：执行体（goroutine）间的通信设施。
	// 接口（interface）：对一组满足某个契约的类型的抽象。

	// # 结构体
	// 所有的Go语言类型（指针类型除外）都可以有自己的方法。在这个背景下，Go语言的结构体只是很普通的复合类型
	//type Rect struct {
	//	x, y float64
	//	width, height float64
	//}
	//func (r *Rect) Area() float64 {
	//	return r.width * r.height
	//}

}
