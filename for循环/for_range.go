package main

import (
	"fmt"
)

func f1() {
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

func f2() {
	var s = []int{1, 2, 3}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
}

func f3() {
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "sfjsj"

	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
}

func f4() {
	s := "hello world"
	for _, v := range s {
		fmt.Printf("v: %c\n", v)
	}
}

func main() {
	// f1()
	// f2()
	// f3()
	f4()
}
