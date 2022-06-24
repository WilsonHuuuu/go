package main

import "fmt"

type Pet interface {
	eat(string) string
}

type Dog struct {
	name string
}

// 不会修改值
// func (dog Dog) eat(name string) string {
// 	dog.name = "超级无敌"
// 	fmt.Printf("name: %v\n", name)
// 	return "吃得好"
// }

// 指针会修改值
func (dog *Dog) eat(name string) string {
	dog.name = "超级无敌"
	fmt.Printf("name: %v\n", name)
	return "吃得好"
}

func main() {
	// dog := Dog{
	// 	name: "无敌",
	// }

	// s := dog.eat("火鸡")
	// fmt.Printf("s: %v\n", s)
	// fmt.Printf("dog: %v\n", dog)

	dog := &Dog{
		name: "无敌",
	}

	s := dog.eat("火鸡")
	fmt.Printf("s: %v\n", s)
	fmt.Printf("dog: %v\n", dog)
}
