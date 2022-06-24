package main

import "fmt"

type Person struct {
	id   int
	name string
}

// 传递结构体本身不会修改内容
func showPerson(per Person) {
	per.id = 101
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

// 传递指针会修改内容
func showPerson2(per *Person) {
	per.id = 102
	per.name = "kite"
	// fmt.Printf("per: %v\n", per)
}

func main() {
	tom := Person{
		id:   100,
		name: "tom",
	}

	per := &tom
	fmt.Printf("tom: %v\n", tom)
	showPerson2(per)
	fmt.Println("-----------")
	fmt.Printf("per: %v\n", *per)

	// fmt.Printf("tom: %v\n", tom)
	// fmt.Println("--------------")
	// showPerson(per)
	// fmt.Printf("tom: %v\n", tom)
}
