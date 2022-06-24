package main

import "fmt"

func main() {
	// 1、初始化变量可以直接赋初值
	// 类型名称（string、int、bool）可以不写
	var name string = "Tom"
	var age int = 20
	var m bool = true

	// 2、批量初始化
	// var name, age, b = "Tom", 20, true

	// 3、短变量声明
	// name := "Tom"
	// age := 20
	// m := true

	// 4、匿名变量
	// func getNameAndAge()(string,int){
	// 	return "tom",20
	// }
	// func main()  {
	// 	Name, Age := getNameAndAge()
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)
	// }

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("m: %v\n", m)

	// 输出行命令：name + . + printf

}
