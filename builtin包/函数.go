package main

import "fmt"

func testAppend() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i: %v\n", i)
	s1 := []int{4, 5, 6}
	i2 := append(s, s1...) //将两个切片合在一起
	fmt.Printf("i2: %v\n", i2)
}

func testLen() {
	s := "hello world"
	fmt.Printf("len(s): %v\n", len(s))
	s1 := []int{4, 5, 6}
	fmt.Printf("len(s1): %v\n", len(s1))
}

func testPrint() {
	name := "tom"
	age := 20
	print(name, " ", age, "\n")
	fmt.Println("-----------------")
	println(name, " ", age) //可以不用加 \n 来换行
}

func testPanic() {
	defer fmt.Println("panic后还执行")
	panic("寄") //panic后不执行
	fmt.Println("end")
}

func testNew() {
	b := new(bool)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %v\n", *b) //要用指针输出值
	i := new(int)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("i: %v\n", *i)
	s := new(string)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("s: %v\n", *s)
}

func testMake() {
	var p *[]int = new([]int)
	fmt.Printf("p: %v\n", p)
	v := make([]int, 10)
	fmt.Printf("v: %v\n", v)
}

func main() {

}
