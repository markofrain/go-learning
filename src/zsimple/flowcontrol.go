package main

import "fmt"

func main() {

	// # 条件语句
	// a := ifelse(1)
	// fmt.Println(a)

	// # 选择语句
	// 类似ifelse判断，只会进入其中一个。
	// 甚至可以在case中包含表达式，而switch后的条件则不是必须的。类似数据库的case when then end
	i := 4
	switchtest(i)
	fmt.Println()
	// # 循环
	//str := "123123"
	//for i := 0; i < len(str); i++ {
	//	fmt.Print(string(str[i]))
	//}
	//fmt.Println()
	//for i,v := range str {
	//	fmt.Print(i,"-",string(v) ," ")
	//}

	// 无限循环
	//num := 0
	//for{
	//	num++
	//	fmt.Println(num)
	//	if num>100 {
	//		break
	//	}
	//}

	// 跳出指定循环
	//JLoop:
	//for j := 0; j < 5; j++ {
	//	for i := 0; i < 10; i++ {
	//		if i > 5 {
	//			break JLoop
	//		}
	//		fmt.Println(i)
	//	}
	//}

	// # 跳转语句 goto
	gototest()

}

/**
跳转语句
*/
func gototest() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func switchtest(i int) {
	switch i {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}
}

func ifelse(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}
