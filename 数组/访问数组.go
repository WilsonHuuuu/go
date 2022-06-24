package main

import "fmt"

// 可以给数组赋值
func f1() {
	var a1 [2]int
	a1[0] = 100
	a1[1] = 200
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
	fmt.Println("-------------------")
	a1[0] = 1000
	a1[1] = 2000
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[1]: %v\n", a1[1])
}

//通过内置函数 len 来访问数组长度
func f2() {
	var a1 = [3]int{1, 2, 3}
	fmt.Printf("len(a1): %v\n", len(a1))

	var a2 = [...]int{1, 2, 3, 4}
	fmt.Printf("len(a2): %v\n", len(a2))
}

// 数组遍历
func f3() {
	var a1 = [3]int{1, 2, 3}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%v]: %v\n", i, a1[i])
	}
}

//for_range 数组遍历
func f4() {
	var a1 = [3]int{1, 2, 3}
	for i, v := range a1 { //可以直接使用 forr
		fmt.Printf("a1[%v]: %v\n", i, v)
	}
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
