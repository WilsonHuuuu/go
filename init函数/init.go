package main

import "fmt"

var i int = initVar()

// init 函数必须没有参数和返回值
// 先执行 init 函数
// 哪个 init 函数写在前面就先运行谁

func init() {
	fmt.Println("init2")
}

func init() {
	fmt.Println("init")
}

func initVar() int {
	fmt.Println("initVar")
	return 100
}

func main() {
	fmt.Println("main")
}
