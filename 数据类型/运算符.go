package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 20

	x := a + b //第一次声明一定要加 ：
	fmt.Printf("x: %v\n", x)
	x = a - b //之后赋值不用加 ：
	fmt.Printf("x: %v\n", x)
	x = a / b
	fmt.Printf("x: %v\n", x)
	x = a * b
	fmt.Printf("x: %v\n", x)

	r := a == b
	fmt.Printf("r: %v\n", r)
	r = a > b
	fmt.Printf("r: %v\n", r)
	r = a < b
	fmt.Printf("r: %v\n", r)
	r = a >= b
	fmt.Printf("r: %v\n", r)
	r = a <= b
	fmt.Printf("r: %v\n", r)
	r = a != b
	fmt.Printf("r: %v\n", r)

	c := true
	d := false
	k := c && d
	fmt.Printf("k: %v\n", k) //与
	k = c || d
	fmt.Printf("k: %v\n", k)  //或
	fmt.Printf("c: %v\n", !c) //取反

	e := 4
	fmt.Printf("e: %b\n", e) //0100
	f := 8
	fmt.Printf("f: %b\n", f) //1000
	y := a & b               //参与运算的两数各对应的二进制相与
	fmt.Printf("y: %v\n", y)
	y = a | b //参与运算的两数各对应的二进制相或
	fmt.Printf("y: %v\n", y)
	y = a ^ b //参与运算的两数各对应的二进制位相异或，当两对应的二进位相异时，结果为1
	fmt.Printf("y: %v\n", y)
	y = e << 2 //左移两位
	fmt.Printf("y: %v\n", y)
	y = e >> 2 //右移两位
	fmt.Printf("y: %v\n", y)

}
