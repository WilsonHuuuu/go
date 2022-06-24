package main

import "fmt"

// 格式： func + 函数名 + （[函数参数名 + 函数的参数类型]） + [返回值名称 + 返回值的类型]
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func comp(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return max
}

func main() {
	r := sum(1, 2)
	fmt.Printf("r: %v\n", r)

	r = comp(1, 2)
	fmt.Printf("r: %v\n", r)
}
