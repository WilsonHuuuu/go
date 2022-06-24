package main

import "fmt"

// 格式： type + 结构体名称 + struct + {  成员定义  }
// 自定义一个类型
type Person struct {
	id   int
	name string
	age  int
}

type Customer struct {
	id, age int
	name    string
}

func main() {
	var tom Person
	tom.id = 101
	tom.age = 20
	tom.name = "tom"
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom.name: %v\n", tom.name)

	var kite Customer
	fmt.Printf("kite: %v\n", kite)

	// 匿名类型的结构体
	var jack struct {
		id   int
		name string
		age  int
	}

	jack.id = 101
	jack.age = 18
	jack.name = "jack"
	fmt.Printf("jack: %v\n", jack)
}
