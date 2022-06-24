package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	//交替来运行函数
	//做网页登录的时候就需要多协程
	go showMsg("java")   //加了 go 启动了一个协程来执行
	go showMsg("golang") //没有加 go 则是 main 函数中的主协程

	time.Sleep(time.Millisecond * 2000) //如果给时间另外的协程去进行则可以进行其他的协程
	fmt.Println("main end")             //主函数退出程序就结束了，主协程结束其他协程也就结束
}
