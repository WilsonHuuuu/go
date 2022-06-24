package main

import (
	"fmt"
	"time"
)

func main() {
	//方法1
	timer := time.NewTimer(time.Second * 2) //() 里面是时间
	fmt.Printf("time.Now(): %v\n", time.Now())
	t1 := <-timer.C //阻塞的指定时间
	fmt.Printf("t1: %v\n", t1)

	//方法2
	fmt.Printf("time.Now(): %v\n", time.Now())
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Printf("time.Now(): %v\n", time.Now())

	//方法3
	time.Sleep(time.Second)

	//方法4
	<-time.After(time.Second * 2)
	fmt.Println("....")

	//方法5
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("func")
	}()

	s := timer.Stop()
	if s {
		fmt.Println("stop")
	}

	time.Sleep(time.Second * 3)
	fmt.Println("main end")

	//方法6
	fmt.Println("before")
	timer3 := time.NewTimer(time.Second * 5) //原来设置是5s
	timer3.Reset(time.Second * 1)            //重新设置为1s
	<-timer3.C
	fmt.Println("after")
}
