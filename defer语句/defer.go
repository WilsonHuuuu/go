package main

import "fmt"

// defer 用于注册延迟条用
func f1() {
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")
}
func main() {
	f1()
}
