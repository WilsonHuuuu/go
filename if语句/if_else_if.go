package main

import (
	"fmt"
)

func f1() {
	score := 90

	if score >= 60 && score <= 70 {
		fmt.Println("c")
	} else if score > 70 && score <= 90 {
		fmt.Println("b")
	} else {
		fmt.Println("a")
	}
}

func f2() {
	var c string
	fmt.Println("请输入一个字符：")
	fmt.Scan(&c)

	if c == "M" {
		fmt.Println("Monday")
	} else if c == "T" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&c)
		if c == "u" {
			fmt.Println("Tuesday")
		} else {
			fmt.Println("Thursday")
		}
	} else if c == "W" {
		fmt.Println("Wednesday")
	} else if c == "F" {
		fmt.Println("Friday")
	} else if c == "S" {
		fmt.Println("请输入第二个字符：")
		fmt.Scan(&c)
		if c == "a" {
			fmt.Println("Saturday")
		} else {
			fmt.Println("Sunday")
		}
	} else {
		fmt.Println("输入不正确")
	}
}
func main() {
	// f1()
	f2()
}
