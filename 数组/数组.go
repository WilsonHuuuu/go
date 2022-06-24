package main

import "fmt"

func f1() {
	var a1 [2]int
	var a2 [3]string
	var a3 [2]bool

	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}

// 数组的初始化
// bool 类型默认：false
func f2() {
	var a1 = [2]int{1, 2}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [2]string{"hello", "world"}
	fmt.Printf("a2: %v\n", a2)
	var a3 = [2]bool{true, false}
	fmt.Printf("a3: %v\n", a3)
}

// 可以省略数组的长度,利用 ... 代替数组长度
func f3() {
	var a1 = [...]int{1, 2}
	fmt.Printf("len(a1): %v\n", len(a1))
	fmt.Printf("a1: %v\n", a1)
	var a2 = [...]string{"hello", "world"}
	fmt.Printf("len(a2): %v\n", len(a2))
	fmt.Printf("a2: %v\n", a2)
	var a3 = [...]bool{true, false}
	fmt.Printf("a3: %v\n", a3)
}

// 可以指定索引值来初始化
func f4() {
	var a1 = [...]int{0: 1, 3: 100, 5: 10}
	fmt.Printf("a1: %v\n", a1)
	var a2 = [...]bool{2: true, 4: false}
	fmt.Printf("a2: %v\n", a2)
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
