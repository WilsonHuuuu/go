package main

import "fmt"

func f1() {
	var m1 map[string]string
	m1 = make(map[string]string)
	fmt.Printf("m1: %v\n", m1)
	fmt.Printf("m1: %T\n", m1)
}

// map初始化
func f2() {
	var m1 = map[string]string{"name": "tom", "age": "20"}
	fmt.Printf("m1: %v\n", m1)

	m2 := make(map[string]string)
	m2["name"] = "tom"
	m2["age"] = "20"
	fmt.Printf("m2: %v\n", m2)
}

// 访问map
func f3() {
	var m1 = map[string]string{"name": "tom", "age": "20"}
	var key = "name"
	var value = m1[key]
	fmt.Printf("value: %v\n", value)
}

// 判断元素是否存在
func f4() {
	var m1 = map[string]string{"name": "tom", "age": "20"}
	var k1 = "name"
	var k2 = "age1"
	v, ok := m1[k1]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("------------")
	v, ok = m1[k2]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
