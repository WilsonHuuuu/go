package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int, 0)

var chanStr = make(chan string)

func main() {

	go func() {
		chanInt <- 100
		chanStr <- "hello"
		// 有close则会返回默认值
		// defer close(chanInt)
		// defer close(chanStr)
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Printf("r: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("r: %v\n", r)
		default:
			fmt.Println("default") // 没有就会发生死锁
		}

		time.Sleep(time.Second)
	}
}
