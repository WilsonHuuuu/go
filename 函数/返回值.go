package main

import "fmt"

// 没有返回值
func f1() {
	fmt.Println("没有参数也没有返回值")
}

// 有一个返回值
func sum(a int, b int) int {
	return a + b
}

// 多个返回值，且在 return 中指定返回的内容
func f2() (name string, age int) {
	name = "tom"
	age = 20
	return name, age
}

// 多个返回值，名称没有使用
func f3() (name string, age int) {
	name = "tom"
	age = 20
	return
}

// return 覆盖命名返回值，返回值名称没有被使用
func f4() (name string, age int) {
	n := "tom"
	a := 30
	return n, a
}
func main() {
	// f1()
	// name, age := f2()
	// name, age := f3()
	n, a := f3()
	fmt.Printf("name: %v\n", n)
	fmt.Printf("age: %v\n", a)

}
