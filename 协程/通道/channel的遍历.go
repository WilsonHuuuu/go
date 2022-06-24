package main

import "fmt"

var c = make(chan int)

func main() {

	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}() // （）：自己调用自己

	//使用死循环进行遍历，记得要加上 break ，不然程序会一直跑下去
	for {
		v, ok := <-c
		if ok {
			fmt.Printf("v: %v\n", v)
		} else {
			break
		}
	}

	//建议使用for range
	for v := range c {
		fmt.Printf("v: %v\n", v)
	}

	//如果没有 close 若下面的循环大于 go 协程里的循环，则会出现死锁的情况
	for i := 0; i < 3; i++ {
		r := <-c
		fmt.Printf("r: %v\n", r)
	}
}
