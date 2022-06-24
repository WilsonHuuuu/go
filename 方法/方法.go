package main

import (
	"fmt"
)

// 结构体定义了属性
type Person struct {
	name string
}

// 属性和方法分开来写
func (per Person) eat() {
	fmt.Printf("%v,eat...", per.name)
}

// 将结构体和属性关联到了一起
func (per Person) sleep() {
	fmt.Printf("%v,sleep...\n", per.name)
}

type Customer struct {
	name string
}

func (customer Customer) login(name string, pwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func main() {
	per := Person{
		name: "tom",
	}
	per.eat()
	per.sleep()

	cus := Customer{
		name: "tom",
	}

	b := cus.login("tom", "123")
	fmt.Printf("b: %v\n", b)

}
