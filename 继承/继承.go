package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("eat")
}

func (a Animal) sleep() {
	fmt.Println("sleep")
}

type Dog struct {
	Animal //可以理解为继承
	color  string
}

type Cat struct {
	Animal
	bbb string
}

func main() {
	dog := Dog{
		Animal{name: "tom", age: 2},
		"black",
	}

	dog.eat()
	dog.sleep()
	fmt.Printf("dog.color: %v\n", dog.color)
	fmt.Printf("dog.age: %v\n", dog.age)

	cat := Cat{
		Animal{name: "kite", age: 2},
		"英短",
	}
	cat.eat()
	dog.sleep()
	fmt.Printf("cat: %v\n", cat)
	fmt.Printf("cat.bbb: %v\n", cat.bbb)
}
