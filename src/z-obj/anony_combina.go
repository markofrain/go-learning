package main

import (
	"fmt"
	"strconv"
)

// 匿名组合
// 确切地说，Go语言也提供了继承，但是采用了组合的文法，所以我们将其称为匿名组合
// 例如：若子类中含有某父类成员，则该子类会自动继承其方法，可以直接点出来匿名成员的方法

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println(base.Name)
}
func (base *Base) Bar() {
	fmt.Println(base.Name + "a")
}

type Foo struct {
	Base
	Age int
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
	fmt.Println(foo.Name + strconv.FormatInt(int64(foo.Age), 10))
}

func main() {

	// 以上代码定义了一个Base类（实现了Foo()和Bar()两个成员方法），然后定义了一个
	//Foo类，该类从Base类“继承”并改写了Bar()方法（该方法实现时先调用了基类的Bar()方法）。

	// 在“派生类”Foo没有改写“基类”Base的成员方法时，相应的方法就被“继承”，

	t1 := new(Base)
	t1.Name = "chen"
	t1.Bar()
	t1.Foo()
	t2 := new(Foo)
	t2.Age = 18
	t2.Base = *t1
	t2.Bar()
	t2.Foo()

	// 另外，在Go语言中，你还可以以指针方式从一个类型“派生”：
	//type Foo struct {
	//	*Base
	//	...
	//}
	// 这段Go代码仍然有“派生”的效果，只是Foo创建实例的时候，需要外部提供一个Base类实例的指针

	// 注意

	//type X struct {
	//	Name string
	//}
	//type Y struct {
	//	X
	//	Name string
	//}
	//组合的类型和被组合的类型都包含一个Name成员，会不会有问题呢？答案是否定的。
	//所有的Y类型的Name成员的访问都只会访问到最外层的那个Name变量，X.Name变量相当于被隐藏起来了

	//type Logger struct {
	//	Level int
	//}
	//type Y struct {
	//	*Logger
	//	Name string
	//	*log.Logger
	//}
	// 匿名组合类型相当于以其类型名称（去掉包名部分）
	// 作为成员变量的名字。按此规则，Y类型中就相当于存在两个名为Logger的成员，虽然类型不同。
	// 因此，我们预期会收到编译错误
}
