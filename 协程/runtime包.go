package main

import (
	"fmt"
	"runtime"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg: %v\n", msg)
	}
}

func main() {
	go show("java")
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 执行其他子协程
		fmt.Printf("\"golang\": %v\n", "golang")
	}
	fmt.Println("end")
}
