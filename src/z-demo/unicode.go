package main

import (
	"fmt"
	"unicode"
)

/*
对unicode支持的三个包
unicode
unicode/utf8
unicode/utf16

一个rune代表一个unicode编码,使用单引号代表一个字符

unicode包包含基本的字符判断函数
func IsControl(r rune) bool  // 是否控制字符
func IsDigit(r rune) bool  // 是否阿拉伯数字字符，即 0-9
func IsGraphic(r rune) bool // 是否图形字符
func IsLetter(r rune) bool // 是否字母
func IsLower(r rune) bool // 是否小写字符
func IsMark(r rune) bool // 是否符号字符
func IsNumber(r rune) bool // 是否数字字符，比如罗马数字Ⅷ也是数字字符
func IsOneOf(ranges []*RangeTable, r rune) bool // 是否是 RangeTable 中的一个
func IsPrint(r rune) bool // 是否可打印字符
func IsPunct(r rune) bool // 是否标点符号
func IsSpace(r rune) bool // 是否空格
func IsSymbol(r rune) bool // 是否符号字符
func IsTitle(r rune) bool // 是否 title case
func IsUpper(r rune) bool // 是否大写字符
func Is(rangeTab *RangeTable, r rune) bool // r 是否为 rangeTab 类型的字符
func In(r rune, ranges ...*RangeTable) bool  // r 是否为 ranges 中任意一个类型的字符

b := unicode.Is(unicode.Han,'我') //true

utf8包主要负责rune和byte之间的转换


utf16包负责rune额unit16数组之间的转换



*/

func main() {
	b := unicode.Is(unicode.Han, '我')
	fmt.Println(b)
}
