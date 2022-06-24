package main

import "fmt"

func f1() {
	m1 := map[string]string{"name": "tom", "age": "20"}
	for k, v := range m1 {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func main() {
	f1()
}
