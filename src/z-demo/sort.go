package main

import (
	"fmt"
	"sort"
)

type StuScore struct {
	name  string
	score int
}

type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}

func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}
	// 打印未排序的 stus 数据
	fmt.Println("Default:\n\t", stus)
	//StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
	sort.Sort(stus)
	// 判断是否已经排好顺序，将会打印 true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	// 打印排序后的 stus 数据
	fmt.Println("Sorted:\n\t", stus)

	s := []int{1, 2, 6, 23, 3, 5}
	sort.Ints(s)
	fmt.Println(s)
}
