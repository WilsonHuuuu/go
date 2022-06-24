package main

import "fmt"

func f1() {
	name := "tom "
	age := "20"
	//在函数内部做一些运算
	f1 := func() string {
		return name + age
	}

	msg := f1()
	fmt.Printf("msg: %v\n", msg)
}

func main() {
	/*	// 匿名函数
		max := func(a int, b int) int {
			if a > b {
				return a
			} else {
				return b
			}
		}

		r := max(1, 2)
		fmt.Printf("r: %v\n", r)
	*/

	// 自己调用
	r := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}(1, 2)

	fmt.Printf("r: %v\n", r)

	f1()
}
