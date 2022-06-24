package main

import "fmt"

// 类型定义
func f1() {
	type MyInt int
	var i MyInt
	i = 100
	fmt.Printf("i: %T,%v\n", i, i)
}

// 类型别名（没有定义新的类型，只是将 MyInt 类型变成了 int 类型）
func f2() {
	type MyInt = int
	var i MyInt
	i = 100
	fmt.Printf("i: %T,%v\n", i, i)
}

func main() {
	// f1()
	f2()
}
