package main

import "fmt"

func f1() {
	a := 1
	b := 2
	if a > b {
		fmt.Printf("a: %v\n", a)
	} else {
		fmt.Printf("b: %v\n", b)
	}
}

func f2() {
	var name string
	var age int
	var email string
	fmt.Println("请输入name，age，email，用空格分隔")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
	// 在输入时需要使用终端输入，且当文件在目录下的文件夹中时需要输入: 包名称\文件名.go
}

func f3() {
	var num int
	fmt.Println("请输入一个数字")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	// 在终端想要输入上一个命令直接按下上箭头即可
}

func main() {
	// f1()
	// f2()
	f3()
}
