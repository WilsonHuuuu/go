package main

import "fmt"

func main() {
	a := 1
	b := 2

	if a > b {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}

	// 根据布尔值判断
	flag1 := true
	if flag1 {
		fmt.Printf("flag1: %v\n", flag1)
	}

	// 可以在 if 中声明变量，但是作用范围仅限于 if 语句中
	if age := 20; age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}
