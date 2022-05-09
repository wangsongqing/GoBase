package controllers

import (
	"fmt"
	"sort"
)

// StuScore 学生成绩结构体
type StuScore struct {
	//姓名
	name string
	//成绩
	score int
}

type StuScores []StuScore

// Len Len()
func (s StuScores) Len() int {
	return len(s)
}

// Less Less():成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

// Swap Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func stateSlice11() {
	stus := StuScores{
		{"alan", 60},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90}}

	// fmt.Println("Default:")
	//原始顺序
	//for _, v := range stus {
	//	fmt.Println(v.name, ":", v.score)
	//}
	// fmt.Println()
	//StuScores已经实现了sort.Interface接口
	sort.Sort(stus)

	fmt.Println("Sorted:")
	//排好序后的结构
	for _, v := range stus {
		fmt.Println(v.name, ":", v.score)
	}

	//判断是否已经排好顺序，将会打印true
	fmt.Println("IS Sorted?", sort.IsSorted(stus))
}
