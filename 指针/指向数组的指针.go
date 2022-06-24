package main

import "fmt"

// 格式： var + 变量名 + [] + * + 变量类型
// 指针不允许进行运算
func main() {
	a := [3]int{1, 2, 3}
	var pa [3]*int

	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(a); i++ {
		pa[i] = &a[i]
	}

	fmt.Printf("pa: %v\n", pa)

	for i := 0; i < len(pa); i++ {
		fmt.Printf("%v\n", *pa[i])
	}
}
