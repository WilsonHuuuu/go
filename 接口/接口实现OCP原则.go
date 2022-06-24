package main

import "fmt"

// OCP(Open-Closed Principle) 开闭原则：对扩展是开发的，对修改是关闭的
type Pet interface {
	eat()
	sleep()
}

type Dog struct {
}

type Cat struct {
}

// Dog 实现Pet 接口
func (dog Dog) eat() {
	fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
	fmt.Println("dog sleep...")
}

// Cat 实现 Pet 接口
func (cat Cat) eat() {
	fmt.Println("cat eat...")
}

func (cat Cat) sleep() {
	fmt.Println("cat sleep...")
}

type Person struct {
}

func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	person := Person{}
	person.care(dog)
	person.care(cat)
}
