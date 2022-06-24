package main

import "fmt"

// 普通指针
func f1() {
	var name string
	name = "tom"
	var p_name *string
	p_name = &name

	fmt.Printf("name: %v\n", name)
	fmt.Printf("p_name: %v\n", p_name)
	fmt.Printf("p_name: %v\n", *p_name)
}

func f2() {
	type Person struct {
		id   int
		age  int
		name string
	}

	// 初始化结构体
	tom := Person{
		id:   101,
		age:  20,
		name: "tom",
	}

	var p_person *Person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %v\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)
}

func f3() {
	type Person struct {
		id   int
		age  int
		name string
	}

	var tom = new(Person)
	tom.id = 101
	tom.name = "tom"
	// (*tom).id=101
	// (*tom).name="tom"

	fmt.Printf("tom: %v\n", *tom)
}

func main() {
	// f1()
	// f2()
	f3()
}
