package main

import "fmt"

// & (取地址)和 * （根据地址取值）
// 指针格式： var + 变量名 + 变量类型

func main() {
	var ip *int
	fmt.Printf("ip: %v\n", ip)

	var i int = 100
	ip = &i
	fmt.Printf("ip: %v\n", ip)  // 输出的是地址
	fmt.Printf("ip: %v\n", *ip) // 加了 * 后输出的是值

	var sp *string
	var s string = "hello"
	sp = &s
	fmt.Printf("sp: %T\n", sp)
	fmt.Printf("sp: %v\n", sp)

	var bp *bool

	var b bool = true
	fmt.Printf("bp: %v\n", bp)
	bp = &b
	fmt.Printf("bp: %v\n", *bp)

}
