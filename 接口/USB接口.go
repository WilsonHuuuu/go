package main

import "fmt"

// 声明了多少个接口就一定要实现多少个接口
type USB interface {
	read()
	write()
}

type Computer struct {
	name string
}

type Mobile struct {
	model string
}

func (m Mobile) read() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("read....")
}

func (m Mobile) write() {
	fmt.Printf("m.model: %v\n", m.model)
	fmt.Println("write....")
}

func (c Computer) read() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("read...")
}

func (c Computer) write() {
	fmt.Printf("c.name: %v\n", c.name)
	fmt.Println("write...")
}

func main() {
	c := Computer{
		name: "苹果",
	}

	m := Mobile{
		model: "4G",
	}

	c.read()
	c.write()

	m.read()
	m.write()
}
