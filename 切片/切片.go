//切片可以理解为：可变长度的数组，底层是使用数组实现的
package main

import "fmt"

func f1() {
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)

}

func f2() {
	var s2 = make([]int, 2)
	fmt.Printf("s2: %v\n", s2)
}

// 访问切片的长度和容量
func f3() {
	var s1 = []int{1, 2, 3}
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
}

func main() {
	// f1()
	// f2()
	f3()
}
