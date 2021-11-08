package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/*
		a := "gopher"
		b := "hello World"
		// equals 比较
		fmt.Println(strings.EqualFold(a, b))
		// 是否包含
		fmt.Println(strings.Contains(b,"hello"))
		// 任何一个字符串中的字符在s中就返回true
		fmt.Println(strings.ContainsAny(b,"xyzh"))
		// 任何一个字符或数字型的字符码在s中就返回true
		fmt.Println(strings.ContainsRune(b,'e'))
		// 字符串出现的数量
		fmt.Println(strings.Count(b,"l"))

		//使用一个或多个连续空格分隔
		fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
		fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))
		fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))

		//是否包含前缀和后缀
		fmt.Println(strings.HasPrefix("Gopher","Go"))
		fmt.Println(strings.HasSuffix("Amigo","o"))
		// 索引位置
		fmt.Println(strings.Index("Gopoher","o"))

		han := func(c rune) bool {
			return unicode.Is(unicode.Han, c) // 汉字
		}
		fmt.Println(strings.IndexFunc("Hello, world", han))
		fmt.Println(strings.IndexFunc("Hello, 世界", han))
		// join
		fmt.Println(strings.Join([]string{"name=xxx","age=xxx"},"&"))
		// repeat 重复
		fmt.Println(strings.Repeat("name",2))

		// 字符串替换
		fmt.Println(strings.Replace("abcd","a","m",1))
		// map映射，每个字符进行映射，如果返回<0则舍弃该字符
		mapping := func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'Z': // 大写字母转小写
				return r + 32
			case r >= 'a' && r <= 'z': // 小写字母不处理
				return r
			case unicode.Is(unicode.Han, r): // 汉字换行
				return '\n'
			}
			return -1 // 过滤所有非字母、汉字的字符
		}
		fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher..."))
	*/

	// trim 匹配第二个参数任意一个字符，直到第一个字符串的后一个字符无法在第二个参数中匹配到
	fmt.Println(strings.TrimLeft("abcdefg", "adgbc"))
	// 必须完全匹配
	fmt.Println(strings.TrimPrefix("abcdefg", "ab"))
	// 左右间隔符匹配去除
	fmt.Println(strings.TrimSpace(" a b c d e fg\n"))
	// 带函数匹配,返回true则请求该字符
	f := func(r rune) bool {
		return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	}
	fmt.Println(strings.TrimFunc("abc我是你zxc", f))
	// 匹配器,多个oldnew对匹配
	r := strings.NewReplacer("&lt;", "<", "&gt;", ">")
	fmt.Println(r.Replace("我的年龄&gt;他的年龄&lt;她的年龄"))

}
