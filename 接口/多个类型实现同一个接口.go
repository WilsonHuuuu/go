package main

import "fmt"

type Pet interface {
	eat()
}

type Dog struct {
}

type Cat struct {
}

func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func main() {
	dog := Dog{}
	cat := Cat{}
	dog.eat()
	cat.eat()

	// 第二种写法
	var pet Pet
	pet = Dog{}
	pet.eat()
	pet = Cat{}
	pet.eat()
}
