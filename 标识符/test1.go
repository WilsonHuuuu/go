package main

import "fmt"

func main() {
	// 变量 + 变量名称 + 类型
	var name string
	var age int
	var m bool

	// var 1name string
	// var &test string
	// 下面是错误的命名

	// go语言声明变量后必须使用

	name = "Tom"
	age = 20
	m = true

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("m: %v\n", m)

}
