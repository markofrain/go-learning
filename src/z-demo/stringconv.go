package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1
	// 整数转字符串
	fmt.Println("my age is " + strconv.FormatInt(int64(num), 4))
	// 布尔转字符串
	fmt.Println(strconv.FormatBool(true) + " asd")
	// 要转换的浮点数,格式化类型,有效数字(所有位数;fmt='b'时无效),字节位数(32,64)
	// 对于 'e', 'E' 和 'f'，有效数字用于小数点之后的位数；对于 'g' 和 'G'，则是所有的有效数字
	fmt.Println(strconv.FormatFloat(1234.2346, 'g', 2, 32))
	// 即转换后的底数的小数点后的位置
	fmt.Println(strconv.FormatFloat(1234.2346, 'e', 2, 32))
}
