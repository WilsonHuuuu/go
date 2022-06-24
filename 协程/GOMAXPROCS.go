package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("a: %v\n", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("b: %v\n", i)
	}
}

func main() {

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(2) // 括号里如果是一则无法看到输出为交替

	go a()
	go b()

	time.Sleep(time.Second)
}
