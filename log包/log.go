package main

import (
	"fmt"
	"log"
	"os"
)

func test1() {
	log.Print("my log")
	log.Printf("my log %d", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)
}

func test2() {
	defer fmt.Println("panic 结束后再执行")
	log.Print("my log")
	log.Panic("my panic") //panic 会执行 defer
	fmt.Println("end")
}

func test3() {
	defer fmt.Println("defer")
	log.Print("my log")
	log.Fatal("fatal") //fatal 不会执行 defer
	fmt.Println("end")
}

// 两种初始化 log 输出
func init() {
	/* log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //可以设置 log 的输出：Ldate：日期；Ltime：时间；Lshortfile：文件夹；Llongfile：路径
	log.SetPrefix("MYLOG:")                              //设置前缀
	f, err := os.OpenFile("a.log", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal("日志文件错误")
	}

	log.SetOutput(f) */

	f, err := os.OpenFile("a.log", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal("日志文件错误")
	}

	logger = log.New(f, "Mylog:", log.Ldate|log.Ltime|log.Lshortfile)
}

var logger *log.Logger

// log包实现的是打印日志的功能
func main() {
	/* 	i := log.Flags()         //获得配置
	   	fmt.Printf("i: %v\n", i) //打印配置
	   	log.Print("my log") */

	logger.Print("my log")
}
