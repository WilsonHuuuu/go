package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	site := WebSite{Name: "duoke360"}
	fmt.Printf("site: %v\n", site)
	fmt.Printf("site: %#v\n", site) //加 # 后会显示哪个包哪个类下面的
	fmt.Printf("site: %T\n", site)  //输出类型

	a := 100
	fmt.Printf("a: %T\n", a)

	b := true
	fmt.Printf("b: %t\n", b)

	// 整型字符输出
	i := 8
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)   //二进制表示
	fmt.Printf("i: %c\n", i)   //Unicode码
	fmt.Printf("i: %x\n", 100) //十六进制

	x := 100
	p := &x
	fmt.Printf("p: %v\n", p) //指针
}
