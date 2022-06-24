package main

import (
	"fmt"
	"runtime"
	"time"
)

func showMsg() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			runtime.Goexit() // 退出当前协程
		}
	}
}

func main() {
	go showMsg() //子协程
	time.Sleep(time.Second)
	fmt.Println("end") //主协程
}
