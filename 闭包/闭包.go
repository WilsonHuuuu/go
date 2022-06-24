package main

import "fmt"

// 闭包可以理解为定义在一个函数内部的函数,闭包指的是一个函数和与其相关的引用环境组合而成的实体

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}

	return add, sub
}

// 值会被保留下来
func main() {
	f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	r = f(30)
	fmt.Printf("r: %v\n", r)
	fmt.Println("-------------")
	// 又被重新定义,作用域是和函数绑定的
	f1 := add()
	r = f1(100)
	fmt.Printf("r: %v\n", r)
	r = f1(100)
	fmt.Printf("r: %v\n", r)

	fmt.Println("-------------")
	add, sub := cal(100)
	r = add(100)
	fmt.Printf("r: %v\n", r)
	r = sub(50)
	fmt.Printf("r: %v\n", r)
	fmt.Println("--------------")
	add1, sub1 := cal(100)
	r = add1(1)
	fmt.Printf("r: %v\n", r)
	r = sub1(1)
	fmt.Printf("r: %v\n", r)
}
