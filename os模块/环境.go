package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取所有环境变量
	// fmt.Printf("os.Environ(): %v\n", os.Environ())

	// 抓取某个环境变量
	s := os.Getenv("GOPATH")
	fmt.Printf("s: %v\n", s)

	// 抓取某个但是可以知道有没有此环境变量
	s2, b := os.LookupEnv("GOPATH1")
	if b {
		fmt.Printf("s2: %v\n", s2)
	} else {
		fmt.Println("没有此环境变量")
	}

	// 设置环境变量
	os.Setenv("env1", "env1")
	// 替换
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

}
