package main

import (
	"io"
	"os"
)

func main() {

	// # error接口
	//对于大多数函数，如果要返回错误，大致上都可以定义为如下模式，将error作为多种返回值中的最后一个，但这并非是强制要求
	//func Foo(param int)(n int, err error) {
	//	// ...
	//}
	//调用时的代码建议按如下方式处理错误情况：
	//n, err := Foo(0)
	//if err != nil {
	//	// 错误处理
	//} else {
	//	// 使用返回值n
	//}

	// # defer关键字，将关键字放在需要释放资源的方法前面，则在方法返回结束之前会正常关闭流
	//CopyFile();
	//如果觉得一句话干不完清理的工作，也可以使用在defer后加一个匿名函数的做法：
	//defer func() {
	//	// 做你复杂的清理工作
	//} ()
	//另外，一个函数中可以存在多个defer语句，因此需要注意的是，defer语句的调用是遵照先进后出的原则，即最后一个defer语句将最先被执行

	// Go语言引入了两个内置函数panic()和recover()以报告和处理运行时错误和程序中的错误场景
	// func panic(interface{})
	// func recover() interface{}
	//当在一个函数执行过程中调用panic()函数时，正常的函数执行流程将立即终止，但函数中
	//之前使用defer关键字延迟执行的语句将正常展开执行，之后该函数将返回到调用函数，并导致
	//逐层向上执行panic流程，直至所属的goroutine中所有正在执行的函数被终止。错误信息将被报
	//告，包括在调用panic()函数时传入的参数，这个过程称为错误处理流程

	//recover()函数用于终止错误处理流程。一般情况下，recover()应该在一个使用defer
	//关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复
	//过程（使用recover关键字），会导致该goroutine所属的进程打印异常信息后直接退出
	//defer func() {
	//	if r := recover(); r != nil {
	//		log.Printf("Runtime error caught: %v", r)
	//	}
	//}()
	//foo()
	//无论foo()中是否触发了错误处理流程，该匿名defer函数都将在函数退出时得到执行。假
	//如foo()中触发了错误处理流程，recover()函数执行将使得该错误处理过程终止。如果错误处
	//理流程被触发时，程序传给panic函数的参数不为nil，则该函数还会打印详细的错误信息

}

func CopyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
