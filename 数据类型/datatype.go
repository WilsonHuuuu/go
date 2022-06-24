package main

import "fmt"

func main() {
	var name string = "Tom"
	age := 20
	b := true
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("b: %v\n", b)

	a := 100
	p := &a
	fmt.Printf("p: %v\n", p)

	c := [2]int{2, 3}
	fmt.Printf("c: %v\n", c)
}
