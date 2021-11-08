package main

/**
接口
Go语言的接口并不是其他语言（C++、Java、C#等）中所提供的接口概念。与他们有些许不同

在Go语言出现之前，接口主要作为不同组件之间的契约存在。对契约的实现是强制的，你(必须声明你的确实现了该接口)。也就是明确要实现哪个具体接口。
这类接口我们称为侵入式接口。“侵入式”的主要表现在于实现类需要明确声明自己实现了某个接口

2. 非侵入式接口
在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口。
即假如有一个接口定义了多个方法，任意一个实现了这些方法的类均可认为实现了该接口，而不需要明确的使用类似Java的implement关键字来实现具体某个接口。

例:IReader，IWriter，ICloser接口分别定义了Read,Write,Close方法，而File类均有同名同参同返回值的方法，则File类均可以认为是这三个接口的实现。
尽管File类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是File类实现了这些接口，可以进行赋值
var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWriter = new(File)
var file4 ICloser = new(File)

在Go中，类的继承树并无意义，你只需要知道这个类实现了哪些方法，每个方法是啥含义就足够了
其二，实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理。接口由使用方按需定义，而不用事前规划
其三，不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦合。接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口
var file2 IReader = new(File)
file2.Read(nil)

3. 接口赋值
接口赋值在Go语言中分为如下两种情况：
	将对象实例赋值给接口；
	将一个接口赋值给另一个接口

(1). 将某种类型的对象实例赋值给接口，这要求该对象实例实现了接口要求的所有方法

赋值可以有两种方式
var b LessAdder = &a  (1)
var b LessAdder = a   (2)

通过类型type.go的示例中，可以知道对象方法可以写成值传递和指针传递，值传递是值的拷贝，不改变原值，要想改变原值需要指针传递，及在方法名前添加(f *File)类似的this指针引用
首先,Go语言可以根据下面的Less函数生成一个新的Less函数
例如原来有个
func (a Integer) Less(b Integer) bool {
 return a < b
}
自动生成下面的含指针方法。
func (a *Integer) Less(b Integer) bool {
 return (*a).Less(b)
}
无法反向生成
func (a Integer) Add(b Integer) {
 (&a).Add(b)
}
因此可以解释来说，因为Go会自动生成带指针的函数，而无法反向生成。因此Add方法是不会反向生成值传递方法的。而Add只有指针传递的方法，
那么将a以值传递方式赋值的时候就会出现编译错误，编译器会让你自动加上&指针传递
假如仅实现了Adder接口，则只能按照指针对象进行传递，因为没有生成值传递方法。
假如仅实现了Lesser接口,则因为GO自动生成引用传递，则使用值传递a和指针传递&a都可以匹配到对应的方法

简言之的解释就是要调用的方法是需要指针类型调用还是值类型调用。

摘自：Go语言调用接口方法的值传递与指针传递的区别 https://blog.csdn.net/weixin_38089997/article/details/108767637
若定义的方法是指针类型作为入参，则实参给值时必须是指针类型的，如果是值类型就会报错
而如果定义的方法参数是值类型，则调用方法传递的参数是值类型还是指针类型都是可以的。

在Java中，我们不用考虑值或是引用，一般除基本类型外都会被认为引用类型，会直接修改并作用到调用的参数。

(2) 将一个接口赋值给另一个接口
在Go语言中，只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值
假如在两个包中，分别有两个名称不同的接口，但是方法列表一样，但是顺不一样，那么在Go语言中，这两个接口实际上并无区别
	任何实现了one.ReadWriter接口的类，均实现了two.IStream；
	任何one.ReadWriter接口对象可赋值给two.IStream，反之亦然；
	在任何地方使用one.ReadWriter接口与使用two.IStream并无差异
以下这些代码可编译通过：
var file1 two.IStream = new(File)
var file2 one.ReadWriter = file1
var file3 two.IStream = file2

接口赋值并不要求两个接口必须等价。如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A
及接口A在左侧，接口B在右侧，接口B的方法列表是包含接口A的，因此可以把接口B赋值给接口A，但反过来不成立。

4. 接口查询
类似Java中的instance关键字，确定是否实现了某接口。
接口查询是否成功，要在运行期才能够确定。它不像接口赋值，编译器只需要通过静态类型检查即可判断赋值是否可行
//if file5,ok:=file2.(IReader);ok {
//	file5.Read(nil)
//}

5. 类型查询
即获取该实例的类型具体是什么类型
//switch a:=file2.(type){
//case *File:
//	fmt.Println(a)
//}
该方式必须和switch联合使用，否则会出现编译器报错

6. Any类型
类似Java中的Object类，是所有类型的父类
interface{}
由于Go语言中任何对象实例都满足空接口interface{}，所以interface{}看起来像是可以指向任何对象的Any类型
当函数可以接受任意的对象实例时，我们会将其声明为interface{}，最典型的例子是标准库fmt中PrintXXX系列的函数

*/
func main() {
	//var file2 IReader = new(File)
	//file2.Read(nil)

	//if file5,ok:=file2.(IReader);ok {
	//	file5.Read(nil)
	//}
	//switch a:=file2.(type){
	//case *File:
	//	fmt.Println(a)
	//}

	//var a Integer = 1
	//var b LessAdder = &a
	//b.Add(2)
	//var a Integer = 1
	//var b Adder = &a
	//b.Add(2)
	//var a Integer = 1
	//var b Lesser = a
	//b.Less(2)

}

//type Integer int
//
//func (a Integer) Less(b Integer) bool {
//	return a<b
//}
//func (a *Integer) Add(b Integer) {
//	*a += b
//}
//type LessAdder interface {
//	Less(b Integer) bool
//	Add(b Integer)
//}
//type Adder interface {
//	Add(b Integer)
//}
//type Lesser interface {
//	Less(b Integer) bool
//}
//
//
/////////////////////////////////////////////////
//type File struct {
//
//}
//
//func (f *File) Read(buf []byte) (n int,err error)  {
//	fmt.Println("调用了read方法")
//	return 0, err
//}
//
//type IReader interface {
//	Read(buf []byte) (n int,err error)
//}
//
//type IWriter interface {
//	Writer(buf []byte) (n int,err error)
//}
//type ICloser interface {
//	Close() error
//}
