package main

import "fmt"

// 添加
func f1() {
	var s1 = []int{}
	s1 = append(s1, 100)
	s1 = append(s1, 200)
	s1 = append(s1, 300)
	fmt.Printf("s1: %v\n", s1)
}

// 删除：利用 append 删除 index ，a = append(a[: index], a[index+1 :]...)
func f2() {
	var s1 = []int{1, 2, 3, 4}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

// 修改
func f3() {
	var s1 = []int{1, 2, 3, 4}
	s1[1] = 100
	fmt.Printf("s1: %v\n", s1)
}

// 查询
func f4() {
	var s1 = []int{1, 2, 3, 4}
	var key = 2
	for i, v := range s1 {
		if v == key {
			fmt.Printf("i: %v\n", i)
		}
	}
}

// 拷贝
func f5() {
	var s1 = []int{1, 2, 3, 4}
	var s2 = make([]int, 4)
	copy(s2, s1)
	s2[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	f5()
}
