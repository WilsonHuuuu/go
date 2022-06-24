package main

import "fmt"

func main() {
	type Person struct {
		id, age int
		name    string
	}

	// 没有初始化的值默认为 0
	var tom Person
	tom = Person{
		id:   101,
		age:  20,
		name: "tom",
	}
	fmt.Printf("tom: %v\n", tom)

	// 直接按照结构体定义的数据类型输入，顺序一定要一致，不需要变量名（例如：id、age），不能混合使用
	var jack Person
	jack = Person{
		10101,
		20,
		"jack",
	}
	fmt.Printf("jack: %v\n", jack)

}
