package main

import "fmt"

func f1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

// 初始值可以写在 for 循环的外面，但是不能漏掉 ；
func f2() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

// 初始值和第二个条件可以写在 for 循环的外面
func f3() {
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

func main() {
	// f1()
	// f2()
	f3()
}
