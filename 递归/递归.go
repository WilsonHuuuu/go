package main

import "fmt"

// 循环求阶乘
func f1() {
	var s int = 1
	for i := 1; i <= 5; i++ {
		s *= i
	}
	fmt.Printf("s: %v\n", s)
}

// 递归法求阶乘
func f2(a int) int {
	if a == 1 {
		// 退出条件
		return 1
	} else {
		// 自己调用自己
		return a * f2(a-1)
	}
}

// 斐波那契数列
func f3(a int) int {
	if a == 1 || a == 2 {
		return 1
	} else {
		return f3(a-1) + f3(a-2)
	}
}

func main() {
	f1()
	i := f2(5)
	fmt.Printf("i: %v\n", i)
	j := f3(9)
	fmt.Printf("j: %v\n", j)
}
